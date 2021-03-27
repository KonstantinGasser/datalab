package app

import (
	"context"
	"errors"
	"net/http"
	"sync"

	appSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/app_service"
	userSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// AppItem represents one App in the database ? do we need this? don't we have a def in the grpc already???
type AppItem struct {
	// mongoDB pk (document key)
	UUID        string `bson:"_id"`
	AppName     string `bson:"name"`
	OwnerUUID   string `bson:"owner_uuid"`
	OrgnDomain  string `bson:"orgn_domain"`
	Description string `bson:"description"`
	// Member is a list of user_uuids mapped to this app
	Member   []string `bson:"member"`
	Settings []string `bson:"setting"`
	AppToken string   `bson:"app_token"`
}

// AppItemLight is a minimum representation of an application
type AppItemLight struct {
	// mongoDB pk (document key)
	UUID    string `bson:"_id"`
	AppName string `bson:"name"`
}

// GetByID collects all the app details for a given appUUID
// it fetches the user data for the owner and all members from the user-service
func (app app) GetApp(ctx context.Context, mongo storage.Storage, userService userSrv.UserServiceClient, appUUID string) (int, *appSrv.ComplexApp, error) {

	var queryData AppItem
	err := mongo.FindOne(ctx, appDatabase, appCollection, bson.M{"_id": appUUID}, &queryData)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	// prepare data of app append with user data if calls successful
	var appData *appSrv.ComplexApp = &appSrv.ComplexApp{
		Name:        queryData.AppName,
		Description: queryData.Description,
		Settings:    queryData.Settings,
		AppToken:    queryData.AppToken,
	}

	// fetch user information needed: owner and all members
	// fetch concurrently
	var wait sync.WaitGroup
	// spin-up goroutine to get app member details
	wait.Add(1)
	var respUserList *userSrv.GetUserListResponse
	var userListErr error
	go func() {
		respUserList, userListErr = userService.GetUserList(ctx, &userSrv.GetUserListRequest{
			Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
			UuidList:   queryData.Member,
		})
		wait.Done()
	}()

	// spin-up goroutine to get app owner details
	wait.Add(1)
	var respOwner *userSrv.GetUserResponse
	var ownerErr error
	go func() {
		respOwner, ownerErr = userService.GetUser(ctx, &userSrv.GetUserRequest{
			Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
			CallerUuid: "", //request.GetCallerUuid(),
			ForUuid:    queryData.OwnerUUID,
		})
		wait.Done()
	}()
	// wait until calls are done
	wait.Wait()

	if userListErr == nil && respUserList.GetStatusCode() == 200 {
		// merge response list from user service to something the app service can understand.
		// since they do not share a common.proto with a ComplexUser message translation is required
		// yes I still need to figure out why I am not able it import and generate from different .proto files..
		for _, item := range respUserList.GetUserList() {
			appData.Member = append(appData.Member, &appSrv.ComplexUser{
				Uuid:          item.GetUuid(),
				Username:      item.GetUsername(),
				FirstName:     item.GetFirstName(),
				LastName:      item.GetLastName(),
				OrgnDomain:    item.GetOrgnDomain(),
				OrgnPosition:  item.GetOrgnPosition(),
				ProfileImgUrl: item.GetProfileImgUrl(),
			})
		}
	} else {
		logrus.Errorf("<%v>[app.GetApp] could not get member user data:%v\n", ctx_value.GetString(ctx, "tracingID"), userListErr)
	}
	if ownerErr == nil && respOwner.GetStatusCode() == 200 {
		// merge complexUser from user service to app service complexUser
		appData.Owner = &appSrv.ComplexUser{
			Uuid:          respOwner.GetUser().GetUuid(),
			Username:      respOwner.GetUser().GetUsername(),
			FirstName:     respOwner.GetUser().GetFirstName(),
			LastName:      respOwner.GetUser().GetLastName(),
			OrgnDomain:    respOwner.GetUser().GetOrgnDomain(),
			OrgnPosition:  respOwner.GetUser().GetOrgnPosition(),
			ProfileImgUrl: respOwner.GetUser().GetProfileImgUrl(),
		}
	} else {
		logrus.Errorf("<%v>[app.GetApp] could not get owner user data:%v\n", ctx_value.GetString(ctx, "tracingID"), ownerErr)
	}

	return http.StatusOK, appData, nil
}

// GetApps collects all apps for a requests owner UUID -> all apps where owner == forUUID will be returned
func (app app) GetAppList(ctx context.Context, mongo storage.Storage, forUUID string) (int, []AppItemLight, error) {

	var appList []AppItemLight
	if err := mongo.FindMany(ctx, appDatabase, appCollection, bson.D{{"owner_uuid", forUUID}}, &appList); err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, appList, nil
}

// DeleteApp deletes an app based on the provided appUUID
func (app app) DeleteApp(ctx context.Context, mongo storage.Storage, appUUID string) (int, error) {
	if err := mongo.DeleteOne(ctx, appDatabase, appCollection, bson.M{"_id": appUUID}); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// CreateApp some docs
func (app app) CreateApp(ctx context.Context, mongo storage.Storage, appItem AppItem) (int, error) {

	// duplicate names may exists in the system but owners can only hold unique app names
	exists, err := mongo.Exists(ctx, appDatabase, appCollection, bson.M{"appName": appItem.AppName, "ownerUUID": appItem.OwnerUUID})
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if exists {
		return http.StatusBadRequest, errors.New("duplicated app names are not possible")
	}

	if err := mongo.InsertOne(ctx, appDatabase, appCollection, appItem); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// AddMember adds all provided []member to the app if the the requested caller is the owner of the app
// must ensure that requested members belong to the same organization as the app owner
func (app app) AddMember(ctx context.Context, storage storage.Storage, ownerUUID, appUUID string, member []string) (int, error) {

	// filter must ensure that caller has permissions (aka is owner) of the app
	filterAppAndOwner := bson.M{
		"_id":        appUUID,
		"owner_uuid": ownerUUID,
	}

	updateQuery := bson.D{
		{
			"$addToSet", bson.M{
				"member": bson.M{
					"$each": member,
				},
			},
		},
	}
	// updated shows if documents have been updated or not
	updated, err := storage.UpdateOne(ctx, appDatabase, appCollection, filterAppAndOwner, updateQuery)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if updated == 0 {
		// not yet sure what to do with this information
		// return http.StatusForbidden, errors.New("user not permitted to modify app data")
	}
	return http.StatusOK, nil
}

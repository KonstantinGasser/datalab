package app

import (
	"context"
	"errors"
	"net/http"
	"sync"

	appSrv "github.com/KonstantinGasser/datalabs/protobuf/app_service"
	userSrv "github.com/KonstantinGasser/datalabs/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/service_app/pkg/storage"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetByID collects all the app details for a given appUUID
// it fetches the user data for the owner and all members from the user-service
func (app app) Get(ctx context.Context, storage storage.Storage, userService userSrv.UserClient, appUUID, callerUUID string) (int, *appSrv.ComplexApp, error) {

	// search for app_uuid where either owner_uuid or one of the members
	// match the callerUUID else not permitted
	appQuery := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.M{"_id": appUUID},
				bson.D{
					{
						Key: "$or",
						Value: bson.A{
							bson.M{"owner_uuid": callerUUID},
							bson.M{"member": callerUUID},
						},
					},
				},
			},
		},
	}
	var queryData AppItem
	err := storage.FindOne(ctx, appDatabase, appCollection, appQuery, &queryData)
	// check if no results have been found
	if err == mongo.ErrNoDocuments {
		return http.StatusBadRequest, nil, errors.New("could not find any match")
	}
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	// prepare data of app append with user data if calls successful
	var appData *appSrv.ComplexApp = &appSrv.ComplexApp{
		Uuid:        queryData.UUID,
		URL:         queryData.URL,
		Name:        queryData.AppName,
		Description: queryData.Description,
		Settings:    queryData.Settings,
		AppToken:    queryData.AppToken,
	}

	// fetch user information needed: owner and all members
	// fetch concurrently
	var wg sync.WaitGroup
	// spin-up goroutine to get app member details
	wg.Add(1)
	var respUserList *userSrv.GetListResponse
	var userListErr error
	go func() {
		respUserList, userListErr = userService.GetList(ctx, &userSrv.GetListRequest{
			Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
			UuidList:   queryData.Member,
		})
		wg.Done()
	}()

	// spin-up goroutine to get app owner details
	wg.Add(1)
	var respOwner *userSrv.GetResponse
	var ownerErr error
	go func() {
		respOwner, ownerErr = userService.Get(ctx, &userSrv.GetRequest{
			Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
			CallerUuid: "", //request.GetCallerUuid(),
			ForUuid:    queryData.OwnerUUID,
		})
		wg.Done()
	}()
	wg.Wait()

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

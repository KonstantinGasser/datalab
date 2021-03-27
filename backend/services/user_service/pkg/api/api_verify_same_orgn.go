package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/user"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) VerifySameOrgn(ctx context.Context, request *userSrv.VerifySameOrgnRequest) (*userSrv.VerifySameOrgnResposne, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[userService.AreInSameOrgn] received request\n", ctx_value.GetString(ctx, "tracingID"))

	// create Comparator as base to compare values with
	comparator := user.Comparator{
		Fetch: func() (map[string]interface{}, error) {
			var data map[string]interface{}
			err := srv.storage.FindOne(ctx, "datalabs_user", "user", bson.D{{"_id", request.GetBaseObject()}}, &data)
			if err != nil {
				return nil, err
			}
			return data, nil
		},
		By:    "orgn_domain",
		Value: map[string]interface{}{}, // since the Fetch is provided the Value will be assigned with the queried data
	}
	// create Comparable from the request data
	comparable := user.Comparable{
		Fetch: func() ([]map[string]interface{}, error) {
			var data []map[string]interface{}
			err := srv.storage.FindMany(ctx, "datalabs_user", "user", bson.D{{"_id", bson.M{"$in": request.GetCompareWith()}}}, &data)
			if err != nil {
				return nil, err
			}
			return data, nil
		},
		By:        "orgn_domain",
		StorageID: "_id",
	}

	status, result, err := srv.user.Compare(ctx, srv.storage, comparator, comparable)
	if err != nil {
		logrus.Errorf("<%v>[userService.AreInSameOrgn] could not compare items: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.VerifySameOrgnResposne{StatusCode: int32(status), Msg: "could not compare items"}, nil
	}
	return &userSrv.VerifySameOrgnResposne{
		StatusCode:    int32(status),
		Msg:           "all items have been compared",
		TruthfulValid: result.TruthfulValid,
		InvalidList:   result.MissItems,
	}, nil
}

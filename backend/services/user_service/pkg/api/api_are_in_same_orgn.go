package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/user"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) AreInSameOrgn(ctx context.Context, request *userSrv.AreInSameOrgnRequest) (*userSrv.AreInSameOrgnResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[userService.AreInSameOrgn] received request\n", ctx_value.GetString(ctx, "tracingID"))

	// create Comparator as base to compare values with
	comparator := user.Comparator{
		Filter: bson.M{"_id": request.GetCompareTo()},
		By:     "orgn_domain",
		Value:  map[string]interface{}{}, // since the Filter is provided the Value will be assigned with the queried data
	}
	// create Comparable from the request data
	comparable := user.Comparable{
		// Filter selects all users by the uuid given in the request and returns only the orgn_domain
		Filter: bson.D{
			{"_id", bson.M{"$in": request.GetValues()}},
			// {"orgn_domain", 1},
		},
		By:        "orgn_domain",
		StorageID: "_id",
	}

	status, result, err := srv.user.Compare(ctx, srv.storage, comparator, comparable)
	if err != nil {
		logrus.Errorf("<%v>[userService.AreInSameOrgn] could not compare items: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.AreInSameOrgnResponse{StatusCode: int32(status), Msg: "could not compare items"}, nil
	}
	logrus.Warn(result)
	return &userSrv.AreInSameOrgnResponse{
		StatusCode:    int32(status),
		Msg:           "compared items",
		TruthfulValid: result.TruthfulValid,
		MissMatches:   result.MissItems,
	}, nil
}

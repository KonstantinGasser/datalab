package app

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (app app) AppendMember(ctx context.Context, mongo storage.Storage, req *appSrv.AppendMemberRequest) (int, error) {
	// add tracingID from request to context
	ctx = ctx_value.AddValue(ctx, "tracingID", req.GetTracing_ID())
	logrus.Infof("<%v>[appService.GetApps] received get apps request\n", ctx_value.GetString(ctx, "tracingID"))

	memberList := req.GetMember()
	var err error
	for _, member := range memberList {
		// perform upsert on member list for app
		err = mongo.UpdateByID(ctx, dbName, appCollection, req.GetAppUuid(), bson.D{
			{"$addToSet", bson.M{"member": member}},
		})
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

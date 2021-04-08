package api

import (
	"context"
	"net/http"
	"strings"

	appSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/app_service"
	"github.com/KonstantinGasser/datalabs/backend/services/app_service/pkg/app"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/KonstantinGasser/datalabs/utils/hash"
	"github.com/KonstantinGasser/datalabs/utils/unique"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
<<<<<<< HEAD
func (srv AppService) CreateApp(ctx context.Context, in *appSrv.CreateAppRequest) (*appSrv.CreateAppResponse, error) {
=======
func (srv AppService) Create(ctx context.Context, in *appSrv.CreateRequest) (*appSrv.CreateResponse, error) {
>>>>>>> feature_app_token
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.CreateApp] received request\n", ctx_value.GetString(ctx, "tracingID"))

	appUUID, err := unique.UUID()
	if err != nil {
		return &appSrv.CreateResponse{StatusCode: http.StatusInternalServerError, Msg: "could not create app", AppUuid: ""}, nil
	}
	orgnAppHash := hash.Sha256([]byte(strings.Join([]string{in.GetOrganization(), in.GetName()}, "/"))).String()
	status, err := srv.app.CreateApp(ctx, srv.storage, app.AppItem{
<<<<<<< HEAD
		UUID:        appUUID,
		AppName:     in.GetName(),
		Description: in.GetDescription(),
		OwnerUUID:   in.GetOwnerUuid(),
		OrgnDomain:  in.GetOrganization(),
		Member:      append(in.GetMember(), in.GetOwnerUuid()), // by default owner will always be in the app team
		Settings:    in.GetSettings(),
=======
		UUID:           appUUID,
		AppName:        in.GetName(),
		Description:    in.GetDescription(),
		OwnerUUID:      in.GetOwnerUuid(),
		OrgnDomain:     in.GetOrganization(),
		Member:         append(in.GetMember(), in.GetOwnerUuid()), // by default owner will always be in the app team
		Settings:       in.GetSettings(),
		OrgnAndAppHash: orgnAppHash,
>>>>>>> feature_app_token
	})
	if err != nil {
		logrus.Errorf("<%v>[appService.CreateApp] could not create app: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.CreateResponse{StatusCode: int32(status), Msg: "could not create app"}, nil
	}
	return &appSrv.CreateResponse{
		StatusCode: int32(status),
		Msg:        "app has been created",
		AppUuid:    appUUID,
	}, nil
}

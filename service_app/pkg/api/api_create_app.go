package api

import (
	"context"
	"net/http"
	"strings"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	"github.com/KonstantinGasser/datalab/service_app/pkg/app"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/utils/hash"
	"github.com/KonstantinGasser/datalab/utils/unique"
	"github.com/KonstantinGasser/required"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
func (srv AppService) Create(ctx context.Context, in *appSrv.CreateRequest) (*appSrv.CreateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.CreateApp] received request\n", ctx_value.GetString(ctx, "tracingID"))

	// initialize configurations
	respCfg, err := srv.configService.Init(ctx, &configSrv.InitRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
	})
	if err != nil {
		logrus.Errorf("<%v>[appService.CreateApp] could not init config: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.CreateResponse{StatusCode: http.StatusInternalServerError, Msg: "Could not create app", AppUuid: ""}, nil
	}
	appUUID, err := unique.UUID()
	if err != nil {
		return &appSrv.CreateResponse{StatusCode: http.StatusInternalServerError, Msg: "could not create app", AppUuid: ""}, nil
	}
	orgnAppHash := hash.Sha256([]byte(strings.Join([]string{in.GetOrganization(), in.GetName()}, "/"))).String()
	appItem := app.AppItem{
		UUID:           appUUID,
		AppName:        in.GetName(),
		URL:            in.GetAppUrl(),
		Description:    in.GetDescription(),
		OwnerUUID:      in.GetOwnerUuid(),
		OrgnDomain:     in.GetOrganization(),
		Member:         append(in.GetMember(), in.GetOwnerUuid()), // by default owner will always be in the app team
		ConfigRef:      respCfg.GetUUID(),
		OrgnAndAppHash: orgnAppHash,
	}
	if err := required.Atomic(&appItem); err != nil {
		logrus.Errorf("<%v>[appService.CreateApp] %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.CreateResponse{StatusCode: http.StatusBadRequest, Msg: "Mandatory fields missing"}, nil
	}
	if err := srv.app.Create(ctx, srv.storage, appItem); err != nil {
		logrus.Errorf("<%v>[appService.CreateApp] %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &appSrv.CreateResponse{StatusCode: err.Code(), Msg: err.Error()}, nil
	}
	return &appSrv.CreateResponse{
		StatusCode: http.StatusOK,
		Msg:        "app has been created",
		AppUuid:    appUUID,
	}, nil
}

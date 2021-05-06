package api

import (
	"context"
	"net/http"

	apptokenSrv "github.com/KonstantinGasser/datalab/protobuf/apptoken_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppTokenServer) Get(ctx context.Context, in *apptokenSrv.GetRequest) (*apptokenSrv.GetResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.Issue] received request\n", ctx_value.GetString(ctx, "tracingID"))

	token, exp, err := srv.apptoken.Get(ctx, srv.storage, in.GetAppUuid())
	if err != nil {
		return &apptokenSrv.GetResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Token:      "",
			TokenExp:   0,
		}, nil
	}
	return &apptokenSrv.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "App-Token Information",
		Token:      token,
		TokenExp:   exp,
	}, nil
}

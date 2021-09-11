package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppConfigServer) AppendPermission(ctx context.Context, in *proto.AppendPermissionRequest) (*proto.AppendPermissionResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.AppendPermission] received request\n", tracingId)

	err := server.modifyService.AddPermission(ctx, in.GetAppUuid(), in.GetUserUuid())
	if err != nil {
		logrus.Errorf("[%v][server.AppendPermission] could not append member permission: %v\n", tracingId, err)
		return &proto.AppendPermissionResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.AppendPermissionResponse{
		StatusCode: http.StatusOK,
		Msg:        "Appended member permission",
	}, nil
}

func (server AppConfigServer) RollbackAppendPermission(ctx context.Context, in *proto.RollbackAppendPermissionRequest) (*proto.RollbackAppendPermissionResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.RollbackAppendPermission] received request\n", tracingId)

	err := server.modifyService.RollbackPermission(ctx, in.GetAppUuid(), in.GetUserUuid())
	if err != nil {
		logrus.Errorf("[%v][server.RollbackAppendPermission] could not rollback member permission: %v\n", tracingId, err)
		return &proto.RollbackAppendPermissionResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.RollbackAppendPermissionResponse{
		StatusCode: http.StatusOK,
		Msg:        "Permission rollback successful",
	}, nil
}

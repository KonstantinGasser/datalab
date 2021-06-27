package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppTokenServer) AppendPermission(ctx context.Context, in *proto.AppendPermissionRequest) (*proto.AppendPermissionResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.AppendPermission] received request\n", tracingId)

	err := server.modifySevice.AddPermission(ctx, in.GetAppUuid(), in.GetUserUuid())
	if err != nil {
		logrus.Errorf("[%v][server.AppendPermission] could not append App-Token permissions: %v\n", tracingId, err.Error())
		return &proto.AppendPermissionResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.AppendPermissionResponse{
		StatusCode: http.StatusOK,
		Msg:        "App-Token permissions appended",
	}, nil
}

func (server AppTokenServer) RollbackAppendPermission(ctx context.Context, in *proto.RollbackAppendPermissionRequest) (*proto.RollbackAppendPermissionResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.RollbackAppendPermission] received request\n", tracingId)

	err := server.modifySevice.RollbackPermission(ctx, in.GetAppUuid(), in.GetUserUuid())
	if err != nil {
		logrus.Errorf("[%v][server.RollbackAppendPermission] could not rollback App-Token permissions: %v\n", tracingId, err.Error())
		return &proto.RollbackAppendPermissionResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.RollbackAppendPermissionResponse{
		StatusCode: http.StatusOK,
		Msg:        "App-Token permissions appended",
	}, nil
}

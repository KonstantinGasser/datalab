package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server UserMetaServer) GetColleagues(ctx context.Context, in *proto.GetColleaguesRequest) (*proto.GetColleaguesResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.GetColleagues] received request\n", tracingId)

	userList, err := server.fetchService.FetchByOrganization(ctx, in.GetOrganization())
	if err != nil {
		logrus.Errorf("[%v][server.GetColleagues] could not get users: %v\n", tracingId, err.Error())
		return &proto.GetColleaguesResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Colleagues: nil,
		}, nil
	}

	var users = make([]*common.UserInfo, len(userList))
	for i, user := range userList {
		users[i] = &common.UserInfo{
			Uuid:         user.Uuid,
			Username:     user.Username,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			OrgnDomain:   user.Organization,
			OrgnPosition: user.Position,
			Avatar:       user.Avatar,
		}
	}
	return &proto.GetColleaguesResponse{
		StatusCode: http.StatusOK,
		Msg:        "User profile",
		Colleagues: users,
	}, nil
}

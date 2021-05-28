package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	grpcAppToken "github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

type ClientAppToken struct {
	Conn grpcAppToken.AppTokenClient
}

func NewClientAppToken(clientAddr string) (*ClientAppToken, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := grpcAppToken.NewAppTokenClient(conn)
	return &ClientAppToken{
		Conn: client,
	}, nil
}

func (client ClientAppToken) CollectAppToken(ctx context.Context, appUuid string, authedUser *common.AuthedUser, resC chan struct {
	Field string
	Value interface{}
}, errC chan error) {
	resp, err := client.Conn.Get(ctx, &grpcAppToken.GetRequest{
		Tracing_ID: ctx.Value("tracingID").(string),
		AuthedUser: authedUser,
		AppUuid:    appUuid,
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("apptoken.get err: %s", resp.GetMsg())
		return
	}
	resC <- struct {
		Field string
		Value interface{}
	}{
		Field: "apptoken",
		Value: resp.GetToken(),
	}
}

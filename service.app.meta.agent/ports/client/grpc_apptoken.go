package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports/client/intercepter"
	grpcAppToken "github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

type ClientAppToken struct {
	conn grpcAppToken.AppTokenClient
}

func NewClientAppToken(clientAddr string) (*ClientAppToken, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure(), intercepter.WithUnary(intercepter.WithAuth))
	if err != nil {
		return nil, err
	}
	client := grpcAppToken.NewAppTokenClient(conn)
	return &ClientAppToken{
		conn: client,
	}, nil
}

// Emit tirggers the initialize endpoint for the AppTokenSerivce asking it to
// initialize and acknowlege that a new app token with the given reference needs to be
// created
func (client ClientAppToken) Emit(ctx context.Context, event *ports.Event, errC chan error) {
	resp, err := client.conn.Initialize(ctx, &grpcAppToken.InitializeRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppRefUuid: event.App.Uuid,
		AppOwner:   event.App.OwnerUuid,
		AppHash:    event.App.Hash,
		AppOrigin:  event.App.URL,
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("apptoken.init err: %s", resp.GetMsg())
		return
	}
	errC <- nil
}

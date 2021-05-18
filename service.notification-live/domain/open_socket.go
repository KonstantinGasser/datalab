package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/utils/ctx_value"
)

var (
	ErrNotAuthorized = fmt.Errorf("user is not authorized")
)

func (svc notificationlogic) OpenSocket(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims := ctx_value.GetAuthedUser(ctx)
	if claims == nil {
		return ErrNotAuthorized
	}
	svc.notifyHub.OpenSocket(ctx, w, r, claims.Uuid, claims.Organization)
	return nil
}

package ports

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps"
)

type EventEmitter interface {
	EmitInit(ctx context.Context, event *InitEvent, errC chan error)
	EmitAppendPermissions(ctx context.Context, event *PermissionEvent, errC chan error)
	EmitRollbackAppendPermissions(ctx context.Context, event *PermissionEvent, errC chan error)
}

type InitEvent struct {
	App *apps.App
}

func NewInitEvent(app *apps.App) *InitEvent {
	return &InitEvent{App: app}
}

type PermissionEvent struct {
	AppUuid  string
	UserUuid string
}

func NewPermissionEvent(appUuid, userUuid string) *PermissionEvent {
	return &PermissionEvent{
		AppUuid:  appUuid,
		UserUuid: userUuid,
	}
}

package ports

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps"
)

type EventEmitter interface {
	Emit(ctx context.Context, event *Event, errC chan error)
}

type Event struct {
	App *apps.App
}

func NewEvent(app *apps.App) *Event {
	return &Event{App: app}
}

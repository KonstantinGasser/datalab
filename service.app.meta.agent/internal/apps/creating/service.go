package creating

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateDefaultApp(ctx context.Context, name, URL, ownerUuid, ownerOrgn, desc string) (string, errors.Api)
}

type service struct {
	repo            apps.AppsRepository
	emitterAppToken ports.EventEmitter
	emitterAppConf  ports.EventEmitter
}

func NewService(repo apps.AppsRepository, emitterAppToken ports.EventEmitter, emitterAppConf ports.EventEmitter) Service {
	return &service{
		repo:            repo,
		emitterAppToken: emitterAppToken,
		emitterAppConf:  emitterAppConf,
	}
}

func (s service) CreateDefaultApp(ctx context.Context, name, URL, ownerUuid, ownerOrgn, desc string) (string, errors.Api) {
	app, err := apps.NewDefault(name, URL, ownerUuid, ownerOrgn, desc)
	if err != nil {
		return "", errors.New(http.StatusBadRequest, err, "App is missing fields")
	}
	err = s.repo.Store(ctx, *app)
	if err != nil {
		return "", errors.New(http.StatusInternalServerError, err, "Could not create App")
	}
	err = s.emitInitEvent(ctx, app)
	if err != nil {
		return "", errors.New(http.StatusInternalServerError, err, "Could not create App")
	}
	return app.Uuid, nil
}

// emitInitEvent distributes the event that a new app has been created triggering the init endpoints
// of the AppTokenSerivce and AppConfigService
func (s service) emitInitEvent(ctx context.Context, app *apps.App) error {
	withCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	var errC = make(chan error)
	emitterEvent := ports.NewEvent(app)

	go s.emitterAppToken.Emit(withCancel, emitterEvent, errC)
	go s.emitterAppConf.Emit(withCancel, emitterEvent, errC)

	for i := 0; i < 2; i++ {
		err := <-errC
		if err != nil {
			logrus.Errorf("[%s][creating.EmitInit] emit cause error: %v\n", ctx.Value("tracingID"), err)
			// if there is an error while emitting events
			// here the emiited events must succed in order for the
			// transaction to succeed - hence if err cancel context and
			// role back (if that would have been implmeneted)
			return err
		}
	}
	close(errC)
	return nil
}

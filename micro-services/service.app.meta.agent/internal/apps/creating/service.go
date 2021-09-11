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
	CreateDefaultApp(ctx context.Context, name, URL, ownerUuid, ownerOrgn string, tags []string, isPrivate bool) (string, errors.Api)
}

type service struct {
	repo            apps.AppsRepository
	emitterAppToken ports.EventEmitter
	emitterAppConf  ports.EventEmitter
	// emitterUserPermissions ports.EventEmitter
}

func NewService(repo apps.AppsRepository, emitterAppToken ports.EventEmitter,
	emitterAppConf ports.EventEmitter) Service {
	return &service{
		repo:            repo,
		emitterAppToken: emitterAppToken,
		emitterAppConf:  emitterAppConf,
	}
}

func (s service) CreateDefaultApp(ctx context.Context, name, URL, ownerUuid, ownerOrgn string, tags []string, isPrivate bool) (string, errors.Api) {
	app, err := apps.NewDefault(name, URL, ownerUuid, ownerOrgn, tags, isPrivate)
	if err != nil {
		return "", errors.New(http.StatusBadRequest, err, "App is missing fields")
	}
	err = s.repo.Store(ctx, *app)
	if err != nil {
		return "", errors.New(http.StatusInternalServerError, err, "Could not create App")
	}
	err = s.emitInitEvent(ctx, app)
	if err != nil {
		// if transaction fails - rollback app creation
		if err := s.compentsateCreate(ctx, app.Uuid); err != nil {
			logrus.Errorf("[create.Rollback] could not rollback App creation: %v\n", err)
		}
		return "", errors.New(http.StatusInternalServerError, err, "Could not create App")
	}
	return app.Uuid, nil
}

func (s service) compentsateCreate(ctx context.Context, appUuid string) error {
	return s.repo.CompensateCreate(ctx, appUuid)
}

// emitInitEvent distributes the event that a new app has been created triggering the init endpoints
// of the AppTokenService and AppConfigService
func (s service) emitInitEvent(ctx context.Context, app *apps.App) error {
	withCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	var errC = make(chan error)
	emitterEvent := ports.NewInitEvent(app)

	go s.emitterAppToken.EmitInit(withCancel, emitterEvent, errC)
	go s.emitterAppConf.EmitInit(withCancel, emitterEvent, errC)
	// go s.emitterUserPermissions.Emit(withCancel, emitterEvent, errC)

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

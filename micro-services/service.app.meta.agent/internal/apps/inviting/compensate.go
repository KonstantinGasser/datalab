package inviting

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps"
	"github.com/sirupsen/logrus"
)

func (s service) compensateInvite(ctx context.Context, appUuid string, member apps.Member) error {
	logrus.Infof("[inviting.compensateInvite] running rollback on accepted invite member\n")
	return s.repo.CompensateMemberStatus(ctx, appUuid, member)
}

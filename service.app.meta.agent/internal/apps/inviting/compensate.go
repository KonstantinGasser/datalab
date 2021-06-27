package inviting

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps"
)

func (s service) compensateInvite(ctx context.Context, appUuid string, member apps.Member) error {
	return s.repo.CompensateMemberStatus(ctx, appUuid, member)
}

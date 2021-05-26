package permissions

import "context"

type Repository interface {
	HasRWAccess(ctx context.Context, uuid, ownerUuid string) error
}

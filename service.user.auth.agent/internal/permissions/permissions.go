package permissions

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/required"
)

var (
	ErrMissingFields = fmt.Errorf("missing fields to create Permissions")
)

type PermissionRepo interface {
	Store(ctx context.Context, permission Permission) error
	GetById(ctx context.Context, userUuid string, stored interface{}) error
	AddApp(ctx context.Context, userUuid, appUuid string) error
}

type Permission struct {
	UserUuid string   `bson:"_id" required:"yes"`
	UserOrgn string   `bson:"user_orgn" required:"yes"`
	Apps     []string `bson:"apps"`
}

func NewDefaultPermission(userUuid, userOrgn string) (*Permission, error) {
	permission := Permission{
		UserUuid: userUuid,
		UserOrgn: userOrgn,
		Apps:     make([]string, 0),
	}
	if err := required.Atomic(&permission); err != nil {
		return nil, ErrMissingFields
	}
	return &permission, nil
}

// AddApp adds a app permissions the the Permission.Apps slice
func (p *Permission) AddApp(appUuid string) *Permission {
	for _, app := range p.Apps {
		if app == appUuid {
			return p
		}
	}
	p.Apps = append(p.Apps, appUuid)
	return p
}

// AllowedApps yields back all apps listed in permissions
func (p Permission) AllowedApps() []string {
	return p.Apps
}

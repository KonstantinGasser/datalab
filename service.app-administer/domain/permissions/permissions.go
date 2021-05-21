package permissions

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotAuthorized  = fmt.Errorf("caller is not authorized to perform the action")
	ErrNotFound       = fmt.Errorf("could not find app information")
	ErrNoAppHashFound = fmt.Errorf("could not find any app hash for app uuid")
	ErrNoAppsFound    = fmt.Errorf("could not find any apps the user is allowed to have access to")
)

func IsOwner(ctx context.Context, repo repo.Repo, callerUuid, appUuid string) error {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.M{"_id": appUuid},
				bson.M{"owner_uuid": callerUuid},
			},
		},
	}
	ok, err := repo.Exists(ctx, config.AppDB, config.AppColl, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotAuthorized
		}
		return err
	}
	if !ok {
		return ErrNotAuthorized
	}
	return nil
}

func IsMember(ctx context.Context, repo repo.Repo, permissions *common.UserTokenClaims, appUuid string) error {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "_id", Value: appUuid},
				},
				bson.D{
					{
						Key: "$and",
						Value: bson.A{
							bson.D{{Key: "member.uuid", Value: permissions.GetUuid()}},
							bson.D{{Key: "member.status", Value: types.InviteAccepted}},
						},
					},
				},
			},
		},
	}
	ok, err := repo.Exists(ctx, config.AppDB, config.AppColl, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotAuthorized
		}
		return err
	}
	if !ok {
		return ErrNotAuthorized
	}
	return nil
}

func IsOwnerOrMember(ctx context.Context, repo repo.Repo, permission *common.UserTokenClaims, appUuid string) error {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.M{"_id": appUuid},
				bson.D{
					{
						Key: "$or",
						Value: bson.A{
							bson.D{{Key: "owner_uuid", Value: permission.GetUuid()}},
							bson.D{
								{
									Key: "$and",
									Value: bson.A{
										bson.D{{Key: "member.uuid", Value: permission.GetUuid()}},
										bson.D{{Key: "member.status", Value: types.InviteAccepted}},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	ok, err := repo.Exists(ctx, config.AppDB, config.AppColl, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotAuthorized
		}
		return err
	}
	if !ok {
		return ErrNotAuthorized
	}
	return nil
}

func HasInvite(ctx context.Context, repo repo.Repo, permissions *common.UserTokenClaims, appUuid string) error {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "_id", Value: appUuid},
				},
				bson.D{
					{
						Key: "$and",
						Value: bson.A{
							bson.D{{Key: "member.uuid", Value: permissions.GetUuid()}},
							bson.D{{Key: "member.status", Value: types.InvitePending}},
						},
					},
				},
			},
		},
	}
	ok, err := repo.Exists(ctx, config.AppDB, config.AppColl, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotAuthorized
		}
		return err
	}
	if !ok {
		return ErrNotAuthorized
	}
	return nil
}

// IsCorrectHash checks if the provided app hash matches with the database records
// is used to authorize certain action performed on an app
func IsCorrectHash(ctx context.Context, repo repo.Repo, appUuid, hash string) error {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.M{"_id": appUuid},
				bson.M{"app_hash": hash},
			},
		},
	}
	ok, err := repo.Exists(ctx, config.AppDB, config.AppColl, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ErrNotAuthorized
		}
		return err
	}
	// double check if no docs found implies no permissions, right?
	if !ok {
		return ErrNotAuthorized
	}
	return nil
}

// CanAccess looks up all apps the current user has permissions to access
func CanAccess(ctx context.Context, repo repo.Repo, permissions *common.UserTokenClaims) ([]string, error) {
	// filters for all apps where user is either owner or
	// is listed as member in app - returning only the app uuids
	filter := bson.D{
		{
			Key: "$or",
			Value: bson.A{
				bson.D{
					{
						Key:   "owner_uuid",
						Value: permissions.GetUuid(),
					},
				},
				bson.D{
					{
						Key: "$and",
						Value: bson.A{
							bson.D{{Key: "member.uuid", Value: permissions.GetUuid()}},
							bson.D{{Key: "member.status", Value: types.InviteAccepted}},
						},
					},
				},
			},
		},
	}
	var allowedApps []types.AppInfo
	err := repo.FindMany(ctx, config.AppDB, config.AppColl, filter, &allowedApps)
	if err != nil {
		return nil, err
	}
	if len(allowedApps) == 0 || allowedApps == nil {
		return []string{}, nil
	}
	var appUuids = make([]string, len(allowedApps))
	for i, item := range allowedApps {
		appUuids[i] = item.Uuid
	}
	return appUuids, nil
}

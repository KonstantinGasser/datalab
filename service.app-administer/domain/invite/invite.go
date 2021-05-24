package invite

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrNoAppFound = fmt.Errorf("could not find requested app info")
)

func ToApp(ctx context.Context, repo repo.Repo, userUuid, ownerUuid, appUuid string) error {

	var invite = types.Invite{
		Uuid:   userUuid,
		Status: types.InvitePending,
	}
	filter := bson.M{"_id": appUuid, "owner_uuid": ownerUuid}
	updateCount, err := repo.UpdateOne(ctx, config.AppDB, config.AppColl, filter, bson.D{
		{
			Key: "$addToSet",
			Value: bson.M{
				"member": invite,
			},
		},
	}, false)
	if err != nil {
		if updateCount == 0 {
			return ErrNoAppFound
		}
		return err
	}
	return nil
}

func Accept(ctx context.Context, repo repo.Repo, appUuid, userUuid string) error {
	return updateInviteStatus(ctx, repo, appUuid, userUuid, types.InviteAccepted)
}

func Reject(ctx context.Context, repo repo.Repo, appUuid, userUuid string) error {
	return updateInviteStatus(ctx, repo, appUuid, userUuid, types.InviteRejected)
}

func updateInviteStatus(ctx context.Context, repo repo.Repo, appUuid, userUuid string, status types.InviteStatus) error {

	// loop up app where user is listed as member in pending state
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.D{
					{
						Key:   "_id",
						Value: appUuid,
					},
				},
				bson.D{
					{
						Key: "$and",
						Value: bson.A{
							bson.D{{Key: "member.uuid", Value: userUuid}},
							bson.D{{Key: "member.status", Value: types.InvitePending}},
						},
					},
				},
			},
		},
	}
	query := bson.D{
		{
			Key:   "$set",
			Value: bson.M{"member.$.status": status},
		},
	}
	_, err := repo.UpdateOne(ctx, config.AppDB, config.AppColl, filter, query, false)
	return err
}

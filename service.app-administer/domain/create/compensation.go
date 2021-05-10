package create

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"go.mongodb.org/mongo-driver/bson"
)

// CompensateApp rolls back a previous create app by deleting the data from the
// database
func CompensateApp(ctx context.Context, repo repo.Repo, appUuid string) error {
	err := repo.DeleteOne(ctx, config.AppDB, config.AppColl, bson.M{"_id": appUuid})
	if err != nil {
		return err
	}
	return nil
}

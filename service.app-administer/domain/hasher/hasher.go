package hasher

import (
	"context"
	"fmt"
	"strings"

	"github.com/KonstantinGasser/datalab/service.app-administer/config"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/types"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"github.com/KonstantinGasser/datalab/utils/hash"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ErrHashMisMatch = fmt.Errorf("provided app-hash does not match records")
)

// Compare loads the referenced app-hash from the database and compares it with the
// given one
func Compare(ctx context.Context, repo repo.Repo, hash string, appUuid string) error {

	var data types.AppInfo
	err := repo.FindOne(ctx, config.AppDB, config.AppColl, bson.M{"_id": appUuid}, &data)
	if err != nil {
		return err
	}
	if data.AppHash != hash {
		return ErrHashMisMatch
	}
	return nil
}

// Build generates the app hash
func Build(appName string, organization string) string {
	concated := strings.Join([]string{appName, organization}, "/")
	return hash.Sha256([]byte(concated)).String()
}

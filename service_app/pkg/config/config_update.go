package config

import (
	"context"

	"github.com/KonstantinGasser/datalabs/service_app/pkg/storage"
)

// Upsert updates the app configurations based on the updateFlag. The updateFlag if set to "*" will update all passed
// configs for the given app else only the one passed in the updateFlag
func (cfg config) Upsert(ctx context.Context, storage storage.Storage, changes Cfgs, updateFlag string) (int, error) {

}

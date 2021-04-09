package user

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalabs/backend/services/user_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// Update updates the allowed fields of the user record
func (user user) Update(ctx context.Context, storage storage.Storage, userItem UserItemUpdateable) (int, error) {

	updateQuery := bson.D{
		{
			Key:   "$set",
			Value: userItem,
		},
	}
	if err := storage.UpdateOne(ctx, userDatabase, userCollection, bson.M{"_id": userItem.UUID}, updateQuery); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

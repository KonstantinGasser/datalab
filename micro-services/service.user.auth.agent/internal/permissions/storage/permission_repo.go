package storage

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	nameDB   = "datalab_userauth"
	nameColl = "userpermission"
)

// MongoClient implements the apptokens.Repository interface
type MongoClient struct {
	conn *mongo.Client
}

func NewMongoClient(conn *mongo.Client) (*MongoClient, error) {
	return &MongoClient{conn: conn}, nil
}

// InsertOne inserts one data point into the mongo database for a given db name and
// collection name. Query must be any of bson.* or a struct with bson tags
// Returned data from the coll.InsertOne are ignored and will not be returned by the function
func (client MongoClient) Store(ctx context.Context, permission permissions.Permission) error {

	data, err := bson.Marshal(permission)
	if err != nil {
		return err
	}

	coll := client.conn.Database(nameDB).Collection(nameColl)
	_, err = coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// GetById looks up the app  behind the uuid and writes the result in the passed pointer
// to the result. If none found returns mongo.ErrNoDocuments
func (client MongoClient) GetById(ctx context.Context, userUuid string, stored interface{}) error {
	filter := bson.M{"_id": userUuid}

	coll := client.conn.Database(nameDB).Collection(nameColl)
	if err := coll.FindOne(ctx, filter).Decode(stored); err != nil {
		return err
	}
	return nil
}

func (client MongoClient) AddApp(ctx context.Context, userUuid, appUuid string) error {
	query := bson.D{
		{
			Key:   "$addToSet",
			Value: bson.M{"apps": appUuid},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	_, err := coll.UpdateByID(ctx, userUuid, query)
	if err != nil {
		return err
	}
	return nil
}

package storage

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	nameDB   = "datalab_usermeta"
	nameColl = "usermeta"
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
func (client MongoClient) Store(ctx context.Context, user users.User) error {

	data, err := bson.Marshal(user)
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

func (client MongoClient) UpdateUser(ctx context.Context, updatable users.UpdatableUser) error {
	query := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"first_name": updatable.FirstName,
				"last_name":  updatable.LastName,
				"position":   updatable.Position,
			},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	_, err := coll.UpdateByID(ctx, updatable.Uuid, query)
	if err != nil {
		return err
	}
	return nil
}

// GetById looks up the app  behind the uuid and writes the result in the passed pointer
// to the result. If none found returns mongo.ErrNoDocuments
func (client MongoClient) GetById(ctx context.Context, uuid string, stored interface{}) error {
	filter := bson.M{"_id": uuid}

	coll := client.conn.Database(nameDB).Collection(nameColl)
	if err := coll.FindOne(ctx, filter).Decode(stored); err != nil {
		return err
	}
	return nil
}

// GetAll collects all Apps which uuid is in the slice of uuids
func (client MongoClient) GetByOrganization(ctx context.Context, organization string, stored interface{}) error {

	filter := bson.M{"organization": organization}

	coll := client.conn.Database(nameDB).Collection(nameColl)
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}
	if err := cur.All(ctx, stored); err != nil {
		return err
	}
	return nil
}

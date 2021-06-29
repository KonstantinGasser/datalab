package storage

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/libconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	nameDB   = "datalab_appconfig"
	nameColl = "appconfig"
)

// MongoClient implements the apptokens.Repository interface
type MongoClient struct {
	conn *mongo.Client
}

func NewMongoClient(conn *mongo.Client) *MongoClient {
	return &MongoClient{conn: conn}
}

func (client MongoClient) Load(ctx context.Context, appUuid string) (*libconfig.Config, error) {
	filter := bson.M{"_id": appUuid}

	var conf libconfig.Config
	coll := client.conn.Database(nameDB).Collection(nameColl)
	if err := coll.FindOne(ctx, filter).Decode(&conf); err != nil {
		return nil, err
	}
	fmt.Println("Config: ", conf)
	return &conf, nil
}

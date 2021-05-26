package dao

import (
	"context"
	"time"

	"github.com/KonstantinGasser/datalab/service.app-configuration/config"
	"github.com/KonstantinGasser/datalab/service.app-configuration/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient implements the Storage interface and wraps its function to serve as Mongo
// storage option
type MongoClient struct {
	conn *mongo.Client
}

func NewMongoDB(addr string) (*MongoClient, error) {
	opts := options.Client().ApplyURI(addr)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return &MongoClient{conn: conn}, nil
}

func (client MongoClient) InsertInitConfig(ctx context.Context, initConfig domain.ConfigInfo) error {
	coll := client.conn.Database(config.CfgDB).Collection(config.CfgColl)
	_, err := coll.InsertOne(ctx, initConfig)
	if err != nil {
		return err
	}
	return nil
}

// GetById looks up the stored App Config for a given uuid
func (client MongoClient) GetById(ctx context.Context, uuid string) (*domain.ConfigInfo, error) {
	filter := bson.M{"_id": uuid}

	coll := client.conn.Database(config.CfgDB).Collection(config.CfgColl)
	var config domain.ConfigInfo
	if err := coll.FindOne(ctx, filter).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// UodateByFlag updates one config within the AppConfig document, relpacing the old config
func (client MongoClient) UpdateByFlag(ctx context.Context, flag string, uuid string, appConfig []interface{}) error {
	filter := bson.M{"_id": uuid}
	query := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				flag: appConfig,
			},
		},
	}

	coll := client.conn.Database(config.CfgDB).Collection(config.CfgColl)
	if _, err := coll.UpdateOne(ctx, filter, query); err != nil {
		return err
	}
	return nil
}

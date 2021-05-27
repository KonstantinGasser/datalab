package mongo

import (
	"context"
	"time"

	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	nameDB   = "datalab_appconfig"
	nameColl = "appconfig"
)

// MongoClient implements the apptokens.Repository interface
type MongoClient struct {
	conn *mongo.Client
}

func NewMongoClient(addr string) (*MongoClient, error) {
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

// InsertOne inserts one data point into the mongo database for a given db name and
// collection name. Query must be any of bson.* or a struct with bson tags
// Returned data from the coll.InsertOne are ignored and will not be returned by the function
func (client MongoClient) Initialize(ctx context.Context, appConfig appconfigs.AppConfig) error {

	data, err := bson.Marshal(appConfig)
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

// GetById looks up the app token behind the uuid and writes the result in the passed pointer
// // to the result. If none found returns mongo.ErrNoDocuments
func (client MongoClient) GetById(ctx context.Context, uuid string, result interface{}) error {
	filter := bson.M{"_id": uuid}

	coll := client.conn.Database(nameDB).Collection(nameColl)
	if err := coll.FindOne(ctx, filter).Decode(result); err != nil {
		return err
	}
	return nil
}

// UpdateByFlag updates one config of the AppConfig (funnel, campaign, btn_time)
func (client MongoClient) UpdateByFlag(ctx context.Context, uuid, flag string, data interface{}) error {
	query := bson.D{
		{
			Key: "$addToSet",
			Value: bson.M{
				flag: data,
			},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	if _, err := coll.UpdateByID(ctx, uuid, query); err != nil {
		return err
	}
	return nil
}

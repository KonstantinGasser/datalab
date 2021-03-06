package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// time to wait for service to connect to MongoDB
	connectTimeout = time.Second * 10
	collectTimeout = time.Second * 10
	// DB, Collection names
	dbName         = "datalabs_user"
	userCollection = "user"
)

type MongoClient struct {
	conn *mongo.Client
}

// NewMongoClient opens a connection to a given MongoDB and returns a pointer
// to a mongoClient
func NewMongoClient(connString string) (*MongoClient, error) {
	opts := options.Client().ApplyURI(connString)

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()
	conn, err := mongo.Connect(ctx, opts)
	if err != nil {
		logrus.Errorf("[mongo.Conn.Connect] could not connect to mongoDB <%s>: %v\n", connString, err)
		return nil, fmt.Errorf("could not connect to mongoDB <%s>: %v", connString, err)
	}
	// check if connection is alive
	if err := conn.Ping(context.TODO(), nil); err != nil {
		logrus.Errorf("[mongo.Conn.Ping] could not ping <%s>: %v\n", connString, err)
		return nil, fmt.Errorf("could not ping <%s>: %v", connString, err)
	}
	return &MongoClient{
		conn: conn,
	}, nil

}

// InsertOne inserts one data point into the mongo database for a given db name and
// collection name. Returned data from the coll.InsertOne are ignored and will not be returned
// by the function
func (client MongoClient) InsertOne(ctx context.Context, db, collection string, data []byte) error {
	coll := client.conn.Database(db).Collection(collection)
	_, err := coll.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("mongo client, could not execute InsertOne: %v", err)
	}
	return nil
}

// FindOne takes a database and collection name and a bson.M query to find a single result
// returns an error or the result (result can be an empty bson.M map if not found in db/collection)
func (client MongoClient) FindOne(ctx context.Context, db, collection string, data bson.M) (bson.M, error) {
	coll := client.conn.Database(db).Collection(collection)

	var result bson.M
	if err := coll.FindOne(ctx, data).Decode(&result); err != nil {
		// Decode will return ErrNoDocuments if the query returns no result
		// this is less an error but similar to io.EOF and means NoRecoredFound
		if err == mongo.ErrNoDocuments {
			return bson.M{}, nil
		}
		return nil, fmt.Errorf("mongo client, could not decode FindOne result: %v", err)
	}
	return result, nil
}

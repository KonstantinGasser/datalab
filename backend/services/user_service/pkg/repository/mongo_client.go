package repository

import (
	"context"
	"fmt"
	"log"
	"time"

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
		return nil, fmt.Errorf("could not connect to mongoDB <%s>: %v", connString, err)
	}
	// check if connection is alive
	if err := conn.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("could not ping <%s>: %v", connString, err)
	}
	return &MongoClient{
		conn: conn,
	}, nil

}

func (client MongoClient) InsertUser(data MongoUser) error {

	// marshal data to bson
	bdata, err := bson.Marshal(data)
	if err != nil {
		return fmt.Errorf("could not marshal user to bson: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), collectTimeout)
	defer cancel()

	collection := client.conn.Database(dbName).Collection(userCollection)
	result, err := collection.InsertOne(
		context.Background(),
		bdata,
	)
	if err != nil {
		return fmt.Errorf("could not insert user in <%s:%s>:%v", dbName, userCollection, err)
	}
	// just to see what mongo returns
	log.Printf("result from mongo: %v", result)
	return nil
}

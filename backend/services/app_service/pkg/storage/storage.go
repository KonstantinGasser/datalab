package storage

import (
	"context"
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
)

// Storage is the api def to speak with the mongoDB
type Storage interface {
	// FindAll and FindOne both take in a pointer to a struct in which the mongo
	// db will deserialize the found results (must be passed as pointer else no data)
	FindAll(ctx context.Context, dbName, collName string, filter bson.D, result interface{}) error
	FindOne(ctx context.Context, dbName, collName string, filter bson.M, result interface{}) error
	Exsists(ctx context.Context, dbName, collName string, filter bson.M) (bool, error)
	InsertOne(ctx context.Context, dbName, collName string, data []byte) error
	DeleteOne(ctx context.Context, dbName, collName string, filter bson.D) error
	UpdateByID(ctx context.Context, dbName, collName, appUUID string, data bson.D) error
}

// New returns a new mongoDB client implementing the Storage interface
func New(connString string) Storage {
	opts := options.Client().ApplyURI(connString)

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()
	conn, err := mongo.Connect(ctx, opts)
	if err != nil {
		logrus.Errorf("[storage.New] could not connect to mongoDB: %v\n", err)
		return nil
	}
	if err := conn.Ping(context.TODO(), nil); err != nil {
		logrus.Errorf("[storage.New] could not ping mongoDB: %v\n", err)
		return nil
	}
	return &mongoC{conn: conn}
}
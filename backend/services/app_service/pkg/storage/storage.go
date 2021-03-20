package storage

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
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
	FindAll(ctx context.Context, db, collection string, filter, result interface{}) error
	FindOne(ctx context.Context, db, collection string, filter, result interface{}) error
	Exsists(ctx context.Context, db, collection string, filter interface{}) (bool, error)
	InsertOne(ctx context.Context, db, collection string, query interface{}) error
	DeleteOne(ctx context.Context, db, collection string, filter interface{}) error
	UpdateOne(ctx context.Context, db, collection string, filter, query interface{}) error
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
	return &mongoClient{conn: conn}
}

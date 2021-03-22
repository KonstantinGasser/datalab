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
	// InsertOne appends the database with the given data
	InsertOne(ctx context.Context, db, collection string, query interface{}) error
	// UpdateOne updates on document in the database with the given data and based on the given filter
	UpdateOne(ctx context.Context, db, collection string, filter, query interface{}) error
	// FindOne selects zero or one match from the filter an assigns the data in the given pointer to results
	FindOne(ctx context.Context, db, collection string, filter, result interface{}) error
	// FindMany select zero or many matches from the filter an assigns the data in the given pointer to results
	FindMany(ctx context.Context, db, collection string, filter, results interface{}) error
	// Exists checks whether something exists in the storage based on the filter
	Exists(ctx context.Context, db, collection string, filter interface{}) (bool, error)
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

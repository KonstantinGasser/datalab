package storage

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// time to wait for service to connect to MongoDB
	connectTimeout = time.Second * 10
	collectTimeout = time.Second * 10
)

var (
	ErrDocExists    = errors.New("requested document exists")
	ErrInsertFailed = errors.New("\"insert\" of data failed")
	ErrFindeFailed  = errors.New("\"find\" document failed")
	ErrUpdateFailed = errors.New("\"update\" of document failed")
)

// Storage is the api def to speak with the mongoDB
type Storage interface {
	// InsertOne appends the database with the given data
	InsertOne(ctx context.Context, db, collection string, query interface{}) error
	// UpdateOne updates on document in the database with the given data and based on the given filter
	// returns the number of updated documents and an error. Accepts mongo.UpdateOption.Upsert to upsert documents
	UpdateOne(ctx context.Context, db, collection string, filter, query interface{}, upsert bool) (int, error)
	// FindOne selects zero or one match from the filter an assigns the data in the given pointer to results
	FindOne(ctx context.Context, db, collection string, filter, result interface{}) error
	// FindMany select zero or many matches from the filter an assigns the data in the given pointer to results
	FindMany(ctx context.Context, db, collection string, filter, results interface{}) error
	// Exists checks whether something exists in the storage based on the filter
	Exists(ctx context.Context, db, collection string, filter interface{}) (bool, error)
	// Delete deletes a document based on a given filter
	DeleteOne(ctx context.Context, db, collection string, filter interface{}) error
}

// New returns a new mongoDB client implementing the Storage interface
func New(connString string) (Storage, error) {
	opts := options.Client().ApplyURI(connString)

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()
	conn, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	return &mongoClient{conn: conn}, nil
}

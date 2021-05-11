package repo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/KonstantinGasser/datalab/service.app-administer/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoClient implements the Storage interface and wraps its function to serve as Mongo
// storage option
type mongoClient struct {
	conn *mongo.Client
}

func NewMongoDB(addr string) (Repo, error) {
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
	return &mongoClient{conn: conn}, nil
}

// FindAll takes in a query, a db and collection name
// found results will be written in the result interface{} which MUST be a pointer else no data
func (client mongoClient) FindMany(ctx context.Context, db, collection string, filter, result interface{}) error {
	coll := client.conn.Database(db).Collection(collection)

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}
	if err = cur.All(ctx, result); err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}
	return nil
}

// FindOne takes a database and collection name and a bson.M query to find a single result
// found results will be written in the passed result interface{} which MUST be a pointer else no data.
// If no documents are found the function will return a mongo.ErrNoDocuments error indicating that the passed result
// will be empty. Caller must check for this error
func (client mongoClient) FindOne(ctx context.Context, db, collection string, filter, result interface{}) error {
	coll := client.conn.Database(db).Collection(collection)

	if err := coll.FindOne(ctx, filter).Decode(result); err != nil {
		// Decode will return ErrNoDocuments if the query returns no result
		// this is less an error but similar to io.EOF and means NoRecoredFound
		if err == mongo.ErrNoDocuments {
			return errors.ErrAPI{
				Status: http.StatusBadGateway,
				Err:    err,
				Msg:    "Could not find any document for request",
			}
		}
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}
	return nil
}

// InsertOne inserts one data point into the mongo database for a given db name and
// collection name. Query must be any of bson.* or a struct with bson tags
// Returned data from the coll.InsertOne are ignored and will not be returned by the function
func (client mongoClient) InsertOne(ctx context.Context, db, collection string, query interface{}) error {
	// check if interface{} is a struct if so needs marshaling
	var data interface{} = query
	var err error
	if reflect.ValueOf(query).Kind() == reflect.Struct {
		fmt.Println("1")
		data, err = bson.Marshal(query)
		fmt.Println("2")
		if err != nil {
			return err
		}
	}
	fmt.Println("3")
	coll := client.conn.Database(db).Collection(collection)
	_, err = coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOne is a generic function to delete one record in a given database/collection based on a given
// filter - can be any of type bson.*
func (client mongoClient) DeleteOne(ctx context.Context, db, collection string, filter interface{}) error {
	coll := client.conn.Database(db).Collection(collection)

	if _, err := coll.DeleteOne(ctx, filter); err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}
	return nil
}

// UpdateOne is a generic function which updates on record in the given database/collection based on the filter
// filter and query can be any of bson.*. returns the number of modified documents and an error
func (client mongoClient) UpdateOne(ctx context.Context, db, collection string, filter, query interface{}, upsert bool) (int, error) {
	// check if interface{} is a struct if so needs marshaling
	var data interface{} = query
	var err error
	if reflect.ValueOf(query).Kind() == reflect.Struct {
		data, err = bson.Marshal(query)
		if err != nil {
			return 0, errors.ErrAPI{
				Status: http.StatusInternalServerError,
				Err:    err,
			}
		}
	}

	coll := client.conn.Database(db).Collection(collection)
	updated, err := coll.UpdateOne(ctx, filter, data, &options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		return int(updated.ModifiedCount), errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}
	return int(updated.ModifiedCount), nil
}

// Exists checks whether a recored in the given database/collection based on the given filter exists
// if error happens bool will be zero value (false) which does not mean that the record may not exists
func (client mongoClient) Exists(ctx context.Context, db, collection string, filter interface{}) (bool, error) {
	coll := client.conn.Database(db).Collection(collection)

	var records bson.M
	if err := coll.FindOne(ctx, filter).Decode(&records); err != nil {
		// Decode will return ErrNoDocuments if the query returns no result
		// this is less an error but similar to io.EOF and means NoRecoredFound
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
		}
	}
	if len(records) == 0 {
		return false, nil
	}
	return true, nil
}

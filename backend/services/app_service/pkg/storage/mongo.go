package storage

import (
	"context"
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// mongoClient implements the Storage interface and wraps its function to serve as Mongo
// storage option
type mongoClient struct {
	conn *mongo.Client
}

// FindAll takes in a query, a db and collection name
// found results will be written in the result interface{} which MUST be a pointer else no data
func (client mongoClient) FindMany(ctx context.Context, db, collection string, filter, result interface{}) error {
	coll := client.conn.Database(db).Collection(collection)

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}
	if err = cur.All(ctx, result); err != nil {
		return err
	}
	return nil
}

// FindOne takes a database and collection name and a bson.M query to find a single result
// found results will be written in the passed result interface{} which MUST be a pointer else no data lol
func (client mongoClient) FindOne(ctx context.Context, db, collection string, filter, result interface{}) error {
	coll := client.conn.Database(db).Collection(collection)

	if err := coll.FindOne(ctx, filter).Decode(result); err != nil {
		// Decode will return ErrNoDocuments if the query returns no result
		// this is less an error but similar to io.EOF and means NoRecoredFound
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return fmt.Errorf("mongo client, could not decode FindOne result: %v", err)
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
		data, err = bson.Marshal(query)
		if err != nil {
			return err
		}
	}

	coll := client.conn.Database(db).Collection(collection)
	_, err = coll.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("mongo client, could not execute InsertOne: %v", err)
	}
	return nil
}

// DeleteOne is a generic function to delete one record in a given database/collection based on a given
// filter - can be any of type bson.*
func (client mongoClient) DeleteOne(ctx context.Context, db, collection string, filter interface{}) error {
	coll := client.conn.Database(db).Collection(collection)

	if _, err := coll.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

// UpdateOne is a generic function which updates on record in the given database/collection based on the filter
// filter and query can be any of bson.*. returns the number of modified documents and an error
func (client mongoClient) UpdateOne(ctx context.Context, db, collection string, filter, query interface{}) (int, error) {
	// check if interface{} is a struct if so needs marshaling
	var data interface{} = query
	var err error
	if reflect.ValueOf(query).Kind() == reflect.Struct {
		data, err = bson.Marshal(query)
		if err != nil {
			return 0, err
		}
	}

	coll := client.conn.Database(db).Collection(collection)
	updated, err := coll.UpdateOne(ctx, filter, data)
	if err != nil {
		logrus.Error(err)
		return int(updated.ModifiedCount), err
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
		return false, fmt.Errorf("mongo client, could not decode FindOne result: %v", err)
	}
	if len(records) == 0 {
		return false, nil
	}
	return true, nil
}

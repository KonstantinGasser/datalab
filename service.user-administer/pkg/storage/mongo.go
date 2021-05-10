package storage

import (
	"context"
	"fmt"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoClient struct {
	conn *mongo.Client
}

// InsertOne is a generic wrapper function to insert one data point in the MongoDB
// the query interface{} must satisfy the mongo.Client.Database.Collection.InsertOne interface{}
// can be any of bson.* or structs specifying the `bson:"some_field_name"` tag those will be marshaled via bson.Marshal
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
		return err
	}
	return nil
}

// UpdateOne is a generic wrapper for the function to update a document based on a filter
// filter and query must be any of bson.* type or a struct with given bson tags -> will be marshaled in func then
func (client mongoClient) UpdateOne(ctx context.Context, db, collection string, filter, query interface{}) error {
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
	_, err = coll.UpdateOne(ctx, filter, data)
	if err != nil {
		return err
	}
	return nil
}

// FindOne is a generic wrapper to find a document based on a filter (must be of type bson.*), the passed
// in result must be a pointer to an struct in order for it to retrieve data. (similar like json.Decode works)
// can return a mongo.ErrNoDocuments which implies that the passed in result will hold NO data. Caller needs to check
// this error
func (client mongoClient) FindOne(ctx context.Context, db, collection string, filter, result interface{}) error {
	// ensure passed result is a pointer else panic
	if reflect.ValueOf(result).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("passed result must be a pointer, have: %v want: &result", reflect.ValueOf(result).Kind()))
	}
	coll := client.conn.Database(db).Collection(collection)
	if err := coll.FindOne(ctx, filter).Decode(result); err != nil {
		return err
	}
	return nil
}

// FindMany is a generic wrapper to find any number of documents based on the passed filter
// the argument "results" must be a pointer to a slice! in which the results will be written into
func (client mongoClient) FindMany(ctx context.Context, db, collection string, filter, results interface{}) error {
	// ensure passed result is a pointer else panic
	if reflect.ValueOf(results).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("passed result must be a pointer, have: %v want: &result", reflect.ValueOf(results).Kind()))
	}

	coll := client.conn.Database(db).Collection(collection)
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}
	if err := cursor.All(ctx, results); err != nil {
		return err
	}
	return nil
}

// Exists checks based on a filter whether a record lives in the database or not
// if the query to the database fails it will return an error but also the default value of bool (false)
// so watch out if you get an error - which does not mean because the bool is false, that the record does not exits!
func (client mongoClient) Exists(ctx context.Context, db, collection string, filter interface{}) (bool, error) {
	coll := client.conn.Database(db).Collection(collection)

	var found bson.M
	if err := coll.FindOne(ctx, filter).Decode(&found); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err // return false since false is the default value of bool
	}
	// if bson.M has values record found
	if len(found) != 0 {
		return true, nil
	}
	return false, nil
}

package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoC struct {
	conn *mongo.Client
}

// FindAll takes in a query, a db and collection name
// found results will be written in the result interface{} which MUST be a pointer else no data lola
func (client mongoC) FindAll(ctx context.Context, db, collection string, filter bson.D, result interface{}) error {
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
func (client mongoC) FindOne(ctx context.Context, db, collection string, filter bson.M, result interface{}) error {
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
// collection name. Returned data from the coll.InsertOne are ignored and will not be returned
// by the function
func (client mongoC) InsertOne(ctx context.Context, db, collection string, data []byte) error {
	coll := client.conn.Database(db).Collection(collection)
	_, err := coll.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("mongo client, could not execute InsertOne: %v", err)
	}
	return nil
}

func (client mongoC) DeleteOne(ctx context.Context, db, collection string, filter bson.D) error {
	coll := client.conn.Database(db).Collection(collection)

	if _, err := coll.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

func (client mongoC) UpdateByID(ctx context.Context, db, collection, appUUID string, data bson.D) error {
	coll := client.conn.Database(db).Collection(collection)

	if _, err := coll.UpdateByID(ctx, bson.D{{"_id", appUUID}}, data); err != nil {
		return err
	}
	return nil
}

func (client mongoC) Exsists(ctx context.Context, db, collection string, filter bson.M) (bool, error) {
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

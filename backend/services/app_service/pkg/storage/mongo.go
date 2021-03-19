package storage

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoC struct {
	conn *mongo.Client
}

func (client mongoC) FindAll(ctx context.Context, db, collection string, filter bson.D) ([]bson.M, error) {
	coll := client.conn.Database(db).Collection(collection)

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return []bson.M{}, err
	}
	var results []bson.M
	if err = cur.All(ctx, &results); err != nil {
		return []bson.M{}, err
	}
	return results, nil
}

// FindOne takes a database and collection name and a bson.M query to find a single result
// returns an error or the result (result can be an empty bson.M map if not found in db/collection)
func (client mongoC) FindOne(ctx context.Context, db, collection string, data bson.M) (bson.M, error) {
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
	logrus.Infof("Result: %v", result)
	return result, nil
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

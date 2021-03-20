package storage

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// const (
// 	// time to wait for service to connect to MongoDB
// 	connectTimeout = time.Second * 10
// 	collectTimeout = time.Second * 10
// 	// DB, Collection names
// 	dbName         = "datalabs_user"
// 	userCollection = "user"
// )

// type mongoClient struct {
// 	conn *mongo.Client
// //

// // InsertOne inserts one data point into the mongo database for a given db name and
// // collection name. Returned data from the coll.InsertOne are ignored and will not be returned
// // by the function
// func (client MongoClient) InsertOne(ctx context.Context, db, collection string, data []byte) error {
// 	coll := client.conn.Database(db).Collection(collection)
// 	_, err := coll.InsertOne(ctx, data)
// 	if err != nil {
// 		return fmt.Errorf("mongo client, could not execute InsertOne: %v", err)
// 	}
// 	return nil
// }

// // FindOne takes a database and collection name and a bson.M query to find a single result
// // returns an error or the result (result can be an empty bson.M map if not found in db/collection)
// func (client mongoClient) FindOne(ctx context.Context, db, collection string, data bson.M) (bson.M, error) {
// 	coll := client.conn.Database(db).Collection(collection)

// 	var result bson.M
// 	if err := coll.FindOne(ctx, data).Decode(&result); err != nil {
// 		// Decode will return ErrNoDocuments if the query returns no result
// 		// this is less an error but similar to io.EOF and means NoRecoredFound
// 		if err == mongo.ErrNoDocuments {
// 			return bson.M{}, nil
// 		}
// 		return nil, fmt.Errorf("mongo client, could not decode FindOne result: %v", err)
// 	}
// 	return result, nil
// }

// // UpdateOne updates on document for a given database and collection. It searches for the given filter and updates
// // with the passed in data
// func (client MongoClient) UpdateOne(ctx context.Context, db, collection string, filter bson.M, data bson.D) error {
// 	coll := client.conn.Database(db).Collection(collection)

// 	_, err := coll.UpdateOne(ctx, filter, data)
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }

// func (client MongoClient) FindMany(ctx context.Context, db, collection string, filter bson.M, resultSet interface{}) error {
// 	coll := client.conn.Database(db).Collection(collection)

// 	cur, err := coll.Find(ctx, filter)
// 	if err != nil {
// 		return err
// 	}

// 	if err := cur.All(ctx, resultSet); err != nil {
// 		return err
// 	}
// 	return nil
// }

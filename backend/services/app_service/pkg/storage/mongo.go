package storage

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/utils"
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
	logrus.Infof("filter: %v\n cur: %v", filter, cur)
	if err != nil {
		return []bson.M{}, err
	}
	var results []bson.M
	if err = cur.All(ctx, &results); err != nil {
		logrus.Errorf("<%v>[mongo.FindAll] could not unmarshal bson.Raw to bson.M: %v\n", utils.StringValueCtx(ctx, "tracingID"), err)
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
			logrus.Infof("<%v>[mongo.FindOne] could not find any related documents in DB", utils.StringValueCtx(ctx, "tracingID"))
			return bson.M{}, nil
		}
		logrus.Errorf("<%v>[mongo.FindOne], could not decode FindOne result: %v\n", utils.StringValueCtx(ctx, "tracingID"), err)
		return nil, fmt.Errorf("mongo client, could not decode FindOne result: %v", err)
	}
	return result, nil
}

// InsertOne inserts one data point into the mongo database for a given db name and
// collection name. Returned data from the coll.InsertOne are ignored and will not be returned
// by the function
func (client mongoC) InsertOne(ctx context.Context, db, collection string, data []byte) error {
	coll := client.conn.Database(db).Collection(collection)
	_, err := coll.InsertOne(ctx, data)
	if err != nil {
		logrus.Errorf("<%v>[mongo.InsertOne], could not execute InsertOne: %v\n", utils.StringValueCtx(ctx, "tracingID"), err)
		return fmt.Errorf("mongo client, could not execute InsertOne: %v", err)
	}
	return nil
}

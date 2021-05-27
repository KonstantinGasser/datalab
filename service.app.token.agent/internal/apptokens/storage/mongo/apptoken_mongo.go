package mongo

import (
	"context"
	"time"

	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	nameDB   = "datalab_apptoken"
	nameColl = "apptoken"
)

// MongoClient implements the apptokens.Repository interface
type MongoClient struct {
	conn *mongo.Client
}

func NewMongoClient(addr string) (*MongoClient, error) {
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
	return &MongoClient{conn: conn}, nil
}

// InsertOne inserts one data point into the mongo database for a given db name and
// collection name. Query must be any of bson.* or a struct with bson tags
// Returned data from the coll.InsertOne are ignored and will not be returned by the function
func (client MongoClient) Initialize(ctx context.Context, appToken apptokens.AppToken) error {

	data, err := bson.Marshal(appToken)
	if err != nil {
		return err
	}

	coll := client.conn.Database(nameDB).Collection(nameColl)
	_, err = coll.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// GetById looks up the app token behind the uuid and writes the result in the passed pointer
// // to the result. If none found returns mongo.ErrNoDocuments
func (client MongoClient) GetById(ctx context.Context, uuid string, result interface{}) error {
	filter := bson.M{"_id": uuid}

	coll := client.conn.Database(nameDB).Collection(nameColl)
	if err := coll.FindOne(ctx, filter).Decode(result); err != nil {
		return err
	}
	return nil
}

// Update updates the Jwt and Exp of the stored AppToken document
func (client MongoClient) Update(ctx context.Context, uuid, jwt string, exp time.Time) error {
	query := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"app_jwt":     jwt,
				"app_jwt_exp": exp.Unix(),
			},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	if _, err := coll.UpdateByID(ctx, uuid, query); err != nil {
		return err
	}
	return nil
}

// // UpdateMongoClient updates the jwt, expiration time and origin of an existing app token document
// func (client MongoClient) UpdateMongoClient(ctx context.Context, uuid, jwt, origin string, newExp time.Time) error {
// 	filter := bson.M{"_id": uuid}
// 	query := bson.D{
// 		{
// 			Key: "$set",
// 			Value: bson.M{
// 				"app_token":  jwt,
// 				"token_exp":  newExp,
// 				"app_origin": origin,
// 			},
// 		},
// 	}
// 	coll := client.conn.Database(config.TokenDB).Collection(config.TokenColl)
// 	updated, err := coll.UpdateOne(ctx, filter, query)
// 	if err != nil {
// 		return err
// 	}
// 	if int(updated.ModifiedCount) == 0 {
// 		return mongo.ErrNoDocuments
// 	}
// 	return nil
// }

// // HasRWAccess checks if a document where _id = uuid and app_owner = ownerUuid exists
// func (client MongoClient) HasRWAccess(ctx context.Context, uuid, ownerUuid string) error {
// 	filter := bson.D{
// 		{
// 			Key: "$and",
// 			Value: bson.A{
// 				bson.M{"_id": uuid},
// 				bson.M{"app_owner": ownerUuid},
// 			},
// 		},
// 	}
// 	var result bson.M
// 	coll := client.conn.Database(config.TokenDB).Collection(config.TokenColl)
// 	if err := coll.FindOne(ctx, filter).Decode(&result); err != nil {
// 		return err
// 	}
// 	return nil
// }

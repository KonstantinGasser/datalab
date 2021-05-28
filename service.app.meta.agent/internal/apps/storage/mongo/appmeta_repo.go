package mongo

import (
	"context"
	"time"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	nameDB   = "datalab_appmeta"
	nameColl = "appmeta"
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
func (client MongoClient) Store(ctx context.Context, app apps.App) error {

	data, err := bson.Marshal(app)
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

// GetById looks up the app  behind the uuid and writes the result in the passed pointer
// to the result. If none found returns mongo.ErrNoDocuments
func (client MongoClient) GetById(ctx context.Context, uuid string, result interface{}) error {
	filter := bson.M{"_id": uuid}

	coll := client.conn.Database(nameDB).Collection(nameColl)
	if err := coll.FindOne(ctx, filter).Decode(result); err != nil {
		return err
	}
	return nil
}

// GetAll collects all Apps which uuid is in the slice of uuids
func (client MongoClient) GetAll(ctx context.Context, userUuid string, stored interface{}) error {
	// filter all apps where the user is either the owner of the app or listed as member
	// where its status is on InviteAccepted
	filter := bson.D{
		{
			Key: "$or",
			Value: bson.A{
				bson.M{"owner_uuid": userUuid},
				bson.D{
					{
						Key: "$and",
						Value: bson.A{
							bson.M{"member.uuid": userUuid},
							bson.M{"member.status": apps.InviteAccepted},
						},
					},
				},
			},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}
	if err := cur.Decode(stored); err != nil {
		return err
	}
	return nil
}

// AddMember appends the App.Member slice with the given member
func (client MongoClient) AddMember(ctx context.Context, appUuid string, invitedMember apps.Member) error {
	query := bson.D{
		{
			Key: "$addToSet",
			Value: bson.D{
				{
					Key: "member",
					Value: bson.M{
						"uuid":   invitedMember.Uuid,
						"status": invitedMember.Status,
					},
				},
			},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	_, err := coll.UpdateByID(ctx, appUuid, query)
	if err != nil {
		return err
	}
	return nil
}

// MemberStatus updates the status of an member to apps.InviteAccepted
func (client MongoClient) MemberStatus(ctx context.Context, appUuid string, openInvite apps.Member) error {
	// filter for member where uuid AND current (not accepted) status are equal
	// the to openInvite data
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.D{
					{
						Key:   "_id",
						Value: appUuid,
					},
				},
				bson.D{
					{
						Key: "$and",
						Value: bson.A{
							bson.D{{Key: "member.uuid", Value: openInvite.Uuid}},
							bson.D{{Key: "member.status", Value: openInvite.Status}},
						},
					},
				},
			},
		},
	}
	query := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"member.$.status": apps.InviteAccepted,
			},
		},
	}
	coll := client.conn.Database(nameDB).Collection(nameColl)
	_, err := coll.UpdateOne(ctx, filter, query)
	if err != nil {
		return nil
	}
	return nil
}

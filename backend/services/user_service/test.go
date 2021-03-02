package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(connString string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(connString)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := mongo.Connect(ctx, opts)
	if err != nil {
		logrus.Errorf("[mongo.Conn.Connect] could not connect to mongoDB <%s>: %v\n", connString, err)
		return nil, fmt.Errorf("could not connect to mongoDB <%s>: %v", connString, err)
	}
	// check if connection is alive
	if err := conn.Ping(context.TODO(), nil); err != nil {
		logrus.Errorf("[mongo.Conn.Ping] could not ping <%s>: %v\n", connString, err)
		return nil, fmt.Errorf("could not ping <%s>: %v", connString, err)
	}
	return conn, nil
}

func main() {
	c, _ := NewMongoClient("mongodb://userDB:secure@192.168.0.179:27017")
	coll := c.Database("datalabs_user").Collection("user")
	s := time.Now()
	coll.FindOne(context.Background(), bson.M{"username": "tino"})
	log.Printf("passed time: %v", time.Since(s))
}

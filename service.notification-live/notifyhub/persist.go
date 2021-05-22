package notifyhub

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.notification-live/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// HasRecords checks if a user already has a record in the database
func (hub *NotifyHub) HasRecord(userUuid string) (bool, error) {
	filter := bson.M{"_id": userUuid}
	ok, err := hub.repo.Exists(context.Background(), config.NofifyDB, config.NotifyCol, filter)
	if err != nil {
		return false, err
	}
	return ok, nil
}

// PersistInitRecord stores an empty UserNotification if the user
// dose not yet hold a record in the database
func (hub *NotifyHub) PersistInitRecord(conn *Connection) error {
	var userNotifications = UserNotifications{
		UserUuid:      conn.Uuid,
		Organization:  conn.Organization,
		Notifications: []Notification{},
	}
	err := hub.repo.InsertOne(context.Background(), config.NofifyDB, config.NotifyCol, userNotifications)
	if err != nil {
		return err
	}
	return nil
}

func (hub *NotifyHub) SaveEvent(userUuid string, event *IncomingEvent) {
	var notification = Notification{
		Timestamp: event.Timestamp,
		Mutation:  event.Mutation,
		Event:     event.Event,
		Value:     event.Value,
	}
	filter := bson.M{"_id": userUuid}
	query := bson.D{
		{
			Key:   "$addToSet",
			Value: bson.M{"notifications": notification},
		},
	}
	_, err := hub.repo.UpdateOne(context.Background(), config.NofifyDB, config.NotifyCol, filter, query, false)
	if err != nil {
		logrus.Errorf("[notifyhub.SaveEvent] could not save event: %v\n", err)
		return
	}
}

// LookUpAndSend looks up all stored notifications of a given user and
// sends them to the hub.Notify channel
func (hub *NotifyHub) LookUpAndSend(userUuid string) {
	// loop up user records with stored notifications
	filter := bson.M{"_id": userUuid}

	var stored UserNotifications
	err := hub.repo.FindOne(context.Background(), config.NofifyDB, config.NotifyCol, filter, &stored)
	if err != nil {
		// how to handle db error? for no I just log it..sry
		logrus.Errorf("[notifyhub.LoopUpAndSend] could not fetch notifications: %v\n", err)
		return
	}
	// send all messages to client
	hub.batchNotify <- &stored
}

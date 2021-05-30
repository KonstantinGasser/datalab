package notifyhub

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type OrganizationPool []*Connection

var (
	ErrUserInPool      = fmt.Errorf("user already present in pool")
	ErrUserNotFound    = fmt.Errorf("could not find user in pool")
	ErrNoRecevierFound = fmt.Errorf("could not find receiver for message to send to")
	ErrWriteToConn     = fmt.Errorf("could not write to connection")
)

func (pool *OrganizationPool) SendBatch(receiver string, msgs BatchNotification) error {
	for _, conn := range *pool {
		if conn.Uuid == receiver {
			err := conn.Conn.WriteJSON(msgs)
			if err != nil {
				return ErrWriteToConn
			}
			return nil
		}
	}
	return ErrNoRecevierFound
}

func (pool *OrganizationPool) Broadcast(msg *IncomingEvent) error {
	for _, conn := range *pool {
		err := conn.Conn.WriteJSON(msg)
		if err != nil {
			return ErrWriteToConn
		}
	}
	return nil
}

// Send iterate over the pool and sends the message to the connection
// which uuid matches with the receiver ones. If not receiver is found
// Send returns an ErrNoRecevierFound
func (pool *OrganizationPool) Send(receiver string, msg *IncomingEvent) error {
	for _, conn := range *pool {
		if conn.Uuid == receiver {
			err := conn.Conn.WriteJSON(msg)
			if err != nil {
				return ErrWriteToConn
			}
			return nil
		}
	}
	return ErrNoRecevierFound
}

// Add adds a new connection to its pool if the connection
// dose not already exists
func (pool *OrganizationPool) Add(newConn *Connection) error {
	for _, conn := range *pool {
		if newConn.Uuid == conn.Uuid {
			return ErrUserInPool
		}
	}
	*pool = append(*pool, newConn)
	return nil
}

// Remove removes a connection form its pool
func (pool *OrganizationPool) Remove(delConn *Connection) error {
	var offset int
	for i, conn := range *pool {
		if delConn.Uuid == conn.Uuid {
			offset = i
			break
		}
	}
	*pool = append((*pool)[:offset], (*pool)[offset+1:]...)
	return nil
}

func (pool *OrganizationPool) Length() int {
	return len(*pool)
}

func (pool *OrganizationPool) Find(uuid string) *Connection {
	for _, conn := range *pool {
		if conn.Uuid == uuid {
			return conn
		}
	}
	return nil
}

type Connection struct {
	Conn         *websocket.Conn
	Uuid         string
	Organization string
}

// Health checks periodically if the connection is still alive
// if not conn signals to hub that connection can be removed
// since the client should never send any messages Health will block at conn.ReadMessage
// which will interrupt if the connection breaks
func (conn *Connection) Health(hub *NotifyHub) {
	defer func() {
		hub.unsubscribe <- conn
	}()

	// block on ReadMessage
	_, _, err := conn.Conn.ReadMessage()
	if err != nil {
		logrus.Warnf("[connection.Health] closed: %v\n", err)
	}
}

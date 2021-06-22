package connection

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

var (
	ErrEndTimeNotSet = fmt.Errorf("connection end time is not yet set - duration cannot be calculated")
)

type Conn struct {
	PushChan  chan<- map[string]interface{}
	leave     chan<- string
	cookie    string
	wsConn    *websocket.Conn
	startTime time.Time
	endTime   time.Time
}

func New(wsConn *websocket.Conn, pushC chan<- map[string]interface{}, leave chan<- string, cookie string) *Conn {
	return &Conn{
		PushChan:  pushC,
		leave:     leave,
		cookie:    cookie,
		wsConn:    wsConn,
		startTime: time.Now(),
		endTime:   time.Time{}, // careful zero value of time, comparable to nil value for time!
	}
}

func (c Conn) Listen() {
	defer func() {
		c.leave <- c.cookie
	}()

	var msgMap map[string]interface{}
	for {
		err := c.wsConn.ReadJSON(&msgMap)
		if err != nil {
			return
		}
		c.PushChan <- msgMap
	}
}

func (c *Conn) Finish() *Conn {
	c.endTime = time.Now()
	return c
}

// Duration returns the session duration if the endTime is not the default value
func (c Conn) Duration() (time.Duration, error) {
	if c.endTime.IsZero() {
		return 0, ErrEndTimeNotSet
	}
	return c.endTime.Sub(c.startTime), nil
}

package websocket

import (
	"context"

	"github.com/gorilla/websocket"
)

// Client is used for connecting to websocket endpoints.
type Client struct {
	URL      string
	Ctx      context.Context
	c        *websocket.Conn
	Messages chan interface{}
	quit     chan bool
}

// Connect connects to the server.
func (c *Client) Connect() (err error) {
	if c.Ctx == nil {
		c.Ctx = context.Background()
	}
	if c.c, _, err = websocket.DefaultDialer.DialContext(c.Ctx, c.URL, nil); err != nil {
		return
	}
	if err = c.c.WriteJSON(map[string]string{
		"action": "subscribe",
		"topic":  "confirmation",
	}); err != nil {
		return
	}
	c.Messages = make(chan interface{})
	c.quit = make(chan bool)
	go c.loop()
	return
}

// Close closes the connection.
func (c *Client) Close() (err error) {
	err = c.c.Close()
	c.quit <- true
	return
}

func (c *Client) loop() {
	defer close(c.Messages)
	for {
		var m message
		if err := c.c.ReadJSON(&m); err != nil {
			c.c.Close()
			<-c.quit
			return
		}
		select {
		case c.Messages <- m.m:
		case <-c.quit:
			return
		}
	}
}
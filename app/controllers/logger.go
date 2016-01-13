package controllers

import (
	"github.com/revel/revel"
	"github.com/umireon/wsdr_hive/app/logger"
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
	"time"
)

type Logger struct {
	*revel.Controller
}

func (c Logger) Monitor() revel.Result {
	return c.Render()
}

func (c Logger) MonitorWS(user string, ws *websocket.Conn) revel.Result {
	ctx, cancel := context.WithCancel(context.Background())
	newMessage := make(chan string)
	go func() {
		for {
			var msg string
			err := websocket.Message.Receive(ws, &msg)
			if err != nil || msg == "close" {
				revel.INFO.Println("An existing logger connection closed", user)
				cancel()
				return
			}
			newMessage <- msg
		}
	}()

	newEvent := logger.Subscribe()
	defer logger.Cancel(newEvent)

	for {
		var err error
		func() {
			ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			select {
			case <-ctx.Done():
				if ctx.Err() == context.DeadlineExceeded {
					err = websocket.Message.Send(ws, "")
					revel.TRACE.Printf("Keepalive: %p", ws)
				} else {
					err = ctx.Err()
				}
			case e := <-newEvent:
				err = websocket.JSON.Send(ws, e.Message)
			case msg := <-newMessage:
				revel.INFO.Println(msg)
			}
		}()

		if err != nil {
			revel.INFO.Println("Connection closed")
			revel.TRACE.Printf("Closed connection is: %p\n", ws)
			return nil
		}
	}
}

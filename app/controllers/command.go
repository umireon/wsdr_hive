package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/umireon/wsdr_hive/app/activator"
	"golang.org/x/net/websocket"
	"time"
)

var ch chan struct{} = make(chan struct{})

type Command struct {
	*revel.Controller
}

func (c Command) Index() revel.Result {
	return c.Render()
}

func (c Command) Activate() revel.Result {
	activator.Activate()
	return c.RenderJson(123)
}

func (c Command) Activity(user string, ws *websocket.Conn) revel.Result {
	fmt.Println("OK")

	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			if websocket.Message.Receive(ws, &msg) != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	for {
		fmt.Println("inbound")
		var err error
		select {
		case <-ch:
			ws.SetDeadline(time.Now().Add(1 * time.Second))
			err = websocket.JSON.Send(ws, "activate")
			if err == nil {
				var i int
				err = websocket.JSON.Receive(ws, &i)
				fmt.Println(i)
			}
		}
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
}

package logger

import (
	"container/list"
	"github.com/revel/revel"
)

type Event struct {
	Type    string
	Message string
}

func Subscribe() <-chan Event {
	ch := make(chan Event)
	subscribe <- ch
	return ch
}

func Cancel(ch <-chan Event) {
	unsubscribe <- ch
}

func Info(msg string) {
	e := Event{Type: "INFO", Message: msg}
	publish <- e
}

var (
	subscribe   = make(chan (chan Event))
	unsubscribe = make(chan (<-chan Event))
	publish     = make(chan Event)
)

func logger() {
	subscribers := list.New()

	for {
		revel.TRACE.Printf("Listening to logger event...")
		select {
		case ch := <-subscribe:
			revel.TRACE.Println("A new channel subscribing")
			revel.TRACE.Printf("%d channel(s) subscribing\n", subscribers.Len())
			subscribers.PushBack(ch)
		case unsub := <-unsubscribe:
			revel.TRACE.Println("An existing channel unsubscribing")
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				if ch.Value.(chan Event) == unsub {
					subscribers.Remove(ch)
					break
				}
			}
		case e := <-publish:
			revel.TRACE.Println("Publishing event")
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				revel.TRACE.Println(ch)
				ch.Value.(chan Event) <- e
			}
		}
	}
}

func init() {
	go func() {
		ch := Subscribe()
		for e := range ch {
			revel.INFO.Println(e.Message)
		}
	}()

	go logger()
}

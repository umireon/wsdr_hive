package activator

import (
	"container/list"
	"github.com/revel/revel"
)

type Event struct {
	Type    string
	Message string
}

func Join() <-chan Event {
	ch := make(chan Event)
	subscribe <- ch
	return ch
}

func Cancel(ch <-chan Event) {
	unsubscribe <- ch
}

func Activate(msg string) {
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
		revel.TRACE.Println("Listening to an event...")
		select {
		case ch := <-subscribe:
			revel.INFO.Println("Joining a new core")
			subscribers.PushBack(ch)
		case unsub := <-unsubscribe:
			revel.INFO.Println("Canceled joining a core")
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				if ch.Value.(chan Event) == unsub {
					subscribers.Remove(ch)
					break
				}
			}
		case e := <-publish:
			revel.INFO.Println("Activating waiting cores")
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				ch.Value.(chan Event) <- e
			}
			subscribers.Init()
		}
	}
}

func init() {
	go logger()
}

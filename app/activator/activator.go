package activator

import (
	"container/list"
)

type Event struct{}

type Subscription <-chan Event

func Activate() {
}

func Join() Subscription {
	ch := make(Subscription)
	subscribe <- ch
	return ch
}

var (
	subscribe = make(chan Subscription)
	publish   = make(chan Event)
)

func activator() {
	subscribers := list.New()
	for {
		select {
		case ch := <-subscribe:
			subscribers.PushBack(ch)
		case event := <-publish:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				ch.Value.(chan<- Event) <- event
			}
		}
	}
}

func init() {
	go activator()
}

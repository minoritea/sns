package pubsub

import (
	"sync"
)

type PubSub[P any, S any] struct {
	sync.Mutex
	subscribers map[chan P]S
}

func New[P any, S any]() *PubSub[P, S] {
	var pubsub PubSub[P, S]
	pubsub.subscribers = make(map[chan P]S)
	return &pubsub
}

func (p *PubSub[P, S]) Subscribe(subscriber S) (chan P, func()) {
	ch := make(chan P)
	p.Lock()
	defer p.Unlock()
	p.subscribers[ch] = subscriber
	return ch, func() {
		p.Lock()
		defer p.Unlock()
		delete(p.subscribers, ch)
	}
}

func (p *PubSub[P, S]) Publish(payload P, filters ...(func(S) bool)) {
	p.Lock()
	defer p.Unlock()
	for ch, subscriber := range p.subscribers {
		for _, filter := range filters {
			if !filter(subscriber) {
				continue
			}
		}
		ch <- payload
	}
}

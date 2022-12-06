package pubsub

import (
	"sync"
)

type PubSub[P any, S comparable] struct {
	sync.RWMutex
	subscribers map[S]map[chan P]struct{}
}

func New[P any, S comparable]() *PubSub[P, S] {
	var pubsub PubSub[P, S]
	pubsub.subscribers = make(map[S]map[chan P]struct{})
	return &pubsub
}

func (p *PubSub[P, S]) Subscribe(subscriber S) (chan P, func()) {
	ch := make(chan P)
	p.Lock()
	defer p.Unlock()
	if p.subscribers[subscriber] == nil {
		p.subscribers[subscriber] = make(map[chan P]struct{})
	}
	p.subscribers[subscriber][ch] = struct{}{}
	return ch, func() {
		p.Lock()
		defer p.Unlock()
		delete(p.subscribers[subscriber], ch)
	}
}

func (p *PubSub[P, S]) publishToAll(payload P) {
	for _, channels := range p.subscribers {
		for ch := range channels {
			ch <- payload
		}
	}
}

func (p *PubSub[P, S]) publishToSubscribers(payload P, subscribers []S) {
	for _, subscriber := range subscribers {
		channels := p.subscribers[subscriber]
		if channels != nil {
			for ch := range channels {
				ch <- payload
			}
		}
	}
}

func (p *PubSub[P, S]) Publish(payload P, subscribers ...S) {
	p.RLock()
	defer p.RUnlock()

	if len(subscribers) == 0 {
		p.publishToAll(payload)
	} else {
		p.publishToSubscribers(payload, subscribers)
	}
}

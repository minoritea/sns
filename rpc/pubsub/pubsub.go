package pubsub

import (
	"sync"
)

type PubSub[T comparable] struct {
	sync.Mutex
	subscribers map[chan T]struct{}
}

func New[T comparable]() *PubSub[T] {
	var pubsub PubSub[T]
	pubsub.subscribers = make(map[chan T]struct{})
	return &pubsub
}

func (p *PubSub[T]) Subscribe() (chan T, func()) {
	ch := make(chan T)
	p.Lock()
	defer p.Unlock()
	p.subscribers[ch] = struct{}{}
	return ch, func() {
		p.Lock()
		defer p.Unlock()
		delete(p.subscribers, ch)
	}
}

func (p *PubSub[T]) Publish(payload T) {
	p.Lock()
	defer p.Unlock()
	for ch := range p.subscribers {
		ch <- payload
	}
}

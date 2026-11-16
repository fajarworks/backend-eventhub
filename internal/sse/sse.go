package sse

import "sync"

type Broker struct {
	mu   sync.RWMutex
	subs map[int]map[chan int]struct{}
}

func NewBroker() *Broker {
	return &Broker{
		subs: make(map[int]map[chan int]struct{}),
	}

}

func (b *Broker) Subscribe(eventId int) chan int {
	ch := make(chan int, 4)
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.subs[eventId] == nil {
		b.subs[eventId] = make(map[chan int]struct{})
		b.subs[eventId][ch] = struct{}{}
	}
	return ch

}

func (b *Broker) Unsubscribe(eventId int, ch chan int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.subs[eventId], ch)
	if len(b.subs[eventId]) == 0 {
		delete(b.subs, eventId)
	}
}

func (b *Broker) Publish(evenId, count int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for ch := range b.subs[evenId] {
		select {
		case ch <- count:
		default:
		}
	}

}

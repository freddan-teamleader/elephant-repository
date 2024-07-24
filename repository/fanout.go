package repository

import (
	"context"
	"sync"
)

type FanOut[T any] struct {
	m         sync.RWMutex
	listeners map[chan T]func(v T) bool
}

func NewFanOut[T any]() *FanOut[T] {
	return &FanOut[T]{
		listeners: make(map[chan T]func(v T) bool),
	}
}

func (f *FanOut[T]) Listen(ctx context.Context, l chan T, test func(v T) bool) {
	f.m.Lock()
	f.listeners[l] = test
	f.m.Unlock()

	go func() {
		<-ctx.Done()
		f.m.Lock()
		delete(f.listeners, l)
		f.m.Unlock()
	}()
}

func (f *FanOut[T]) Notify(msg T) {
	f.m.RLock()
	listeners := make([]chan T, 0, len(f.listeners))
	tests := make([]func(v T) bool, 0, len(f.listeners))

	for listener, test := range f.listeners {
		listeners = append(listeners, listener)
		tests = append(tests, test)
	}
	f.m.RUnlock()

	for i, listener := range listeners {
		if !tests[i](msg) {
			continue
		}

		select {
		case listener <- msg:
		default:
		}
	}
}

package events

import (
	"sync"
	"time"
)

// teste
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

// Operacoes / Handlres
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}

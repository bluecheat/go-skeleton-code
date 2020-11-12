package event

import "skeleton-code/config"

type EventType int

const (
	ChanEvent EventType = iota
	KafkaEvent
)

type Eventable interface {
	ReadMessage() (interface{}, error)
	WriteMessage(interface{}) error
}

func NewEvent(etype EventType, config config.Config) Eventable {
	switch etype {
	case ChanEvent:
		return newChanEvent()
	case KafkaEvent:
	}
	return nil
}

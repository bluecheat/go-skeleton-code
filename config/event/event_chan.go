package event

import (
	"errors"
	"reflect"
)

type ChanReceiver struct {
	mtype  reflect.Type
	mqueue chan interface{}
}

var (
	typeError = errors.New("read message type is different")
)

func newChanEvent() *ChanReceiver {
	return &ChanReceiver{
		mqueue: make(chan interface{}, 1),
	}
}

func (receiver *ChanReceiver) ReadMessage() (interface{}, error) {
	message := <-receiver.mqueue
	if receiver.mtype != reflect.TypeOf(message) {
		return nil, typeError
	}
	return message, nil
}

func (receiver *ChanReceiver) WriteMessage(message interface{}) error {
	go func() {
		receiver.mqueue <- message
	}()
	return nil
}

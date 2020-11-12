package utils

import (
	"errors"
	"reflect"
	"sync"
)

var errBadChannel = errors.New("event: Subscribe argument does not have sendable channel type")

type Feed struct {
	clients   []reflect.SelectCase
	removeSub chan interface{} // interrupts Send
	feedType  reflect.Type

	once sync.Once
	mu   sync.Mutex
}

const firstSubSendCase = 1

type feedTypeError struct {
	got, want reflect.Type
	op        string
}

func (e feedTypeError) Error() string {
	return "event: wrong type in " + e.op + " got " + e.got.String() + ", want " + e.want.String()
}

func (f *Feed) init() {
	f.removeSub = make(chan interface{})
	f.clients = []reflect.SelectCase{{Chan: reflect.ValueOf(f.removeSub), Dir: reflect.SelectRecv}}
}

func (f *Feed) Subscribe(channel interface{}) {
	f.once.Do(f.init)

	chanval := reflect.ValueOf(channel)
	chantyp := chanval.Type()
	if chantyp.Kind() != reflect.Chan || chantyp.ChanDir()&reflect.SendDir == 0 {
		panic(errBadChannel)
	}
	if !f.typeCheck(chantyp.Elem()) {
		panic(feedTypeError{op: "Subscribe", got: chantyp, want: reflect.ChanOf(reflect.SendDir, f.feedType)})
	}
	cas := reflect.SelectCase{Dir: reflect.SelectSend, Chan: chanval}
	f.mu.Lock()
	defer f.mu.Unlock()
	f.clients = append(f.clients, cas)
}

func (f *Feed) Send(value interface{}) {
	f.once.Do(f.init)

	rvalue := reflect.ValueOf(value)

	if !f.typeCheck(rvalue.Type()) {
		panic(feedTypeError{op: "Send", got: rvalue.Type(), want: reflect.ChanOf(reflect.SendDir, f.feedType)})
	}

	f.mu.Lock()
	clients := f.clients[firstSubSendCase:]
	f.mu.Unlock()

	for i, _ := range clients {
		clients[i].Send = rvalue
	}

	sent := make(chan interface{})

	for {
		chosen, _, _ := reflect.Select(clients)
		clients = f.consume(clients, chosen)
		if len(clients) == 0 {
			go func() {
				sent <- struct{}{}
			}()
			return
		}
	}
}

func (f *Feed) typeCheck(typ reflect.Type) bool {
	if f.feedType == nil {
		f.feedType = typ
		return true
	}
	return f.feedType == typ
}

func (f *Feed) consume(base []reflect.SelectCase, index int) []reflect.SelectCase {
	lastIndex := len(base) - 1
	base[index], base[lastIndex] = base[lastIndex], base[index]
	return base[:lastIndex]
}

package utils

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type testMessage struct {
	msg  string
	time int
}

type failMessage struct {
	msg  string
	time int
}

func TestFeed_SubscribeAndSend(t *testing.T) {
	messageChannel := make(chan testMessage)
	messageChannel2 := make(chan testMessage)
	messageChannel3 := make(chan testMessage)

	var testMessageFeed Feed

	wg := sync.WaitGroup{}
	wg.Add(12)
	go func() {
		for {
			select {
			case msg := <-messageChannel:
				fmt.Println("1 - ", msg)
				wg.Done()
			case msg2 := <-messageChannel2:
				fmt.Println("2 - ", msg2)
				wg.Done()
			case msg3 := <-messageChannel3:
				fmt.Println("3-", msg3)
				wg.Done()
			}
		}
	}()

	go testMessageFeed.Subscribe(messageChannel)
	go testMessageFeed.Subscribe(messageChannel2)
	go testMessageFeed.Subscribe(messageChannel3)
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				testMessageFeed.Send(testMessage{
					msg:  "hello world",
					time: time.Now().Second(),
				})
			}

		}
	}()
	wg.Wait()
}

func TestFeed_SubscribeTypeFail(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			t.Log(v)
			return
		}
	}()

	messageChannel := make(chan testMessage)
	differentChannel := make(chan failMessage)

	var testMessageFeed Feed

	testMessageFeed.Subscribe(messageChannel)
	testMessageFeed.Subscribe(differentChannel)

	t.Error("allowed an subscribe type error")
}

func TestFeed_SendTypeFail(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			t.Log(v)
			return
		}
	}()

	messageChannel := make(chan testMessage)
	messageChannel2 := make(chan testMessage)

	var testMessageFeed Feed

	testMessageFeed.Subscribe(messageChannel)
	testMessageFeed.Subscribe(messageChannel2)

	testMessageFeed.Send(failMessage{
		msg:  "error",
		time: time.Now().Second(),
	})
	t.Error("allowed an send type error")
}

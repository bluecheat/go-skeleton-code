package utils

import (
	"os"
	"os/signal"
	"skeleton-code/logger"
	"syscall"
)

type closeCallback func() error

type Closer struct {
	signChan chan os.Signal
	callback []closeCallback
}

func NewCloser() *Closer {
	scall := make(chan os.Signal, 2)
	signal.Notify(scall, os.Interrupt, os.Kill, syscall.SIGTERM)
	c := &Closer{
		signChan: scall,
		callback: make([]closeCallback, 0),
	}
	go c.watch()
	return c
}

func (c *Closer) watch() {
	for {
		select {
		case <-c.signChan:
			for _, callback := range c.callback {
				if err := callback(); err != nil {
					logger.Error(err)
				}
			}
		}
	}

}

func (c *Closer) Callback(closer closeCallback) {
	logger.Debug("closed ")
	c.callback = append(c.callback, closer)
}

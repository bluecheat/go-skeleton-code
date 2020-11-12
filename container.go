package app

import (
	"context"
	"go.uber.org/fx"
	"sync"
)

type Container interface {
	Invoke(in ...interface{}) *DIContainer
	Provide(pv ...interface{}) *DIContainer
	Module(opts ...fx.Option) *DIContainer
	Run() <-chan bool
	Stop()
}

type DIContainer struct {
	ctx context.Context

	opts []fx.Option
	app  *fx.App

	sync.Mutex
	running chan bool
}

func NewDIContainer() *DIContainer {
	return &DIContainer{
		ctx:     context.Background(),
		opts:    make([]fx.Option, 0),
		running: make(chan bool),
	}
}

func (c *DIContainer) Invoke(in ...interface{}) *DIContainer {
	c.Lock()
	c.opts = append(c.opts, fx.Invoke(in...))
	c.Unlock()
	return c
}

func (c *DIContainer) Provide(pv ...interface{}) *DIContainer {
	c.Lock()
	c.opts = append(c.opts, fx.Provide(pv...))
	c.Unlock()
	return c
}

func (c *DIContainer) Module(opts ...fx.Option) *DIContainer {
	c.Lock()
	for _, opt := range opts {
		c.opts = append(c.opts, opt)
	}
	c.Unlock()
	return c
}

func (c *DIContainer) Run() <-chan bool {
	c.Lock()
	c.app = fx.New(
		c.opts...,
	)
	c.app.Run()
	go func() {
		for {
			select {
			case <-c.app.Done():
				c.Stop()
			}
		}
	}()
	c.Unlock()
	return c.running
}

func (c *DIContainer) Stop() {
	c.app.Stop(c.ctx)
	close(c.running)
}

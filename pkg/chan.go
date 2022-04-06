package pkg

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type chanPool struct {
	EventsTimerChan        chan int64
	LiveInterval           chan string
	SubscribeLogsChan      chan types.Log
	LoggerChan             chan interface{}
}

var Chan = &chanPool{
	EventsTimerChan: make(chan int64,1024),
	LiveInterval: make(chan string, 1024),
	SubscribeLogsChan: make(chan types.Log),
	LoggerChan: make(chan interface{},1024),
}

func (c *chanPool) InEventsTimerChan (ch int64) {
	c.EventsTimerChan <- ch
}

func (c *chanPool) InLiveInterval (ch string) {
	c.LiveInterval <- ch
}

func (c *chanPool) InLoggerChan (logger interface{}) {
	c.LoggerChan <- logger
}
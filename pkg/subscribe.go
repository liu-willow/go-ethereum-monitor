package pkg

import (
	"context"
	"github.com/ethereum/go-ethereum"
)

type subscribe struct {
	*server
	*config
}

func NewSubscribe(client *server,config *config) *subscribe {
	return &subscribe{
		server:    client,
		config:    config,
	}
}

func (s *subscribe) Run() (ethereum.Subscription, error) {
	return s.EthClient.SubscribeFilterLogs(
		context.Background(),
		ethereum.FilterQuery{
			Addresses: s.addresses,
			Topics:    s.topics,
		},
		Chan.SubscribeLogsChan)
}
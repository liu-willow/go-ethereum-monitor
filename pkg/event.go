package pkg

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type event struct {
	*server
	*config
	FromBlock *big.Int
	ToBlock   *big.Int
}

func NewEvent(client *server,config *config) *event {
	return &event{
		server:    client,
		config:    config,
		FromBlock: nil,
		ToBlock:   nil,
	}
}

func (e *event) SetFromBlockWithInt64(from int64) {
	e.FromBlock = big.NewInt(from)
}
func (e *event) SetFromBlockWithBigInt(from *big.Int) {
	e.FromBlock = from
}
func (e *event) SetToBlockWithInt64(to int64) {
	e.ToBlock = big.NewInt(to)
}
func (e *event) SetToBlockWithBigInt(to *big.Int) {
	e.ToBlock = to
}
func (e *event) Run() ([]types.Log, error) {
	eventQuery := ethereum.FilterQuery{
		FromBlock: e.FromBlock,
		ToBlock:   e.ToBlock,
		Addresses: e.config.addresses,
		Topics:    e.config.topics,
	}

	return e.EthClient.FilterLogs(context.Background(), eventQuery)
}
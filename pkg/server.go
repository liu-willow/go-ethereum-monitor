package pkg

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type server struct {
	url       string
	EthClient *ethclient.Client
}

var s *server

func NewServer(url string) *server {
	client, err := ethclient.DialContext(context.Background(),url)
	if err != nil {
		panic(err)
	}
	s = &server{
		url:       url,
		EthClient: client,
	}
	return s
}
func GetServer() *server {
	if s == nil {
		panic("Server Not Ready")
	}
	return s
}

func (cs *server) GetBlock (blockNumber *big.Int) (*types.Block,error) {
	return cs.EthClient.BlockByNumber(context.Background(), blockNumber)
}

func (cs *server) GetBlockByHash (hash common.Hash) (*types.Block,error) {
	return cs.EthClient.BlockByHash(context.Background(),hash)
}

func (cs *server) LastBlock () (int64,error) {
	blockNumber, err := cs.EthClient.BlockNumber(context.Background())
	return int64(blockNumber),err
}

func (cs *server) GetNetworkId () (*big.Int,error) {
	return cs.EthClient.NetworkID(context.Background())
}
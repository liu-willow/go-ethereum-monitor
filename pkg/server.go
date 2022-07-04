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
	ChanId    *big.Int
}

var s *server

func NewServer(url string) *server {
	if s == nil {
		client, err := ethclient.DialContext(context.Background(), url)
		if err != nil {
			panic(err)
		}
		chainId, err := client.NetworkID(context.TODO())
		if err != nil {
			panic(err)
		}
		s = &server{
			url:       url,
			EthClient: client,
			ChanId:    chainId,
		}
	}
	return s
}

func GetServer() *server {
	if s == nil {
		panic("Server Not Ready")
	}
	return s
}

func (cs *server) GetBlock(blockNumber *big.Int) (*types.Block, error) {
	return cs.EthClient.BlockByNumber(context.TODO(), blockNumber)
}

func (cs *server) GetBlockByHash(hash common.Hash) (*types.Block, error) {
	return cs.EthClient.BlockByHash(context.TODO(), hash)
}

func (cs *server) LastBlock() (int64, error) {
	blockNumber, err := cs.EthClient.BlockNumber(context.TODO())
	return int64(blockNumber), err
}

func (cs *server) GetNetworkId() *big.Int {
	return cs.ChanId
}

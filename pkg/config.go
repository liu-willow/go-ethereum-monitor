package pkg

import (
	"github.com/ethereum/go-ethereum/common"
)

type config struct {
	addresses []common.Address
	topics    [][]common.Hash
}

func NewConfig() *config {
	return &config{
		addresses: make([]common.Address, 0),
		topics:    make([][]common.Hash, 0),
	}
}
func (c *config) AddAddressWithAddress(address common.Address) {
	c.addresses = append(c.addresses, address)
}

func (c *config) AddAddressWithString(address string) {
	c.addresses = append(c.addresses, common.HexToAddress(address))
}

func (c *config) AddAddressWithAddresses(addresses []common.Address) {
	c.addresses = addresses
}

func (c *config) AddTopicWithHash(topic common.Hash) {
	if len(c.topics) < 1 {
		c.topics = append(c.topics, []common.Hash{topic})
	} else {
		c.topics[0] = append(c.topics[0], topic)
	}
}

func (c *config) AddTopicWithHashes(topics []common.Hash) {
	if len(c.topics) < 1 {
		c.topics = append(c.topics, topics)
	} else {
		c.topics[0] = append(c.topics[0], topics...)
	}
}

func (c *config) AddTopicWithString(topic string) {
	if len(c.topics) < 1 {
		c.topics = append(c.topics, []common.Hash{common.HexToHash(topic)})
	} else {
		c.topics[0] = append(c.topics[0], common.HexToHash(topic))
	}
}

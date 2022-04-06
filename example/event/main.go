package main

import (
	"encoding/json"
	"fmt"
	"github.com/liu-willow/go-ethereum-monitor/pkg"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"sync"
	"time"
)

const blockHeight = 50

var mu = new(sync.RWMutex)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("err: [%+v]", err)
		}
	}()

	//url can use both https and wss
	url := "wss://speedy-nodes-nyc.moralis.io/0ed910502ea1998707783a43/bsc/testnet/ws"
	ethClient := pkg.NewServer(url)

	netWorkId, err := ethClient.GetNetworkId()
	if err != nil {
		panic(err)
	}
	pkg.Chan.InLoggerChan(fmt.Sprintf("netWorkId: [%+v]", netWorkId.Int64()))
	config := pkg.NewConfig()
	config.AddAddressWithString(os.Getenv("MarketplaceAddress"))

	event := pkg.NewEvent(ethClient, config)

	file := fmt.Sprintf("event/%s", "bsc.json")
	ticker := time.NewTicker(3 * time.Second)
	go func(tc *time.Ticker) {
		for range tc.C {
			configFile := getJsonFile(file)

			lastBlock, _ := event.LastBlock()

			if lastBlock > configFile.BlockNumber {
				startBlock := configFile.BlockNumber
				endBlock := startBlock + blockHeight
				if lastBlock <= endBlock {
					endBlock = lastBlock
				}
				go func(start, end int64) {
					defer func() {
						if e := recover(); e != nil {
							log.Printf("event recover:%+v", e)
							pkg.Chan.InEventsTimerChan(start)
						}
					}()

					event.SetFromBlockWithInt64(start)
					event.SetToBlockWithBigInt(big.NewInt(end))

					logs, err := event.Run()
					if err != nil {
						panic(err)
					}
					for _, v := range logs {
						pkg.Chan.InLoggerChan(fmt.Sprintf("logs:[%+v]", v))
					}

					pkg.Chan.InEventsTimerChan(end)
				}(startBlock, endBlock)
			}
		}
	}(ticker)

	log.Println("event service is starting ...")

	for {
		select {
		case endBlock := <-pkg.Chan.EventsTimerChan:
			content := FileConfig{BlockNumber: endBlock}
			writeToJson(file, content)
		case logger := <-pkg.Chan.LoggerChan:
			log.Printf("logger:%+v\r\n", logger)
		}
	}

}

type FileConfig struct {
	BlockNumber int64 `json:"blockNumber"`
}

func getJsonFile(filePath string) FileConfig {
	mu.RLock()
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("file not exist!")
	}
	u := FileConfig{}
	err = json.Unmarshal(bytes, &u)
	if err != nil {
		panic("parse json config Fail!")
	}
	mu.RUnlock()
	return u
}

func writeToJson(file string, content interface{}) {
	mu.Lock()
	data, err := json.MarshalIndent(content, "", "	")
	if err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile(file, data, 0777); err != nil {
		panic(err)
	}
	mu.Unlock()
}

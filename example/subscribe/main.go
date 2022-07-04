package main

import (
	"fmt"
	"github.com/liu-willow/go-ethereum-monitor/pkg"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	lastTime int64
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("err: [%+v]", err)
		}
	}()

	//url can only use wss
	url := "wss://speedy-nodes-nyc.moralis.io/0ed910502ea1998707783a43/bsc/testnet/ws"
	ethClient := pkg.NewServer(url)

	// //timer Similar to heartbeat
	ticker := time.NewTicker(7 * time.Second)
	go func() {
		for range ticker.C {
			lastBlock, _ := ethClient.LastBlock()
			pkg.Chan.InLoggerChan(fmt.Sprintf("lastBlock: [%+v]", lastBlock))
		}
	}()
	config := pkg.NewConfig()

	/**
	 * set contract address
	 */
	//config.AddAddressWithAddress(common.HexToAddress("0x000000000000000000000000000000000000000000000"))
	config.AddAddressWithString("0x000000000000000000000000000000000000000000000")

	subscribe := pkg.NewSubscribe(ethClient, config)

	sub, err := subscribe.Run()
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	pkg.Chan.InLoggerChan(fmt.Sprintf("netWorkId: [%+v]", subscribe.GetNetworkId().Int64()))

	interval, err := strconv.ParseInt(os.Getenv("LIVEINTERVAL"), 10, 64)
	if err != nil {
		interval = 900
	}
	interval = 9
	tickerLog := time.NewTicker(time.Duration(interval) * time.Second)
	go func(t *time.Ticker) {
		for range t.C {
			now := time.Now().Unix()
			if now-lastTime > interval {
				t.Stop()
				pkg.Chan.InLiveInterval(fmt.Sprintf("revice last log before %d Second", interval))
			}
		}
	}(tickerLog)
	pkg.Chan.InLoggerChan("subscribe service is starting ...")

	for {
		select {
		case err := <-sub.Err():
			ticker.Stop()
			tickerLog.Stop()
			panic(err)
		case vLog := <-pkg.Chan.SubscribeLogsChan:
			lastTime = time.Now().Unix()
			pkg.Chan.InLoggerChan(fmt.Sprintf("vLog: [%+v]", vLog))
		case msg := <-pkg.Chan.LiveInterval:
			ticker.Stop()
			tickerLog.Stop()
			pkg.Chan.LiveInterval = nil
			sub.Unsubscribe()
			panic(msg)
		case logger := <-pkg.Chan.LoggerChan:
			log.Printf("logger:%+v\r\n", logger)
		}
	}
}

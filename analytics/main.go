package main

import (
	cmap "github.com/orcaman/concurrent-map/v2"
	"log"
	"time"
)

// Provider -> token -> operation -> timestamped value, ex. Binance -> USDT -> Buy -> 80.0 23032023 20:00.00
type DataType = map[string]map[string]cmap.ConcurrentMap[string, Order]

var data DataType
var globalProviders ProvidersWeb

func main() {
	var err error
	globalProviders, err = getProviders()
	log.Println("Got providers")
	if err != nil {
		panic(err)
	}
	UpdateData()
	go func() {
		for {
			UpdateData()
			time.Sleep(time.Duration(5) * time.Second)
		}
	}()
	log.Println("Started data updates")
	time.Sleep(time.Duration(10) * time.Second)
	for {
		AnalyzeData()
		time.Sleep(time.Duration(20) * time.Second)
	}
}

package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/tebeka/selenium"
	"net/http"
	"time"
)

var ordersByPageHandle cmap.ConcurrentMap[string, Orders]

func Parse(driver selenium.WebDriver, pages Pages) {
	for {
		for _, page := range pages {
			if err := driver.SwitchWindow(page.handle); err != nil {
				panic(err)
			}
			var orders Orders
			var err error
			if orders, err = page.parse(driver); err != nil {
				println(err) // todo log
			}
			ordersByPageHandle.Set(page.handle, orders)
		}
		time.Sleep(time.Duration(3) * time.Second)
	}
}

func main() {
	ordersByPageHandle = cmap.New[Orders]()
	service, driver := initSelenium()
	defer func(service *selenium.Service) {
		err := service.Stop()
		if err != nil {
			panic(err) // its prob dangerous but whatever
		}
	}(service)
	currencies := []string{
		"USDT",
		//"BTC",
		//"ETH",
	}
	pages, err := initProviders(driver, Providers{
		//Binance{currencies: currencies},
		Garantex{currencies: currencies},
	})
	if err != nil {
		panic(err)
	}
	go Parse(driver, pages)
	http.HandleFunc("/", HandleFunc)
	if err := http.ListenAndServe(":"+"8080", nil); err != nil {
		fmt.Printf("Can't serve err: %v", err)
		return
	}
}

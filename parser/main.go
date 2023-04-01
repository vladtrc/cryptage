package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/tebeka/selenium"
	"log"
	"net/http"
	"time"
)

var ordersByPageHandle cmap.ConcurrentMap[string, Orders]
var handlesByProvider map[string][]string

func Parse(driver selenium.WebDriver, pages Pages) {
	for {
		for _, page := range pages {
			time.Sleep(time.Duration(5) * time.Second)
			if err := driver.SwitchWindow(page.handle); err != nil {
				panic(err)
			}
			var orders Orders
			var err error
			if orders, err = page.parse(driver); err != nil {
				log.Printf("could not parse page:%s", err)
			}
			ordersByPageHandle.Set(page.handle, orders)
		}
		log.Println("Parsed all pages")
	}
}

func main() {
	log.Println("Started parser")
	ordersByPageHandle = cmap.New[Orders]()
	handlesByProvider = make(map[string][]string)
	driver := initSelenium()
	log.Println("Initialized selenium")
	pages, err := initProviders(driver, Providers{
		Binance{currencies: config.currencies},
		Garantex{currencies: config.currencies},
	})
	if err != nil {
		panic(err)
	}
	go Parse(driver, pages)
	http.HandleFunc("/", RouteFunc)
	if err := http.ListenAndServe(":"+config.port, nil); err != nil {
		fmt.Printf("Can't serve err: %v", err)
		return
	}
}

package main

import (
	"fmt"
	"github.com/tebeka/selenium"
)

type Binance struct {
	currencies []string
}

func (b Binance) name() string {
	return "Binance"
}

type OrderTypeCurrency struct {
	OrderType
	string
}

func (o OrderTypeCurrency) BinanceParseOrders(driver selenium.WebDriver) (res Orders, err error) {
	if res, err = parseBinanceP2PTrade(driver); err != nil {
		return
	}
	for i, order := range res {
		order._type = o.OrderType
		order._currency = o.string
		res[i] = order
	}
	return
}
func (b Binance) init(driver selenium.WebDriver) (res Pages, err error) {
	orderTypeTemplates := map[OrderType]string{
		Sell: "https://p2p.binance.com/en/trade/all-payments/%s?fiat=RUB",
		Buy:  "https://p2p.binance.com/en/trade/sell/%s?fiat=RUB&payment=ALL",
	}
	for _, orderType := range []OrderType{Sell, Buy} {
		for _, currency := range b.currencies {
			url := fmt.Sprintf(orderTypeTemplates[orderType], currency)
			var handle string
			if handle, err = createNewTabAndSetCurrent(driver); err != nil {
				return
			}
			if err = driver.Get(url); err != nil {
				return
			}
			if err = prepareBinancePage(driver); err != nil {
				return
			}
			page := Page{
				handle: handle,
				parse:  OrderTypeCurrency{orderType, currency}.BinanceParseOrders,
			}
			res = append(res, page)
		}
	}
	return
}

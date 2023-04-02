package main

import (
	"github.com/tebeka/selenium"
)

type Garantex struct {
	tokens []string
}

func (g Garantex) name() string {
	return "Garantex"
}

type OrderCurrency struct {
	tokens []string
}

func (c OrderCurrency) GarantexParse(driver selenium.WebDriver) (res Orders, err error) {
	for _, orderType := range []OrderType{Sell, Buy} {
		for _, token := range c.tokens {
			var orders Orders
			if orders, err = GarantexParsePage(driver, orderType, token); err != nil {
				return
			}
			res = append(res, orders...)
		}
	}
	return
}

func (g Garantex) init(driver selenium.WebDriver) (res Pages, err error) {
	urlTemplate := "https://garantex.io/trading/usdtrub?lang=en"
	var handle string
	if handle, err = createNewTabAndSetCurrent(driver); err != nil {
		return
	}
	if err = driver.Get(urlTemplate); err != nil {
		return
	}
	page := Page{
		handle: handle,
		parse:  OrderCurrency{g.tokens}.GarantexParse,
	}
	res = append(res, page)
	return
}

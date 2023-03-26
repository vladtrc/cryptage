package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"strings"
)

type Garantex struct {
	currencies []string
}

type OrderCurrency struct {
	string
}

func (c OrderCurrency) GarantexParse(driver selenium.WebDriver) (res Orders, err error) {
	for _, orderType := range []OrderType{Sell, Buy} {
		if res, err = GarantexParsePage(driver, orderType, c.string); err != nil {
			return
		}
	}
	return
}

func (g Garantex) init(driver selenium.WebDriver) (res Pages, err error) {
	urlTemplate := "https://garantex.io/trading/%srub?lang=en"
	for _, currency := range g.currencies {
		var handle string
		if handle, err = createNewTabAndSetCurrent(driver); err != nil {
			return
		}
		if err = driver.Get(fmt.Sprintf(urlTemplate, strings.ToLower(currency))); err != nil {
			return
		}
		page := Page{
			handle: handle,
			parse:  OrderCurrency{currency}.GarantexParse,
		}
		res = append(res, page)
	}
	return
}

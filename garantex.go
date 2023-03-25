package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"strings"
)

type Garantex struct {
	currencies []string
}

func GarantexParsePage(driver selenium.WebDriver, orderType OrderType) (res Orders, err error) {
	var template string
	if orderType == Sell {
		template = "(//th[contains(text(), '(')])[1]/../../../../../div/table/tbody/tr/td[contains(@class, '%s')]/*"
	} else {
		template = "(//th[contains(text(), '(')])[4]/../../../../../div/table/tbody/tr/td[contains(@class, '%s')]/*"
	}
	var prices []string
	var volumes []string
	var pFactors []string
	var amounts []string
	if prices, err = scrapElementTexts(driver, fmt.Sprintf(template, "price")); err != nil {
		return
	}
	if volumes, err = scrapElementTexts(driver, fmt.Sprintf(template, "volume")); err != nil {
		return
	}
	if pFactors, err = scrapElementTexts(driver, fmt.Sprintf(template, "pfactor")); err != nil {
		return
	}
	if amounts, err = scrapElementTexts(driver, fmt.Sprintf(template, "amount")); err != nil {
		return
	}
	for i, price := range prices {
		order := Order{
			price: price,
		}
		if len(volumes) > i {
			order.available = volumes[i]
		}
		if len(pFactors) > i {
			order.commission = pFactors[i]
		}
		if len(amounts) > i {
			order.available += " " + amounts[i]
		}
		res = append(res)
	}
	return
}

type OrderCurrency struct {
	string
}

func (c OrderCurrency) GarantexParse(driver selenium.WebDriver) (res Orders, err error) {
	for _, orderType := range []OrderType{Sell, Buy} {
		if res, err = GarantexParsePage(driver, orderType); err != nil {
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

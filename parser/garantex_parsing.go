package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

func GarantexParsePage(driver selenium.WebDriver, orderType OrderType, currency string) (res Orders, err error) {
	var template string
	if orderType == Buy {
		template = "(//div[contains(text(), 'Buy %s')])[last()]/..//td[contains(@class, '%s')]"
	} else {
		template = "(//div[contains(text(), 'Sell %s')])[last()]/..//td[contains(@class, '%s')]"
	}
	var prices []string
	var volumes []string
	var pFactors []string
	var amounts []string
	if prices, err = scrapElementTexts(driver, fmt.Sprintf(template, currency, "price")); err != nil {
		return
	}
	if volumes, err = scrapElementTexts(driver, fmt.Sprintf(template, currency, "volume")); err != nil {
		return
	}
	if pFactors, err = scrapElementTexts(driver, fmt.Sprintf(template, currency, "pfactor")); err != nil {
		return
	}
	if amounts, err = scrapElementTexts(driver, fmt.Sprintf(template, currency, "amount")); err != nil {
		return
	}
	for i, price := range prices {
		order := Order{
			price:     price,
			_currency: currency,
			_type:     orderType,
			timestamp: time.Now(),
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
		res = append(res, order)
	}
	return
}

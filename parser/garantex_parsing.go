package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"strings"
	"time"
)

func GarantexParsePage(driver selenium.WebDriver, orderType OrderType, currency string) (res Orders, err error) {
	err = xpathForeach(driver, fmt.Sprintf("//a[contains(text(),'%s-RUB')]", currency), func(e selenium.WebElement) (err error) {
		return e.Click()
	})
	if err != nil {
		return
	}
	lwCurrency := strings.ToLower(currency)
	var template string
	if orderType == Buy {
		template = fmt.Sprintf("//tbody[contains(@class, '%srub_bid')]//td[contains(@class, '%s')]", lwCurrency, "%s")
	} else {
		template = fmt.Sprintf("//tbody[contains(@class, '%srub_ask')]//td[contains(@class, '%s')]", lwCurrency, "%s")
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

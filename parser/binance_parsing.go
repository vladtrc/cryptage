package main

import (
	"github.com/tebeka/selenium"
	"log"
	"time"
)

type ByValueCondition struct {
	by       string
	value    string
	expected bool // are we waiting for a button to be visible (true) or otherwise (false)
}

func (c ByValueCondition) checkButtonVisibility(driver selenium.WebDriver) (res bool, err error) {
	we, _ := driver.FindElement(c.by, c.value)
	if we == nil {
		return
	}
	res, _ = we.IsDisplayed()
	if !c.expected {
		res = !res
	}
	return
}
func prepareBinancePage(driver selenium.WebDriver) (err error) {
	acceptAllCookiesXPath := "//button[text()='Accept All Cookies']"
	if err = driver.WaitWithTimeout(ByValueCondition{by: selenium.ByXPATH, value: acceptAllCookiesXPath, expected: true}.checkButtonVisibility, time.Duration(10)*time.Second); err != nil {
		log.Printf("could not wait until the accept cookies button: %s", err)
	}
	err = xpathForeach(driver, acceptAllCookiesXPath, func(e selenium.WebElement) (err error) {
		return e.Click()
	})
	err = xpathForeach(driver, "//input[@placeholder='Enter amount']", func(e selenium.WebElement) (err error) {
		return e.SendKeys("50000" + selenium.ReturnKey)
	})
	if err != nil {
		return
	}
	err = xpathForeach(driver, "//button[@id='C2CofferList_btn_refresh']", func(e selenium.WebElement) (err error) {
		return e.SendKeys(selenium.EnterKey)
	})
	if err != nil {
		return
	}
	err = xpathForeach(driver, "//*[normalize-space(text())='5s to refresh']", func(e selenium.WebElement) (err error) {
		return e.Click()
	})
	return
}
func parseBinanceP2PTrade(driver selenium.WebDriver) (res Orders, err error) {
	pricesPath := "//div[count(./div) = 2]/div[normalize-space(text())='RUB']/../div[not(contains(text(), 'RUB'))]"
	advertisersPath := "//div[count(./div) = 2]/div[normalize-space(text())='RUB']/../../../../div/div[count(./div) = 3]"
	availablePath := "//div[count(./div) = 2]/div[normalize-space(text())='RUB']/../../../../div/div[count(./div) = 2]"
	paymentsPath := "//div[count(./div) = 2]/div[normalize-space(text())='RUB']/../../../../div/div/div/a/../../.."
	var prices []string
	var advertisers []string
	var available []string
	var payments []string
	if prices, err = scrapElementTexts(driver, pricesPath); err != nil {
		return
	}
	if advertisers, err = scrapElementTexts(driver, advertisersPath); err != nil {
		return
	}
	if available, err = scrapElementTexts(driver, availablePath); err != nil {
		return
	}
	if payments, err = scrapElementTexts(driver, paymentsPath); err != nil {
		return
	}
	for i, price := range prices {
		order := Order{
			price: price,
		}
		if len(advertisers) > i {
			order.advertiser = advertisers[i]
		}
		if len(available) > i {
			order.available = available[i]
		}
		if len(payments) > i {
			order.payment = payments[i]
		}
		res = append(res, order)
	}
	return
}

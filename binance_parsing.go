package main

import "github.com/tebeka/selenium"

func xpathForeach(driver selenium.WebDriver, xpath string, handler func(selenium.WebElement) error) (err error) {
	elements, err := driver.FindElements(selenium.ByXPATH, xpath)
	if err != nil {
		return
	}
	for _, e := range elements {
		if err = handler(e); err != nil {
			return
		}
	}
	return
}
func scrapElementTexts(driver selenium.WebDriver, xpath string) (res []string, err error) {
	elements, err := driver.FindElements(selenium.ByXPATH, xpath)
	if err != nil {
		return
	}
	for _, e := range elements {
		var text string
		findElements, _ := e.FindElements(selenium.ByXPATH, "*")
		for er, el := range findElements {
			s, _ := el.Text()
			println(s)
			println(er)
		}
		text, err = e.Text()
		if err != nil {
			return
		}
		res = append(res, text)
	}
	return
}

func prepareBinancePage(driver selenium.WebDriver) (err error) {
	err = xpathForeach(driver, "//input[@placeholder='Enter amount']", func(e selenium.WebElement) (err error) {
		return e.SendKeys("50000")
	})
	if err != nil {
		return
	}
	err = xpathForeach(driver, "//button[text()='Accept All Cookies']", func(e selenium.WebElement) (err error) {
		return e.SendKeys(selenium.EnterKey)
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

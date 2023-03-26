package main

import (
	"github.com/tebeka/selenium"
	"strings"
	"unicode"
)

func removeSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

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
		text, err = e.GetAttribute("innerText")
		text = removeSpaces(text)
		if err != nil {
			return
		}
		res = append(res, text)
	}
	return
}

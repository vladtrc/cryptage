package main

import (
	"github.com/tebeka/selenium"
	"strings"
)

type Page struct {
	handle   string
	provider string
	parse    func(driver selenium.WebDriver) (Orders, error)
}

func (p Page) compare(rhs Page) int {
	return strings.Compare(p.handle, rhs.handle)
}

type Pages []Page

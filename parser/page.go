package main

import (
	"github.com/tebeka/selenium"
)

type Page struct {
	handle string
	parse  func(driver selenium.WebDriver) (Orders, error)
}

type Pages []Page

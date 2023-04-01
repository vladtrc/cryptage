package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"log"
)

type Provider interface {
	init(driver selenium.WebDriver) (Pages, error)
	name() string
}

type Providers []Provider

func initProviders(driver selenium.WebDriver, providers Providers) (pages Pages, err error) {
	for _, provider := range providers {
		var providerPages Pages
		if providerPages, err = provider.init(driver); err != nil {
			return
		}
		for _, page := range providerPages {
			handlesByProvider[provider.name()] = append(handlesByProvider[provider.name()], page.handle)
		}
		pages = append(pages, providerPages...)
		log.Println(fmt.Sprintf("Initialized provider %s", provider.name()))
	}
	return
}

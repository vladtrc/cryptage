package main

import "github.com/tebeka/selenium"

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
			providerByPageHandle[page.handle] = provider.name()
		}
		pages = append(pages, providerPages...)
	}
	return
}

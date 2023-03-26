package main

type Config struct {
	chromeDriverPath string
	reloadSeconds    int
	headless         bool
	port             string
	currencies       []string
}

var config = Config{
	chromeDriverPath: "chromedriver",
	reloadSeconds:    10,
	headless:         true,
	port:             "9090",
	currencies: []string{
		"USDT",
		"BTC",
		"ETH",
	},
}

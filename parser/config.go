package main

type Config struct {
	chromeUrlPrefix string
	chromeArgs      []string
	reloadSeconds   int
	port            string
	currencies      []string
}

var config = Config{
	chromeUrlPrefix: "http://chrome:4444/wd/hub",
	chromeArgs: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		"--headless",
	},
	reloadSeconds: 10,
	port:          "9090",
	currencies: []string{
		"USDT",
		"BTC",
		"ETH",
	},
}

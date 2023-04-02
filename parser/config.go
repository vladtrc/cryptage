package main

type Config struct {
	chromeUrlPrefix   string
	localChromeDriver bool
	chromeArgs        []string
	reloadSeconds     int
	port              string
	tokens            []string
}

var dockerConfig = Config{
	chromeUrlPrefix:   "http://chrome:4444/wd/hub",
	localChromeDriver: false,
	chromeArgs: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		"--headless",
	},
	reloadSeconds: 10,
	port:          "9090",
	tokens: []string{
		"USDT",
		//"BTC",
		"ETH",
	},
}
var localConfig = Config{
	chromeUrlPrefix:   "http://127.0.0.1:4444/wd/hub",
	localChromeDriver: true,
	chromeArgs: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		//"--headless",
	},
	reloadSeconds: 10,
	port:          "9090",
	tokens: []string{
		"USDT",
		"BTC",
		"ETH",
	},
}

var config = dockerConfig

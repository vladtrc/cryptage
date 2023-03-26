package main

type Config struct {
	chromeDriverPath string
	reloadSeconds    int
	headless         bool
}

var config = Config{
	chromeDriverPath: "driver/chromedriver",
	reloadSeconds:    10,
	headless:         true,
}

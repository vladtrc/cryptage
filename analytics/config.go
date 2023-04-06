package main

type Config struct {
	parserURL           string
	tgBotURL            string
	minThresholdPercent float64
}

var localConfig = Config{
	parserURL:           "http://localhost:9090",
	tgBotURL:            "http://localhost:9095",
	minThresholdPercent: 1.5,
}
var dockerConfig = Config{
	parserURL:           "http://parser:9090",
	tgBotURL:            "http://tgbot:9095",
	minThresholdPercent: 1.5,
}
var config = dockerConfig

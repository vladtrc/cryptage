package main

type Config struct {
	parserURL string
	tgBotURL  string
}

var localConfig = Config{
	parserURL: "http://localhost:9090",
	tgBotURL:  "http://localhost:9095",
}
var dockerConfig = Config{
	parserURL: "http://parser:9090",
	tgBotURL:  "http://tgbot:9095",
}
var config = dockerConfig

package main

type Config struct {
	parserURL string
}

var localConfig = Config{
	parserURL: "http://localhost:9090",
}
var dockerConfig = Config{
	parserURL: "http://parser:9090",
}
var config = dockerConfig

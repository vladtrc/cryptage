package main

import (
	"fmt"
	"net/http"
)

var bot TgBot

func main() {
	bot = NewTgBot()
	go bot.Updates()
	http.HandleFunc("/", RouteFunc)
	if err := http.ListenAndServe(":"+config.port, nil); err != nil {
		fmt.Printf("Can't serve err: %v", err)
		return
	}
}

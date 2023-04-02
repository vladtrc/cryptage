package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func TgBotBroadcast(message string) (err error) {
	resource := "/broadcast"
	params := url.Values{}
	params.Add("message", message)
	u, _ := url.ParseRequestURI(config.tgBotURL)
	u.Path = resource
	u.RawQuery = params.Encode()
	resp, err := http.Get(fmt.Sprintf("%v", u))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

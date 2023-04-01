package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type OrderPageWeb struct {
	Price      string `json:"price"`
	Advertiser string `json:"advertiser"`
	Available  string `json:"available"`
	Payment    string `json:"payment"`
	Commission string `json:"commission"`
	Timestamp  string `json:"timestamp"`
}

type OrdersWeb = []OrderPageWeb
type ProviderWeb struct {
	Provider string   `json:"provider"`
	Tokens   []string `json:"currencies"`
}
type ProvidersWeb []ProviderWeb

var parserClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := parserClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
func getProviders() (providers ProvidersWeb, err error) {
	err = getJson(config.parserURL+"/providers", &providers)
	return
}
func getOrders(provider string, token string, op string) (orders OrdersWeb, err error) {
	route := fmt.Sprintf("/%s/%s/%s", provider, token, op)
	err = getJson(config.parserURL+route, &orders)
	return
}

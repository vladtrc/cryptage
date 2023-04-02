package main

import (
	"encoding/json"
	"net/http"
	"strings"
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

type OrdersPageWeb []OrderPageWeb
type ProviderWeb struct {
	Provider   string   `json:"provider"`
	Currencies []string `json:"tokens"`
}
type ProvidersWeb []ProviderWeb

func Ok(res any, w http.ResponseWriter) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	if err := enc.Encode(res); err != nil {
		println(err) // todo log
	}
	w.WriteHeader(http.StatusOK)
}

func GetProviders(w http.ResponseWriter) {
	var providers []string
	for provider, _ := range handlesByProvider {
		providers = append(providers, provider)
	}
	currencies := config.tokens
	var res ProvidersWeb
	for _, p := range providers {
		res = append(res, ProviderWeb{
			Provider:   p,
			Currencies: currencies,
		})
	}
	Ok(res, w)
}
func GetOrders(provider string, currency string, _type string, w http.ResponseWriter) {
	handles := handlesByProvider[provider]
	if len(handles) == 0 {
		w.Write([]byte("no such provider"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var orders Orders
	for _, handle := range handles {
		pageOrders, ok := ordersByPageHandle.Get(handle)
		if !ok {
			w.Write([]byte("invalid handle or parsing has not been done yet"))
			w.WriteHeader(http.StatusNoContent)
			return
		}
		orders = append(orders, pageOrders...)
	}
	var res OrdersPageWeb
	for _, order := range orders {
		orderType := OrderType(_type)
		if order._currency == currency && order._type == orderType {
			orderWeb := OrderPageWeb{
				Price:      order.price,
				Advertiser: order.advertiser,
				Available:  order.available,
				Payment:    order.payment,
				Commission: order.commission,
				Timestamp:  order.timestamp.Format(time.RFC3339),
			}
			res = append(res, orderWeb)
		}
	}
	Ok(res, w)
}
func RouteFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("only get supported"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	trimmedUrl := strings.TrimPrefix(r.URL.Path, "/")
	params := strings.Split(trimmedUrl, "/")
	if len(params) == 3 {
		GetOrders(params[0], params[1], params[2], w)
	} else if len(params) == 1 && params[0] == "providers" {
		GetProviders(w)
	} else {
		w.Write([]byte("path must consist of 3 parameters"))
		w.WriteHeader(http.StatusBadRequest)
	}
}

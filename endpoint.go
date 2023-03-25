package main

import (
	"encoding/json"
	"net/http"
)

type OrderPageWeb struct {
	Price      string `json:"price"`
	Advertiser string `json:"advertiser"`
	Available  string `json:"available"`
	Payment    string `json:"payment"`
}

type OrdersPageWeb []OrderPageWeb

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	orders := map[string]map[string]map[OrderType]OrdersPageWeb{}
	for entry := range ordersByPageHandle.IterBuffered() {
		for _, order := range entry.Val {
			provider := "Binance"
			if orders[provider] == nil {
				orders[provider] = make(map[string]map[OrderType]OrdersPageWeb)
			}
			if orders[provider][order._currency] == nil {
				orders[provider][order._currency] = make(map[OrderType]OrdersPageWeb)
			}
			orders[provider][order._currency][order._type] = append(orders[provider][order._currency][order._type], OrderPageWeb{
				Price:      order.price,
				Advertiser: order.advertiser,
				Available:  order.available,
				Payment:    order.payment,
			})
		}
	}
	switch r.Method {
	case "GET":
		enc := json.NewEncoder(w)
		enc.SetIndent("", "    ")
		if err := enc.Encode(orders); err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
	}
}

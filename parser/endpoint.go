package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type OrderPageWeb struct {
	Price      string `json:"price"`
	Advertiser string `json:"advertiser"`
	Available  string `json:"available"`
	Payment    string `json:"payment"`
}

type OrdersPageWeb []OrderPageWeb

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("only get supported"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	trimmedUrl := strings.TrimPrefix(r.URL.Path, "/")
	params := strings.Split(trimmedUrl, "/")
	if len(params) != 3 {
		w.Write([]byte("path must consist of 3 parameters"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	provider := params[0]
	currency := params[1]
	_type := params[2]

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
			}
			res = append(res, orderWeb)
		}
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	if err := enc.Encode(res); err != nil {
		println(err) // todo log
	}
	w.WriteHeader(http.StatusOK)
}

package main

import (
	cmap "github.com/orcaman/concurrent-map/v2"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Order struct {
	price      float64
	advertiser string
	available  string
	payment    string
	commission string
	timestamp  time.Time
}

func leaveOnlyDigitsAndDots(s string) string {
	return strings.Map(
		func(r rune) rune {
			if r == '.' || unicode.IsDigit(r) {
				return r
			}
			return -1
		},
		s,
	)
}

func ParseFloat(float string) (float64, error) {
	float = leaveOnlyDigitsAndDots(float)
	return strconv.ParseFloat(float, 64)
}
func getBestDeal(orders OrdersWeb, min bool) (res Order) { // should we find minumum (true) or maximum (false)
	for _, order := range orders {
		price, err := ParseFloat(order.Price)
		if err != nil {
			log.Printf("unable to parse float %s", order.Price)
			continue
		}
		timestamp, err := time.Parse(time.RFC3339, order.Timestamp)
		curr := Order{
			price:      price,
			advertiser: order.Advertiser,
			available:  order.Available,
			payment:    order.Payment,
			commission: order.Commission,
			timestamp:  timestamp,
		}
		measurement := price > res.price
		if min {
			measurement = !measurement
		}
		if res.timestamp.IsZero() || measurement {
			res = curr
		}
	}
	return
}
func UpdateData() {
	for _, provider := range globalProviders {
		for _, token := range provider.Tokens {
			if data == nil {
				data = make(DataType)
			}
			if data[provider.Provider] == nil {
				data[provider.Provider] = make(map[string]cmap.ConcurrentMap[string, Order])
			}
			data[provider.Provider][token] = cmap.New[Order]()
			for _, op := range []string{"Sell", "Buy"} {
				m := data[provider.Provider][token]
				orders, err := getOrders(provider.Provider, token, op)
				if err != nil {
					log.Printf("unable to get orders for %s %s %s : %s", provider.Provider, token, op, err)
				}
				if orders == nil {
					continue
				}
				m.Set(op, getBestDeal(orders, op == "Sell")) // find
			}
		}
	}
}

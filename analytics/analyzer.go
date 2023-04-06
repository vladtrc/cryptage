package main

import (
	"fmt"
	"log"
)

type Providers = map[string][]string
type OrderFilter struct {
	op        string
	providers Providers
}

type FullOrderInfo struct {
	provider string
	token    string
	op       string
	order    Order
}

func filterOrders(filter OrderFilter) (orders []FullOrderInfo) {
	var operators []string
	if filter.op != "" {
		operators = []string{filter.op}
	} else {
		operators = []string{"Sell", "Buy"}
	}
	providers := make(Providers)
	if filter.providers != nil {
		providers = filter.providers
	} else {
		for _, provider := range globalProviders {
			var tokens []string
			for _, token := range provider.Tokens {
				tokens = append(tokens, token)
			}
			providers[provider.Provider] = tokens
		}
	}
	for _, op := range operators {
		for provider, tokens := range providers {
			for _, token := range tokens {
				m := data[provider][token]
				order, ok := m.Get(op)
				if !ok {
					//log.Printf("could not find element at %s %s %s", provider, token, op)
					continue
				}
				orders = append(orders, FullOrderInfo{
					provider: provider,
					token:    token,
					op:       op,
					order:    order,
				})
			}
		}
	}
	return
}

func FormatOrder(order Order) string {
	return fmt.Sprintf(
		"%.2f\nПокупатель %s\nДоступно %s\nКомиссия %s\nОплата %s\n%s",
		order.price,
		order.advertiser,
		order.available,
		order.commission,
		order.payment,
		order.timestamp.Format("2006-01-02 15:04:05"),
	)
}
func FormatReportMessage(sellOrder, buyOrder FullOrderInfo) string {
	return fmt.Sprintf(
		"%s/RUB:\n\n%s продажа:%s\n\n%s покупка:%s\n\nПрофит:%.2f percent",
		sellOrder.token,
		sellOrder.provider, FormatOrder(sellOrder.order),
		buyOrder.provider, FormatOrder(buyOrder.order),
		(buyOrder.order.price/sellOrder.order.price-1)*100,
	)
}
func AnalyzeData() {
	minThresholdCoeff := 1 + config.minThresholdPercent/100
	for _, buyOrder := range filterOrders(OrderFilter{op: "Buy"}) {
		providers := getProvidersWithSameToken(buyOrder.token)
		for _, sellOrder := range filterOrders(OrderFilter{op: "Sell", providers: providers}) {
			if sellOrder.order.price*minThresholdCoeff < buyOrder.order.price {
				message := FormatReportMessage(sellOrder, buyOrder)
				err := TgBotBroadcast(message)
				if err != nil {
					log.Printf("could not notify with tg bot: %s", err)
				}
			}
		}
	}
}

func getProvidersWithSameToken(token string) (res Providers) {
	res = make(Providers)
	for _, globalProvider := range globalProviders {
		var tokens []string
		for _, globalToken := range globalProvider.Tokens {
			if globalToken == token {
				tokens = append(tokens, globalToken)
			}
		}
		res[globalProvider.Provider] = tokens
	}
	return
}

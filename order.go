package main

type OrderType string

const (
	Sell OrderType = "Sell"
	Buy            = "Buy"
)

type Order struct {
	price      string
	advertiser string
	available  string
	payment    string
	commission string
	_currency  string
	_type      OrderType
}

type Orders []Order

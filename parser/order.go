package main

import "time"

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
	timestamp  time.Time
	_currency  string
	_type      OrderType
}

type Orders []Order

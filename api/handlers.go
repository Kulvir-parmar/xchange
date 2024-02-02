package api

import "net/http"

type User struct {
	userId  string
	balance map[string]float64
}

type Order struct {
	Price    float64
	Quantity float64
	Side     string
}

const TICKER = "BTC/USDT"

var bid []Order
var ask []Order

var users = []User{
	{userId: "1", balance: map[string]float64{"BTC": 1.00, "USDT": 10000.00}},
	{userId: "2", balance: map[string]float64{"BTC": 1.00, "USDT": 10000.00}},
}

func (s *Server) order(w http.ResponseWriter, r *http.Request) {
	// Create an order
	// Execute the order
}

func (s *Server) depth(w http.ResponseWriter, r *http.Request) {
	// Get the current Order book
}

func (s *Server) balance(w http.ResponseWriter, r *http.Request) {
	// Get the current balance of user
}

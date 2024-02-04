package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type User struct {
	UserId  string
	Balance map[string]float64
}

type Order struct {
	UserId   string
	Side     string
	Price    float64
	Quantity float64
}

const TICKER = "BTC/USDT"

var bids []Order
var asks []Order

// TODO: connect db for users
var users = []User{
	{UserId: "1", Balance: map[string]float64{"BTC": 1.00, "USDT": 100000.00}},
	{UserId: "2", Balance: map[string]float64{"BTC": 1.00, "USDT": 100000.00}},
}

func (s *Server) order(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var order Order
	json.Unmarshal(body, &order)

	remainingQuantity := fillOrder(order)

	if remainingQuantity == 0.0 {
		w.Write([]byte("Order Fulfilled"))
		return
	}
}

func (s *Server) depth(w http.ResponseWriter, r *http.Request) {
	orderBook := make(map[float64][]Order)

	for _, order := range asks {
		orderBook[order.Price] = append(orderBook[order.Price], order)
	}
	for _, order := range bids {
		orderBook[order.Price] = append(orderBook[order.Price], order)
	}

	// TODO: return order book JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orderBook)
}

func (s *Server) balance(w http.ResponseWriter, r *http.Request) {
}

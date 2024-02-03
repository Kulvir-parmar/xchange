package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

type User struct {
	UserId  string
	balance map[string]float64
}

type Order struct {
	UserId   string
	Side     string
	Price    float64
	Quantity float64
}

const TICKER = "BTC/USDT"

var bid []Order
var ask []Order

var users = []User{
	{UserId: "1", balance: map[string]float64{"BTC": 1.00, "USDT": 10000.00}},
	{UserId: "2", balance: map[string]float64{"BTC": 1.00, "USDT": 10000.00}},
}

// TODO: Implement incoming request handling
func (s *Server) order(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var order Order
	json.Unmarshal(body, &order)

	remainingQuantity := fulfillOrder()

	if remainingQuantity == 0.0 {
		w.Write([]byte("Fulfilled"))
	}

	if order.Side == "BID" {
		newOrder := Order{UserId: order.UserId, Price: order.Price, Quantity: remainingQuantity}
		bid = append(bid, newOrder)

		sort.Slice(bid, func(i, j int) bool {
			return bid[i].Price > bid[j].Price
		})
	} else {
		ask = append(ask, order)

		sort.Slice(ask, func(i, j int) bool {
			return ask[i].Price < ask[j].Price
		})
	}

	fmt.Println(ask)
}

// TODO: make an orderBook from all the orders
func (s *Server) depth(w http.ResponseWriter, r *http.Request) {

	for _, order := range ask {
		fmt.Println(order)
	}
	for _, order := range bid {
		fmt.Println(order)
	}
}

func (s *Server) balance(w http.ResponseWriter, r *http.Request) {
}

func fulfillOrder() float64 {
	return 0.0
}

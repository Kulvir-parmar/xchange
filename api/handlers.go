package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	userId  string
	balance map[string]float64
}

type Order struct {
	userId   string
	Price    float64
	Quantity float64
}

const TICKER = "BTC/USDT"

var bid []Order
var ask []Order

var users = []User{
	{userId: "1", balance: map[string]float64{"BTC": 1.00, "USDT": 10000.00}},
	{userId: "2", balance: map[string]float64{"BTC": 1.00, "USDT": 10000.00}},
}

// TODO: Implement incoming request handling
func (s *Server) order(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	body.Decode(&body)

	remainingQuantity := fulfillOrder()

	if remainingQuantity == 0.0 {
		w.Write([]byte("Fulfilled"))
	}

	/* if body.side == "BID" {
		order := Order{userId: body.userId, Price: body.price, Quantity: remainingQuantity}
		bid = append(bid, order)

		sort.Slice(bid, func(i, j int) bool {
			return bid[i].Price > bid[j].Price
		})
	} else {
		order := Order{userId: body.userId, Price: body.price, Quantity: remainingQuantity}
		ask = append(ask, order)

		sort.Slice(ask, func(i, j int) bool {
			return ask[i].Price < ask[j].Price
		})
	} */
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

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
	Balance map[string]float64
}

type Order struct {
	UserId   string
	Side     string
	Price    float64
	Quantity float64
}

// Single dev working on this so only one TICKER available.
const TICKER = "BTC/USDT"

/*
For lower latency we are keep order book in memory only rather than some db.
Downside is that we will loose orderbook in case of system failure.
Maybe Implement some kind of persistence layer in future. IDK now.
*/
var bids []Order
var asks []Order

var Users = []User{
	{UserId: "1", Balance: map[string]float64{"BTC": 1.00, "USDT": 100000.00}},
	{UserId: "2", Balance: map[string]float64{"BTC": 1.00, "USDT": 100000.00}},
}

func (s *Server) order(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var order Order
	json.Unmarshal(body, &order)

	remainingQuantity := fillOrder(order)

	if remainingQuantity == 0.0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]float64{"Filled Quantity": order.Quantity})
		fmt.Println("Order filled completely")
		return
	}

	if order.Side == "BID" {
		order.Quantity = remainingQuantity
		bids = append(bids, order)

		sort.Slice(bids, func(i, j int) bool {
			return bids[i].Price > bids[j].Price
		})

	} else {
		order.Quantity = remainingQuantity
		asks = append(asks, order)

		sort.Slice(asks, func(i, j int) bool {
			return asks[i].Price < asks[j].Price
		})
	}

	fmt.Println(asks)
	fmt.Println(bids)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"Filled Quantity": order.Quantity - remainingQuantity})

	fmt.Println("Order Placed in orderbook")
}

// FIX: fix this api

/* func (s *Server) depth(w http.ResponseWriter, r *http.Request) {
	orderBook := make(map[float64][]Order)

	for _, order := range asks {
	}
	for _, order := range bids {
	}
} */

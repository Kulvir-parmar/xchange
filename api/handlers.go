package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

// Single dev working on this so only one TICKER available.
const TICKER = "BTC/USDT"

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"Filled Quantity": order.Quantity - remainingQuantity})
}

// FIX: fix this api

/* func (s *Server) depth(w http.ResponseWriter, r *http.Request) {
	orderBook := make(map[float64][]Order)

	for _, order := range asks {
	}
	for _, order := range bids {
	}
} */

/*
Get the amount at which Market Order can be filled for given quantity.
go to all ASK orders and add up all their prices until the quantity is filled.
Send this quote to the user and if user agrees then fill the order.

Single dev working hence only quote for BID orders are generated.
Fund me to add more functionality. jk jk just fo!
*/
func (s *Server) quote(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	type QuoteQuantity struct {
		Quantity float64
	}
	var quoteQuantity QuoteQuantity
	json.Unmarshal(body, &quoteQuantity)

	var quotePrice float64

	if len(asks) == 0 {
		http.Error(w, "No sellers available!! GG", http.StatusNotFound)
		return
	}

	for _, order := range asks {
		if quoteQuantity.Quantity > order.Quantity {
			quotePrice += order.Price * order.Quantity
			quoteQuantity.Quantity -= order.Quantity
		} else {
			quotePrice += order.Price * quoteQuantity.Quantity
			break
		}
	}

	http.Header.Add(w.Header(), "content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"Quote": quotePrice})
}

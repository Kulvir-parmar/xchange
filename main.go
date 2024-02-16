package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

const wsendpoint = "wss://fstream.binance.com/stream?streams=btcusdt@depth"

type BinanceOrderBook struct {
	Asks [][]string `json:"a"`
	Bids [][]string `json:"b"`
}

type BinanceDepthResponse struct {
	Stream string           `json:"stream"`
	Data   BinanceOrderBook `json:"data"`
}

func main() {
	fmt.Println("Welcome to JJ Exchange")

	conn, _, err := websocket.DefaultDialer.Dial(wsendpoint, nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}

	var result BinanceDepthResponse

	for {
		if err := conn.ReadJSON(&result); err != nil {
			log.Fatal("read: ", err)
			return
		}

		fmt.Printf("Bids: %v\n", result.Data.Bids)
		fmt.Printf("Asks: %v\n", result.Data.Asks)
	}
}

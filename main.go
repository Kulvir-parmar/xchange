package main

import (
	"fmt"
	"log"

	"github.com/Kulvir-parmar/xchange/api"
)

func main() {
	fmt.Println("Welcome to JJ Exchange")

	server := api.NewServer(":3000")
	log.Fatal(server.Start())
}

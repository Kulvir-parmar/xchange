package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello World! \n")
}

func main() {
	fmt.Println("Welcome to JJ Exchange")

	http.HandleFunc("/", hello)

	fmt.Println("server is up and running wild at port 3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}

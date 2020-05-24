package main

import (
	"fmt"
	"log"
	"net/http"
	"owknight-be/router"
)

func main() {
	fmt.Println("Listen :10010")
	mux := router.V1()

	if err := http.ListenAndServe(":10010", mux); err != nil {
		log.Fatal(err)
	}
}

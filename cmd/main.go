package main

import (
	"log"
	"net/http"

	"github.com/tabi-bito3/crypto-price-fetcher/internal/prices"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/prices", prices.PriceHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

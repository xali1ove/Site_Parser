package main

import (
	"log"
	"lolkek/internal/handlers"
	"net/http"
)

func main() {
	http.Handle("/crypto_rank", handlers.NewCryptoRankHandler())
	http.Handle("/coin_gecko", handlers.NewCoinGeckoHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

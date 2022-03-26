package handlers

import (
	"encoding/json"
	"log"
	"lolkek/internal/pkg/models"
	"net/http"
)

type coinGeckoHandler struct {
}

func NewCoinGeckoHandler() *coinGeckoHandler {
	return &coinGeckoHandler{}
}

func (c *coinGeckoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	cryptos := make([]models.CoinGecko, 0, 65)

	response, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&per_page=65")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewDecoder(response.Body).Decode(&cryptos)
	defer response.Body.Close()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	// todo логика таблицы

	// eto chast nahren
	//debug


	writer.Header().Set("content-type", "application/json")
	err = json.NewEncoder(writer).Encode(cryptos)

	log.Println(err)
}

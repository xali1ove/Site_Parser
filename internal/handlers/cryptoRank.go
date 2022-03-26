package handlers

import (
	"encoding/json"
	"log"
	"lolkek/internal/pkg/models"
	"net/http"
)

type cryptoRankHandler struct {
}

func NewCryptoRankHandler() *cryptoRankHandler {
	return &cryptoRankHandler{}
}

func (c *cryptoRankHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	cryptos := models.NewCryptoRank()

	response, err := http.Get("https://api.cryptorank.io/v1/currencies?limit=3&api_key=ec0e8e9c19da2455ff83c3d0c6988b9a5c0646e6897df4bd18740338d0da")
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

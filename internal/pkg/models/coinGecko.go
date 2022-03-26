package models

type CoinGecko struct {
	Name         string  `json:"name"`
	CurrentPrice float64 `json:"current_price"`
}

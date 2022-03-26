package models

type CryptoRankData struct {
	Rank   int    `json:"rank"`
	Name   string `json:"name"`
	Values struct {
		USD struct {
			Price float64 `json:"price"`
		} `json:"USD"`
	} `json:"values"`
}

type CryptoRank struct {
	Data []CryptoRankData `json:"data"`
}

func NewCryptoRank() *CryptoRank {
	return &CryptoRank{Data: make([]CryptoRankData, 0, 3)}
}

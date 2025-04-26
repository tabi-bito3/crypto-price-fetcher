package prices

import (
	"encoding/json"
	"net/http"
)

const baseURL = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd"

type PriceResponse struct {
	Bitcoin  CoinPrice `json:"bitcoin"`
	Ethereum CoinPrice `json:"ethereum"`
}

type CoinPrice struct {
	USD float64 `json:"usd"`
}

func FetchPrices() (PriceResponse, error) {
	resp, err := http.Get(baseURL)
	if err != nil {
		return PriceResponse{}, err
	}
	defer resp.Body.Close()

	var priceResponse PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		return PriceResponse{}, err
	}
	return priceResponse, nil
}

package prices

import (
	"encoding/json"
	"net/http"
)

func PriceHandler(w http.ResponseWriter, r *http.Request) {
	prices, err := FetchPrices()
	if err != nil {
		http.Error(w, "Failed to fetch prices", http.StatusInternalServerError)
		return
	}

	symbol := r.URL.Query().Get("symbol")
	var price float64

	switch symbol {
	case "BTC":
		price = prices.Bitcoin.USD
	case "ETH":
		price = prices.Ethereum.USD
	default:
		http.Error(w, "Symbol not supported", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"symbol": symbol,
		"price":  price,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package ticker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FinnhubQuote struct {
	Current       float64 `json:"c"`  // current price
	Change        float64 `json:"d"`  // change from prev close
	PercentChange float64 `json:"dp"` // percent change from prev close
	High          float64 `json:"h"`  // daily high
	Low           float64 `json:"l"`  // daily low
	Open          float64 `json:"o"`  // day's open price
	PreviousClose float64 `json:"pc"` // prev close
	Timestamp     int64   `json:"t"`  // unix timestamp
}

func FetchQuote(symbol, apiKey string) (FinnhubQuote, error) {
	url := fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s", symbol)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return FinnhubQuote{}, err
	}

	req.Header.Set("X-Finnhub-Token", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return FinnhubQuote{}, err
	}
	defer resp.Body.Close()

	var quote FinnhubQuote
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&quote)

	if err != nil {
		return FinnhubQuote{}, err
	}

	// check if ticker is invalid
	if quote.Current == 0 {
		return FinnhubQuote{}, fmt.Errorf("ticker symbol '%s' not found", symbol)
	}

	return quote, nil
}

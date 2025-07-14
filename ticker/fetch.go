package ticker

import (
	"encoding/json"
	"fmt"
	"io"
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
	baseurl := "https://finnhub.io/api/v1"
	url := fmt.Sprintf("%s/quote?symbol=%s", baseurl, symbol)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return FinnhubQuote{}, err
	}

	req.Header.Set("X-Finnhub-Token", apiKey)

	resp, err := client.Do(req)
	if resp.StatusCode == 429 {
		return FinnhubQuote{}, fmt.Errorf("error: Rate limit reached at %s", baseurl)
	}

	if err != nil {
		return FinnhubQuote{}, err
	}
	defer resp.Body.Close()

	var quote FinnhubQuote
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&quote)

	if err != nil {
		fmt.Println(symbol)
		return FinnhubQuote{}, err
	}

	// check if ticker is invalid
	if quote.Current == 0 {
		return FinnhubQuote{}, fmt.Errorf("ticker symbol '%s' not found", symbol)
	}

	return quote, nil
}

func FetchSPY() (string, error) {
	url := "https://stockanalysis.com/list/sp-500-stocks"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error: failed to fetch %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: non-500 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error: failed to read page body from %s: %w", url, err)
	}

	return string(body), nil
}

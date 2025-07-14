package data

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ilubbe/tickercli/ticker"
)

type Stock struct {
	Rank   string
	Ticker string
}

func GetTop20() ([]Stock, error) {
	html, err := ticker.FetchSPY()
	if err != nil {
		return nil, fmt.Errorf("error: failed to fetch HTML: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("error: failed to parse HTML: %w", err)
	}

	var stocks []Stock

	doc.Find("table tbody tr").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if i >= 20 {
			return false
		}

		rank := s.Find("td").First().Text()
		rank = strings.TrimSpace(rank)
		ticker := s.Find("td").Eq(1).Text()
		ticker = strings.TrimSpace(ticker)

		if ticker != "" && rank != "" {
			stocks = append(stocks, Stock{
				Rank:   rank,
				Ticker: ticker,
			})
		}

		return true
	})

	return stocks, nil
}

package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/ilubbe/tickercli/cmd"
	"github.com/ilubbe/tickercli/data"
	"github.com/ilubbe/tickercli/ticker"
)

func main() {
	opts := cmd.ParseFlags()

	if opts.Symbol != "" {
		ticker.GetQuote(opts.Symbol)
	}

	if opts.Top20 || opts.Gainers || opts.Losers {
		stocks, err := data.GetTop20()
		if err != nil || len(stocks) == 0 {
			fmt.Fprintf(os.Stderr, "error: could not get list of top 20 stocks in S&P 500")
			os.Exit(1)
		}

		if opts.Top20 {
			var wg sync.WaitGroup

			for _, stockSymbol := range stocks {
				wg.Add(1)
				go func(stockSymbol string) {
					defer wg.Done()
					ticker.GetQuote(stockSymbol)
				}(stockSymbol.Ticker)
			}

			wg.Wait()
		}
	}
}

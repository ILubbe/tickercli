package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ilubbe/tickercli/colors"
	"github.com/ilubbe/tickercli/ticker"
)

func main() {
	symbol := flag.String("symbol", "", "Ticker to lookup (e.g. AAPL)")
	symbolShort := flag.String("s", "", "Ticker to lookup (e.g. AAPL)")
	flag.Parse()

	chosenSymbol := *symbol
	if *symbolShort != "" {
		chosenSymbol = *symbolShort
	}

	if chosenSymbol == "" {
		fmt.Println("Usage: tickercli [-s|-symbol] SYMBOL")
		os.Exit(1)
	}

	chosenSymbol = strings.ToUpper(chosenSymbol)

	apiKeyBytes, err := os.ReadFile("api.key")
	if err != nil {
		fmt.Printf("%sError reading API key:%s%s\n", colors.Red, err, colors.Reset)
		os.Exit(1)
	}

	apiKey := strings.TrimSpace(string(apiKeyBytes))

	quote, err := ticker.FetchQuote(chosenSymbol, apiKey)
	if err != nil {
		fmt.Printf("%sError fetching quote:%s%s\n", colors.Red, err, colors.Reset)
		os.Exit(1)
	}

	timestamp := time.Unix(quote.Timestamp, 0)

	fmt.Println()
	fmt.Printf("Quote for %s%s%s at %s:\n", colors.Blue, chosenSymbol, colors.Reset, timestamp.Format(time.RFC1123))
	fmt.Printf("  Current Price  : %s%.2f%s\n", colors.Blue, quote.Current, colors.Reset)
	colorCode := colors.DetermineColor(quote.Change)
	fmt.Printf("  Change         : %s%.2f (%.2f%%)%s\n", colorCode, quote.Change, quote.PercentChange, colors.Reset)
	fmt.Printf("  Day High       : %.2f\n", quote.High)
	fmt.Printf("  Day Low        : %.2f\n", quote.Low)
	fmt.Printf("  Open           : %.2f\n", quote.Open)
	fmt.Printf("  Previous Close : %.2f\n", quote.PreviousClose)
	fmt.Println()
}

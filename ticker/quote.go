package ticker

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ilubbe/tickercli/colors"
)

var printMutex sync.Mutex

func GetQuote(symbol string) {
	symbol = strings.ToUpper(symbol)
	apiKeyBytes, err := os.ReadFile("api.key")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%serror: could not read API key:%s%s\n", colors.Red, err, colors.Reset)
		os.Exit(1)
	}

	apiKey := strings.TrimSpace(string(apiKeyBytes))

	quote, err := FetchQuote(symbol, apiKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%serror: could not fetch quote: %s%s\n", colors.Red, err, colors.Reset)
		os.Exit(1)
	}

	timestamp := time.Unix(quote.Timestamp, 0)

	printMutex.Lock()
	fmt.Println()
	fmt.Printf("Quote for %s%s%s at %s:\n", colors.Blue, symbol, colors.Reset, timestamp.Format(time.RFC1123))
	fmt.Printf("  Current Price  : %s%.2f%s\n", colors.Blue, quote.Current, colors.Reset)
	colorCode := colors.DetermineColor(quote.Change)
	fmt.Printf("  Change         : %s%.2f (%.2f%%)%s\n", colorCode, quote.Change, quote.PercentChange, colors.Reset)
	fmt.Printf("  Day High       : %.2f\n", quote.High)
	fmt.Printf("  Day Low        : %.2f\n", quote.Low)
	fmt.Printf("  Open           : %.2f\n", quote.Open)
	fmt.Printf("  Previous Close : %.2f\n", quote.PreviousClose)
	fmt.Println()
	printMutex.Unlock()
}

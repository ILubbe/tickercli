package cmd

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "\nUsage of %s:\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "  --symbol, -s   Ticker to lookup (e.g. AAPL)")
		fmt.Fprintln(os.Stderr, "  --top20, -t    Show top 20 tickers")
		fmt.Fprintln(os.Stderr, "  --gainers, -g  Show top 5 gainers (can combine with --losers)")
		fmt.Fprintln(os.Stderr, "  --losers, -l   Show top 5 losers (can combine with --gainers)")
		fmt.Fprintln(os.Stderr, "  --help, -h     Show this help menu")
	}
}

type Options struct {
	Symbol  string
	Top20   bool
	Gainers bool
	Losers  bool
	Help    bool
}

func ParseFlags() Options {
	symbol := flag.String("symbol", "", "Ticker to lookup (e.g. AAPL)")
	symbolShort := flag.String("s", "", "Ticker to lookup (e.g. AAPL)")
	top20 := flag.Bool("top20", false, "List the top 20 companies in the S&P 500")
	top20Short := flag.Bool("t", false, "List the top 20 companies in the S&P 500")
	gainers := flag.Bool("gainers", false, "List top 5 daily gainers in the S&P 500")
	gainersShort := flag.Bool("g", false, "List top 5 daily gainers in the S&P 500")
	losers := flag.Bool("losers", false, "List top 5 daily losers in the S&P 500")
	losersShort := flag.Bool("l", false, "List top 5 daily losers in the S&P 500")
	help := flag.Bool("help", false, "Show this help menu")
	helpShort := flag.Bool("h", false, "Show this help menu")

	flag.Parse()

	chosenSymbol := *symbol
	if *symbolShort != "" {
		chosenSymbol = *symbolShort
	}

	opts := Options{
		Symbol:  chosenSymbol,
		Top20:   *top20 || *top20Short,
		Gainers: *gainers || *gainersShort,
		Losers:  *losers || *losersShort,
		Help:    *help || *helpShort,
	}

	validateFlags(opts)

	return opts

}

func validateFlags(opts Options) {
	num := 0
	if opts.Symbol != "" {
		if opts.Symbol[0] == '-' {
			printHelpAndExit("error: Symbol cannot start with a '-'")
		}
		num++
	}
	if opts.Top20 {
		num++
	}
	if opts.Gainers {
		num++
	}
	if opts.Losers {
		num++
	}
	if opts.Help {
		num++
	}

	if num == 0 {
		printHelpAndExit("error: You must provide at least one flag.")
	}

	if opts.Symbol != "" && (opts.Top20 || opts.Gainers || opts.Losers || opts.Help) {
		printHelpAndExit("error: --symbol or -s cannot be used with other flags.")
	}

	if opts.Top20 && (opts.Symbol != "" || opts.Gainers || opts.Losers || opts.Help) {
		printHelpAndExit("error: --top20 or -t cannot be used with other flags.")
	}

	if (opts.Gainers || opts.Losers) && num > 2 {
		printHelpAndExit("error: Only --gainers and --losers may be combined.")
	}

	if opts.Help {
		printHelpAndExit("")
	}
}

func printHelpAndExit(msg string) {
	if msg != "" {
		fmt.Fprintln(os.Stderr, msg)
	}
	flag.Usage()
	if msg == "" {
		os.Exit(0)
	}
	os.Exit(1)
}

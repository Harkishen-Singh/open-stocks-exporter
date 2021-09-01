package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func parseFlags(cfg *config, args []string) {
	flag.DurationVar(&cfg.fetchInterval, "fetch-interval", time.Minute, "Interval after which stock entities should be updated.")
	flag.StringVar(&cfg.symbolsStr, "symbols", "", "Comma separated symbols of stocks, as per Yahoo finance.")
	flag.StringVar(&cfg.logFormat, "log-format", "logfmt", "Logging format. Valid options: ['logfmt', 'json'].")
	flag.StringVar(&cfg.logLevel, "log-level", "debug", "Logging format. Valid options: ['debug', 'info', 'warn', 'error'].")

	_ = flag.CommandLine.Parse(args)
}

// processFlags processes and validates the received flag values.
func processFlags(cfg *config) error {
	if err := processSymbols(cfg); err != nil {
		return fmt.Errorf("process symbols: %w", err)
	}
	return nil
}

func processSymbols(cfg *config) error {
	s := cfg.symbolsStr
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return fmt.Errorf("empty symbols are not allowed")
	}
	symbols := strings.Split(s, ",")
	cfg.symbols = symbols
	return nil
}

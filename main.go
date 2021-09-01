package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Harkishen-Singh/open-stocks-exporter/log"
)


type config struct {
	fetchInterval time.Duration
	symbols []string
	symbolsStr string
	logFormat string
	logLevel string
}

func main() {
	conf := new(config)
	args := os.Args[1:]

	parseFlags(conf, args)
	err := processFlags(conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, "processing flags:", err.Error())
		os.Exit(1)
	}

	if err = log.Init(log.Config{Level: conf.logLevel, Format: conf.logFormat}); err != nil {
		fmt.Fprintln(os.Stderr, "log init:", err.Error())
		os.Exit(1)
	}
	log.Info("msg", fmt.Sprintf("%v", conf))
}

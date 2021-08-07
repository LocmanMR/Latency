package main

import (
	"./src/urlLatency"
	"flag"
)

func main() {
	var timeout string

	flag.StringVar(&timeout, "timeout", "0ns", "The http client timeout")
	flag.Parse()

	urlLatency.UrlLatency(timeout)
}
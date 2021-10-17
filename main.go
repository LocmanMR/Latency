package main

import (
	"Latency"
	"flag"
)

func main() {
	var timeout string

	flag.StringVar(&timeout, "timeout", "0ns", "The http client timeout")
	flag.Parse()

	Latency.UrlLatency(timeout)
}

package main

import (
	"flag"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
	"sync"
)

const Https = "https://"

var client = http.Client{}

func main() {
	var timeout string

	flag.StringVar(&timeout, "timeout", "0ns", "The http client timeout")
	flag.Parse()

	UrlLatency(timeout)
}

func UrlLatency(timeout string)  {
wg := &sync.WaitGroup{}

	setHttpClientTimeout(timeout)

	for _, domain := range getDomains() {
		wg.Add(1)
		go ping(domain, wg)
	}

	wg.Wait()
}

func getDomains() []string {
	urls := make([]string, 0)

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		urls = append(urls, in.Text())
	}

	return urls
}

func ping(domain string, wg *sync.WaitGroup) {
	defer wg.Done()

	url := Https + domain
	startTime := time.Now()
	request, err := http.NewRequest("GET", url, nil)
	endTime := time.Now()

	if err != nil {
		fmt.Println(domain, 0, endTime.Sub(startTime))
		return
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(domain, 0, endTime.Sub(startTime))
		return
	}
	defer response.Body.Close()

	fmt.Println(domain, response.StatusCode, endTime.Sub(startTime))
}

func setHttpClientTimeout(timeout string) {
	client.Timeout, _ = time.ParseDuration(timeout)
}

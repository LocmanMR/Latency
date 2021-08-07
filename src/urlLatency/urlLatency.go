package urlLatency

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

const Https = "https://"

var client = http.Client{}

func UrlLatency(timeout string)  {
	setHttpClientTimeout(timeout)

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		 url := in.Text()

		 fmt.Println(url)
		 fmt.Println(ping(url))
	}
}

func ping(domain string) (int, time.Duration) {
	url := Https + domain
	startTime := time.Now()
	request, err := http.NewRequest("GET", url, nil)
	endTime := time.Now()
	if err != nil {
		return 0, endTime.Sub(startTime)
	}

	response, err := client.Do(request)
	if err != nil {
		return 0, endTime.Sub(startTime)
	}

	return response.StatusCode, endTime.Sub(startTime)
}

func setHttpClientTimeout(timeout string) {
	client.Timeout, _ = time.ParseDuration(timeout)
}

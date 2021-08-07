package urlLatency

import (
	"bufio"
	"fmt"
	"os"
)

func UrlLatency()  {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		 url := in.Text()
		 fmt.Println(len(url))
	}
}

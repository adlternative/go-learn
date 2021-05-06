package main

import (
	_ "basic_learn/memo"
	_ "basic_learn/memo2"
	_ "basic_learn/memo3"
	_ "basic_learn/memo4"
	"basic_learn/memo5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
func incomingURLs() []string {
	return os.Args[1:]
}
func main() {
	m := memo5.New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			if value != nil {
				fmt.Printf("%s, %s, %d bytes\n",
					url, time.Since(start), len(value.([]byte)))
			}
			n.Done()
		}(url)
	}
	n.Wait()
}

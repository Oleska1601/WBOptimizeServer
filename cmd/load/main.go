package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		url         string
		concurrency int
		duration    int
		method      string
	)

	flag.StringVar(&url, "url", "http://localhost:8081/api/v1/cpu", "target URL")
	flag.IntVar(&concurrency, "c", 5, "number of concurrent requests")
	flag.IntVar(&duration, "d", 30, "duration in seconds")
	flag.StringVar(&method, "method", "GET", "HTTP method")
	flag.Parse()

	var (
		success int64
		failed  int64
		wg      sync.WaitGroup
		stop    = make(chan struct{})
		start   = time.Now()
	)

	fmt.Printf("Load Generator\n")
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Concurrency: %d\n", concurrency)
	fmt.Printf("Duration: %d seconds\n", duration)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			client := &http.Client{Timeout: 10 * time.Second}
			for {
				select {
				case <-stop:
					return
				default:
					var resp *http.Response
					var err error

					switch method {
					case "GET":
						resp, err = client.Get(url)
					case "POST":
						resp, err = client.Post(url, "application/json", nil)
					default:
						resp, err = client.Get(url)
					}

					if err != nil {
						atomic.AddInt64(&failed, 1)
						continue
					}
					io.Copy(io.Discard, resp.Body)
					resp.Body.Close()
					atomic.AddInt64(&success, 1)
				}
			}
		}(i)
	}

	time.Sleep(time.Duration(duration) * time.Second)
	close(stop)
	wg.Wait()

	elapsed := time.Since(start)
	total := success + failed
	rps := float64(success) / elapsed.Seconds()

	fmt.Printf("\nResults\n")
	fmt.Printf("Duration: %.2f seconds\n", elapsed.Seconds())
	fmt.Printf("Successful: %d\n", success)
	fmt.Printf("Failed: %d\n", failed)
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Requests/sec: %.2f\n", rps)
}

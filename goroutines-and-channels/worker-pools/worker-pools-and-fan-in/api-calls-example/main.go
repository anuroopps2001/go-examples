package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type ApiResult struct {
	URL      string
	status   int
	Duration time.Duration

	/*The error built-in interface type is the conventional interface
	for representing an error condition, with the nil value representing no error. */
	Error error
}

func main() {
	// slice of strings
	apis := []string{
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/3",
	}
	results := make(chan ApiResult)

	var wg sync.WaitGroup

	// Number of goroutines managed by waitgroup
	wg.Add(len(apis))

	start := time.Now()

	for _, api := range apis {
		go callApi(api, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for respose := range results {
		if respose.Error != nil {
			fmt.Println("error calling", respose.URL, ":", respose.Error)
			continue
		}
		fmt.Printf("API: %s | status: %d | took: %v\n",
			respose.URL,
			respose.status,
			respose.Duration,
		)
	}

	fmt.Println("Total time:", time.Since(start))
	fmt.Println("All API calls completed")
}

func callApi(apiurl string, result chan ApiResult, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()

	resp, err := http.Get(apiurl)
	if err != nil {
		result <- ApiResult{
			URL:   apiurl,
			Error: err,
		}
		return
	}
	defer resp.Body.Close() // Every successful http.Get must close resp.Body

	// io.ReadAll(resp.Body) // Consume the entire response body so the request completes properly.

	io.Copy(io.Discard, resp.Body) // I am intentionally reading the entire response body and throwing it away.
	result <- ApiResult{
		URL:      apiurl, // resp.Request.URL from http.Get(url)
		status:   resp.StatusCode,
		Duration: time.Since(start),
	}
}

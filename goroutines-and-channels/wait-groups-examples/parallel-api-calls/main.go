package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	apis := []string{
		"https://service-a/health",
		"https://service-b/health",
		"https://service-c/health",
		"https://service-d/health",
	}
	var wg sync.WaitGroup
	wg.Add(len(apis)) // call 5 new child goroutines

	for _, api := range apis {
		go callApi(api, &wg)
	}

	// Below for loop also works
	/* for i := 0; i < len(apis); i++ {
		api := apis[i]
		go callApi(api, &wg)
	} */

	wg.Wait() // Wait for all child goroutines to finish and return count as 0
	fmt.Println("all API calls finished")
}

func callApi(num string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("API", num, "started..")
	time.Sleep(1 * time.Second)
	fmt.Println("API", num, "Finished..!!")
}

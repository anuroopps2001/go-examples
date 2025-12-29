package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	jobs := 10
	wg.Add(jobs)

	for i := 1; i < jobs; i++ {
		go process(i, &wg)
	}
	wg.Wait()

	fmt.Println("All jobs done")

}

func process(job int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("processing job", job)
	time.Sleep(100 * time.Millisecond)
}

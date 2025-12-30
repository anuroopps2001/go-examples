package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	jobs := make(chan int)
	var wg sync.WaitGroup

	workers := 3
	wg.Add(workers)

	// start new child goroutines of single worker pool
	for i := 1; i <= workers; i++ {
		go worker(i, jobs, &wg)
	}
	// sending load into channels to be processed by 3 child goroutines which are group of an one "worker pool"
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	// since main goroutine doesn;t have any work, we don't separate goroutine again for waiting. So below code could
	// be run in main goroutine itself
	/* go func() {
		wg.Wait()
		close(jobs)
	}() */

	close(jobs)
	wg.Wait()

	fmt.Println("Processing of all the jobs is completed..!!")
	defer func() {
		duration := time.Since(startTime)
		fmt.Printf("Execution time: %s\n", duration)
	}()
}

func worker(workerId int, workJob chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// workJob is actual work to be processed and workworkerId is goroutine processing that job
	for jobNum := range workJob {
		fmt.Println("worker goroutine", workerId, "processeing job", jobNum)
	}

}

/*
package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	ch := make(chan int)

	go worker(ch)
	// sending data into channel
	for i := 1; i <= 1000; i++ {
		ch <- i
	}
	close(ch)

	// time.Sleep(5 * time.Second)

	fmt.Println("All data processing completed")
	defer func() {
		duration := time.Since(startTime)
		fmt.Printf("Execution time: %s\n", duration)
	}()

}

func worker(load chan int) {
	for job := range load {
		fmt.Println("worker processing", job)
	}

} */

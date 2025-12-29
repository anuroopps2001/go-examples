package main

import (
	"fmt"
	"sync"
)

func main() {
	/* WaitGroup is a counting semaphore typically used to wait for a group of goroutines or tasks to finish.
		Typically, a main goroutine will start tasks, each in a new goroutine,
		by calling WaitGroup.Go and then wait for all tasks to.

	WaitGroup is shared state
	Shared state must be shared by reference
	*/
	var wg sync.WaitGroup
	wg.Add(4) // statrt 4 new child goroutines counter = counter + 1

	func() {
		for i := 1; i <= 4; i++ {
			go worker(i, &wg)
		}
	}()

	/* Wait function does below tasks:=
	Main goroutine blocks
	Program does not exit early
	All workers are guaranteed to finish */
	wg.Wait() // Pause (block) this main goroutine until the WaitGroup counter becomes 0.
}

func worker(id int, wg *sync.WaitGroup) { // We pass *sync.WaitGroup (a pointer) so that all goroutines operate on the same WaitGroup instance.
	// defer calls will be executed at the end of function
	// counter = counter - 1
	defer wg.Done() // Done decrements the WaitGroup task counter by one. i,e Counter goes: 4 → 3 → 2 → 1 → 0

	fmt.Println("worker", id, "done")
}

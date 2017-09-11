package main

import (
	"allgolangdemo/queuedemo/src/queue"
	"fmt"
	"sync"
)

// stimulate this main func is non-stop server, like the A2 indexer
func main() {
	fmt.Println("main process start...")
	wg := &sync.WaitGroup{}
	// init dispatcher
	queue.StartDispatcher(5, wg)

	// queue those long-time works
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go queue.Collect(fmt.Sprintf("work_%d", i+1))
	}

	wg.Wait()

	fmt.Println("main process done")
}

package queue

import (
	"fmt"
	"sync"
)

var WorkQueue chan Work
var WorkerQueue chan Worker

// StartDispatcher init worker and dispatch work to them
func StartDispatcher(workersNum int, wg *sync.WaitGroup) {
	WorkerQueue = make(chan Worker, workersNum)
	WorkQueue = make(chan Work, 100)

	for i := 0; i < workersNum; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i + 1)
		worker.Start(wg)
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Dispatcher receive work: ", work)
				go func() {
					worker := <-WorkerQueue
					fmt.Printf("Dispatcher dispatch work %v to worker%d\n", work, worker.ID)
					worker.Work <- work
				}()
			}
		}
	}()
}

// Collect add new work to queue
func Collect(name string) {
	work := Work{
		Name: name,
	}

	WorkQueue <- work

	fmt.Println("queue work: ", work)
	return
}

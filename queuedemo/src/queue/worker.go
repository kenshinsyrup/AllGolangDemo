package queue

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	ID   int
	Work chan Work
	Quit chan bool
}

func NewWorker(id int) Worker {
	worker := Worker{
		ID:   id,
		Work: make(chan Work),
		Quit: make(chan bool),
	}
	WorkerQueue <- worker
	return worker
}

// Start watch works, once get work, just do it
func (w *Worker) Start(wg *sync.WaitGroup) {
	go func() {
		for {
			select {
			case work := <-w.Work:
				// Receive a work request.
				fmt.Printf("Worker%d receive work %v\n", w.ID, work)
				time.Sleep(time.Second * 1)
				fmt.Printf("Worker%d finish work %v!\n", w.ID, work)

				WorkerQueue <- *w
				wg.Done()
				fmt.Println("free worker back to queue, now worker num is: ", len(WorkerQueue))

			case <-w.Quit:
				fmt.Printf("Worker%d stop\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

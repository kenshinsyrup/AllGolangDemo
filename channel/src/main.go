package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := make(chan int, 5)
	var count int
	m := &sync.Mutex{}

	go func() {
		queue <- count
	}()

	for preCount := range queue {
		preCount++
		if preCount == 10 {
			fmt.Println("done")
			close(queue)
			break
		}
		enqueue(queue, preCount, m)
	}
}
func enqueue(queue chan int, count int, m *sync.Mutex) {
	for i := 0; i < 3; i++ {
		go func() {
			m.Lock()
			count++
			m.Unlock()
			// over
			if count > 10 {
				return
			}
			// tasks
			time.Sleep(1 * time.Second)
			fmt.Println("count: ", count, "l: ", len(queue))
			queue <- count
		}()
	}
}

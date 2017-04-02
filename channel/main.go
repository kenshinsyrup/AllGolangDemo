package main

import "fmt"
import "sync"
import "time"
import "errors"

func main() {
	fmt.Println("Hello, playground")
	// queue := make(chan int)
	// done := make(chan bool)
	wg := &sync.WaitGroup{}
	var p int
	// for {
	// for i := 0; i < 5; i++ {
	// 	p++
	/*
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go test(queue, done, wg)
		}

		// }
		// }

		// for i := 0; i < 1000; i++ {
		fmt.Println("??")
		for {
			p++
			d, ok := <-done
			fmt.Println("d: ", d, "ok? ", ok)
			if ok && d {
				close(queue)
				break
			} else {
				queue <- p
			}

			// select {
			// case <-done:
			// 	fmt.Println("done")
			// 	close(queue)
			// 	break Point
			// default:
			// 	runtime.Gosched()
			// 	fmt.Println("p: ", p)
			// 	queue <- p
			// }
		}
		close(done)
	*/

	// watch
	for {
		if err := readPage(p, wg); err != nil {
			break
		}
		p += 3
	}

	// wait
	wg.Wait()
	fmt.Println("all done")
}
func readPage(p int, wg *sync.WaitGroup) error {
	fmt.Println("page initial: ", p)
	// 3 goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		// m.Lock()
		p++
		// m.Unlock()
		//  over
		if p >= 10 {
			wg.Done()
			fmt.Println("should done")
			return errors.New("done")
		}
		//  process *p tasks
		go func(page int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 1000)
			fmt.Println(" sleep  task: ", page)
		}(p, wg)
	}
	time.Sleep(1 * time.Second)
	return nil
}

func test(queue chan int, done chan bool, wg *sync.WaitGroup) {
	for j := range queue {
		if j >= 10 {
			fmt.Println("now stop")
			wg.Done()
			done <- true
			// d, ok := <-done
			// if !ok {
			// 	fmt.Println("already send sign to channel done")
			// 	break
			// }

			// if d == false {
			// 	done <- true
			// }
		}
		time.Sleep(1 * time.Second)
		// do your tasks

	}
	fmt.Println("goroutine done")
	wg.Done()

	// for {
	// 	j, ok := <-queue
	// 	if !ok {
	// 		wg.Done()
	// 		return
	// 	}
	// 	// 模拟任务结束
	// 	if j > 100 {
	// 		wg.Done()
	// 		return
	// 	}

	// 	fmt.Println("now ", j)

	// }
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// func Select
	var send1 = make(chan int)
	var increase1 = func(c chan int) {
		for i := 0; i < 8; i++ {
			c <- i
		}
		close(c)
	}
	go increase1(send1)

	send2 := make(chan int)
	increase2 := func(c chan int) {
		for i := 0; i < 8; i++ {
			c <- i * 100
		}
		close(c)
	}
	go increase2(send2)

	var selectCase = make([]reflect.SelectCase, 2)
	selectCase[0].Dir = reflect.SelectRecv
	selectCase[0].Chan = reflect.ValueOf(send1)
	selectCase[1].Dir = reflect.SelectRecv
	selectCase[1].Chan = reflect.ValueOf(send2)

	counter := 0
	for counter < 1 {
		chosen, recv, recvOk := reflect.Select(selectCase)
		if recvOk {
			fmt.Println(chosen, recv.Int(), recvOk)

		} else {
			counter++
			fmt.Println("over")
		}
	}
}

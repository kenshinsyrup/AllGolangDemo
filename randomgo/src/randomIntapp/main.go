package main

import (
	"fmt"
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	// fmt.Println(RandNum(10))

	fmt.Println(random.Intn(30))
}

// func RandNum(n int) int {
// 	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
// }

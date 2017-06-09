package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(RandNum(10))
}
func RandNum(n int) int {
	return rand.New(rand.NewSource(time.Now().UTC().UnixNano())).Intn(n)
}

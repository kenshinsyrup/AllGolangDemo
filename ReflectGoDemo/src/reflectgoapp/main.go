package main

import (
	"fmt"
	"reflect"
)

func main() {
	// func Copy
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	fmt.Println("dst: ", dst)
	fmt.Println("src: ", src)

	n := reflect.Copy(reflect.ValueOf(dst), reflect.ValueOf(src))
	fmt.Println("dst: ", dst)
	fmt.Println("src: ", src)
	fmt.Println("copy num: ", n)
}

package main

import (
	"allgolangdemo/CompareSlice/equal"
	"fmt"
)

func main() {
	// a := []string{"1"}
	// b := []string(nil)
	a := []string(nil)
	b := []string{}
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(equal.StringSliceEqual(a, b))
	fmt.Println(equal.StringSliceReflectEqual(a, b))
}

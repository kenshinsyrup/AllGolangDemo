package main

import (
	"allgolangdemo/CompareSlice/equal"
	"fmt"
)

func main() {
	// a := []string(nil)
	var a []string
	b := []string{}
	fmt.Printf("a: %#v \n", a)
	fmt.Printf("b: %#v \n", b)

	fmt.Println(equal.StringSliceEqual(a, b))
	fmt.Println(equal.StringSliceReflectEqual(a, b))
	fmt.Println(equal.StringSliceEqualBCE(a, b))
}

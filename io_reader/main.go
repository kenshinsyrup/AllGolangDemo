package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	p := []byte(`hello world`)
	r1 := bytes.NewReader(p)
	r2 := bytes.NewReader(p)
	fmt.Println("hahahhah")

	fmt.Println("before r1 len:", r1.Len())
	out1, _ := ioutil.ReadAll(r1)
	fmt.Println("out1: ", string(out1))
	fmt.Println("after r1 len:", r1.Len())
	out2, _ := ioutil.ReadAll(r2)
	fmt.Println("out2: ", string(out2))
}

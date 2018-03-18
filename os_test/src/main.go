package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	d, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	for {
		if d == "" {
			panic("must run under allgolangdemo")
		}
		if strings.HasSuffix(d, "allgolangdemo") {
			break
		}
		d = path.Dir(d)
	}
	fmt.Println("project root path: ", d)
}

package genfile

import (
	"fmt"
	"os"
)

func Genfile() {
	f, err := os.Create(" ")
	fmt.Println(err)
	fmt.Println(f.Name())
}

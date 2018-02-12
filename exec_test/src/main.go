package main

import (
	"fmt"
	"os/exec"
)

func main() {
	args := []string{"-l", "-a"}
	cmd := exec.Command("ls", args...)
	// cmd := exec.Command("ls", "-l", "-a")
	// err := cmd.Run()
	// fmt.Println(err)
	output, err := cmd.Output()
	fmt.Println(err)
	fmt.Println(string(output))
}

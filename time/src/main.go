package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now: ", now)

	localNow := now.Format("2006-01-02 15:04:05 +0800 UTC")
	fmt.Println("local now: ", localNow)

	y := now.Year()
	m := now.Month()
	d := now.Day()
	fmt.Println("Year: ", y)
	fmt.Println("Month: ", m)
	fmt.Println("Day: ", d)

}

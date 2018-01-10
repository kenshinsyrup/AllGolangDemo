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

	tzUTC, err := time.LoadLocation("") //等同于"UTC"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(now.In(tzUTC))

	tzLocal, err := time.LoadLocation("Local") //Local
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(now.In(tzLocal))

	tzBeijing, err := time.LoadLocation("Asia/Shanghai") //北京时间，大写的尴尬 :)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(now.In(tzBeijing))

	// https://www.w3.org/TR/NOTE-datetime W3标准日期时间格式
	nowUTC := time.Now().UTC()
	fmt.Println("UTC now: ", nowUTC)
	fmt.Println("UTC now in W3 standard: ", nowUTC.Format("2006-01-02T15:04:05Z"))
	fmt.Println("Beijing now in W3 standard: ", nowUTC.In(tzBeijing).Format("2006-01-02T15:04:05+08:00"))
}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	// t0 := "\u6B22\u8FCE\u6765\u5230"
	// fmt.Println(t0)

	// type TS1 struct {
	// 	str string `json:"data"`
	// }

	// type TS2 struct {
	// 	run rune `json:"data"`
	// }
	// now := time.Now().Unix() // 一个无效的码点值
	// fmt.Println(now)
	// str := string(now)            // golang是utf-8编码，会对无效码点进行替换
	// fmt.Printf("%X", []byte(str)) // EFBFBD，即字符「�」

	// a := "\U0000d83c\U0000dfa4"
	// fmt.Println(a)

	// d := fmt.Sprintf("{\"data\":%s}",a)
	// fmt.Println(d)

	// var ts1 TS1
	// var ts2 TS2
	// json.Unmarshal([]byte(b), &ts1)
	// json.Unmarshal([]byte(b),&ts2)

	// fmt.Println(ts1)
	// fmt.Println(ts2)

	// 3

	d := '🎤'
	fmt.Println(d)
	fmt.Println(string(d))

	d1 := "🎤"
	fmt.Println(d1)

	data := fmt.Sprintf("{\"nickname\": \"%s\"}", string(d))
	fmt.Println(data)

	b := []byte(string(data))
	fmt.Println(b)
	fmt.Println(string(b))

	type TS1 struct {
		Data string `json:"nickname"`
	}

	var ts1 TS1
	err := json.Unmarshal([]byte(data), &ts1)
	fmt.Println(err)
	fmt.Println(ts1)
}

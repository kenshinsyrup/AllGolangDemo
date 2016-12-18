package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

var alphabetNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	// 涉及验证码等安全度较高需求的操作时，使用crypto/rand 而不是math/rand
	s, err := GetRandomSring(8, alphabetNum)
	fmt.Println(s, err)

	m := map[string]int{}
	for i := 0; i < 100000; i++ {
		k, err := GetRandomSring(8, alphabetNum)
		if err != nil {
			return
		}
		if v, ok := m[k]; ok {
			fmt.Printf("重复: %s %d \n", k, v)
			break
		}
		m[k] = i
		if i%1000 == 0 {
			fmt.Println(k)
		}
	}
}

func GetRandomSring(num int, s string) (string, error) {
	b := make([]byte, num)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	l := len(s)
	for _, v := range b {
		val := int(v)
		for val > l-1 {
			val = val % l
		}
		buf.WriteString(s[val : val+1])
	}

	return buf.String(), nil
}

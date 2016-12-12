package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	m := map[string]int{}
	for i := 0; i < 1000000; i++ {
		k := GetRandomSring(8)
		if v, ok := m[k]; ok {
			fmt.Printf("重复: %s %d \n", k, v)
			break
		}
		m[k] = i
	}
}

var (
	randSeed = int64(1)
	l        sync.Mutex
)

//获取指定长度的随机字符串
//@params num int 生成的随机字符串的长度
//@params str string 可选，指定随机的字符串
func GetRandomSring(num int, str ...string) string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if len(str) > 0 {
		s = str[0]
	}
	l := len(s)
	r := rand.New(rand.NewSource(getRandSeed()))
	var buf bytes.Buffer
	for i := 0; i < num; i++ {
		x := r.Intn(l)
		buf.WriteString(s[x : x+1])
	}
	return buf.String()
}
func getRandSeed() int64 {
	l.Lock()
	if randSeed >= 100000000 {
		randSeed = 1
	}
	randSeed++
	l.Unlock()
	return time.Now().UnixNano() + randSeed
}

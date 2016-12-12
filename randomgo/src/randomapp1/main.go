package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

var (
	idChars    = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	idCharsLen = int64(len(idChars))
)

func main() {

	getID := func() string {
		var c string
		offset := time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC)
		func() {
			seq := int64(0)
			previousID := int64(0)
			// for {
			inc, err := rand.Int(rand.Reader, big.NewInt(9999))
			if err != nil {
				panic(err)
			}
			if seq >= 99999999 {
				seq = 0
			}
			seq += inc.Int64()
			fmt.Println(time.Since(offset))
			fmt.Println(int64(time.Since(offset)))
			fmt.Println(int64(time.Microsecond))
			id := int64(time.Since(offset))/int64(time.Microsecond)/100*10000 + seq
			if id < previousID {
				id++
			}
			previousID = id
			fmt.Println(id)
			c = base36encode(id)
			// }
		}()
		return c
	}

	fmt.Println(getID())
}

func base36encode(n int64) string {
	if n < 0 {
		panic(fmt.Sprintf("%v", n))
	}
	if n < idCharsLen {
		return string(idChars[n])
	}
	ret := ""
	for n != 0 {
		ret = string(idChars[n%idCharsLen]) + ret
		n = n / idCharsLen
	}
	return ret
}

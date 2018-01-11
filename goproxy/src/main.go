package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8888", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// curl --proxy http://127.0.0.1:8888 http://www.baidu.com
		// Hola!
		// &{GET http://www.baidu.com/ HTTP/1.1 1 1 map[User-Agent:[curl/7.54.0] Accept:[*/*] Proxy-Connection:[Keep-Alive]] {} <nil> 0 [] false www.baidu.com map[] map[] <nil> map[] 127.0.0.1:51118 http://www.baidu.com/ <nil> <nil> <nil> 0xc4201121c0}

		// curl --proxy http://127.0.0.1:8888 https://www.google.com
		// curl: (35) error:140770FC:SSL routines:SSL23_GET_SERVER_HELLO:unknown protocol
		// &{CONNECT //www.google.com:443 HTTP/1.1 1 1 map[User-Agent:[curl/7.54.0] Proxy-Connection:[Keep-Alive]] {} <nil> 0 [] false www.google.com:443 map[] map[] <nil> map[] 127.0.0.1:51122 www.google.com:443 <nil> <nil> <nil> 0xc4201000c0}
		fmt.Println(req)
		rw.Write([]byte(`Hola!`))
	}))
}

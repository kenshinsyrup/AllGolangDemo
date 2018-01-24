package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/url"
)

func main() {
	// 普通的HTTP代理
	// http.ListenAndServe(":8888", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
	// 	// curl --proxy http://127.0.0.1:8888 http://www.baidu.com
	// 	// Hola!
	// 	// &{GET http://www.baidu.com/ HTTP/1.1 1 1 map[User-Agent:[curl/7.54.0] Accept:[*/*] Proxy-Connection:[Keep-Alive]] {} <nil> 0 [] false www.baidu.com map[] map[] <nil> map[] 127.0.0.1:51118 http://www.baidu.com/ <nil> <nil> <nil> 0xc4201121c0}

	// 	// curl --proxy http://127.0.0.1:8888 https://www.google.com
	// 	// curl: (35) error:140770FC:SSL routines:SSL23_GET_SERVER_HELLO:unknown protocol
	// 	// &{CONNECT //www.google.com:443 HTTP/1.1 1 1 map[User-Agent:[curl/7.54.0] Proxy-Connection:[Keep-Alive]] {} <nil> 0 [] false www.google.com:443 map[] map[] <nil> map[] 127.0.0.1:51122 www.google.com:443 <nil> <nil> <nil> 0xc4201000c0}
	// 	fmt.Println(req)
	// 	rw.Write([]byte(`Hola!`))
	// }))

	// TCP代理，可以代理HTTP和HTTPS
	s, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Fail to Listen: ", err)
	}
	for {
		cToP, err := s.Accept()
		if err != nil {
			fmt.Println("Fail to Accept: ", err)
			continue
		}
		go handle(cToP)
	}
}

func handle(cToP net.Conn) {
	var b [1024]byte
	n, err := cToP.Read(b[:])
	if err != nil {
		fmt.Println("Fail to Read: ", err)
		return
	}
	fmt.Println("******")
	fmt.Println(string(b[:n]))

	var method, host string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	fmt.Println("method: ", method)
	fmt.Println("host: ", host)

	hostPortURL, err := url.Parse(host)
	if err != nil {
		fmt.Println("Fail to Parse: ", err)
		return
	}
	address := hostPortURL.Scheme + ":443"
	//获得了请求的host和port，就开始拨号吧
	pToS, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}

	if method == "CONNECT" {
		fmt.Fprint(cToP, "HTTP/1.1 200 Connection established\r\n\r\n")
	}

	//进行转发
	// go io.Copy(pToS, cToP)
	// io.Copy(cToP, pToS)

	// 读取传输的数据
	complete := make(chan bool)
	go copyContent(cToP, pToS, complete)
	go copyContent(pToS, cToP, complete)
	// Block until we've completed!
	<-complete
}

func copyContent(from net.Conn, to net.Conn, complete chan bool) {
	var err error = nil
	var bytes []byte = make([]byte, 256)
	var read int = 0
	for {
		// Read data from the source connection.
		read, err = from.Read(bytes)
		// If any errors occured, write to complete as we are done (one of the
		// connections closed.)
		if err != nil {
			complete <- true
			break
		}
		fmt.Println("from content: ", string(bytes))
		// Write data to the destination.
		_, err = to.Write(bytes[:read])
		// Same error checking.
		if err != nil {
			complete <- true
			break
		}
	}
}

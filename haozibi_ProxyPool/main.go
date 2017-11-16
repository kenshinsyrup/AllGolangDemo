package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	proxy "github.com/haozibi/ProxyPool"
)

func main() {
	// serve
	// l, err := net.Listen("tcp", ":8888")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// fmt.Println("Serving...")
	// for {
	// 	client, err := l.Accept()
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// 	fmt.Println("client: ", client)
	// 	go handleClientRequest(client)
	// }

	// destination server, source server, microsoft academic
	go http.ListenAndServe(":6666", &sourceServerHandler{})
	time.Sleep(1 * time.Second)

	// host server,  public server
	go http.ListenAndServe(":7777", &hostServerHandler{})
	time.Sleep(1 * time.Second)

	// proxy server,  listen to crawler's request
	go http.ListenAndServe(":8888", &proxyServerHandler{})
	time.Sleep(1 * time.Second)

	// crwaler request
	res, err := http.DefaultClient.Get("http://127.0.0.1:8888")
	fmt.Println("get err: ", err)
	fmt.Println("res status: ", res.Status)

	// client request
	// res, err := http.DefaultClient.Get("http://127.0.0.1:9090")

	// normal original request and response
	// request, _ := http.NewRequest("GET", "http://127.0.0.1:8888", nil)
	// res, _ := http.DefaultClient.Do(request)
	// fmt.Println("&&&&&&&&&")

	// now let crawlers request from proxy server
	// request, _ := http.NewRequest("GET", "http://127.0.0.1:9090", nil)
	// res, _ := http.DefaultClient.Do(request)
	// res, err := http.DefaultClient.Get("http://127.0.0.1:8888")
	// fmt.Println("get err: ", err)
	// fmt.Println("res status: ", res.Status)

	// getProxy()

	// compare()

	// why remote addr is empty
	// req, err := http.NewRequest("GET", "http://google.com", nil)
	// fmt.Println("req err: ", err)
	// res, err := http.DefaultClient.Do(req)
	// fmt.Println("res err: ", err)
	// fmt.Println("header: ", res.Header)
	// fmt.Println("code: ", res.StatusCode)
	// fmt.Println("res req remoteaddr: ", res.Request.RemoteAddr)
	// fmt.Println("req remoteaddr: ", req.RemoteAddr)

}

func compare() {
	// reg, _ := regexp.Compile("您的IP地址是：\\[.+?\\]")

	req2, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		fmt.Println("new request is err: ", err)
		return
	}
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			log.Printf("resolved to: %s", connInfo.Conn.RemoteAddr())
		},
	}
	req2 = req2.WithContext(httptrace.WithClientTrace(req2.Context(), trace))
	res, err := http.Get("http://google.com/") //请求并获取到对象
	// client2 := &http.Client{}
	// res, err := client2.Do(req2)

	// fmt.Println("dataproxy data: ", string(dataproxy))
	fmt.Println("res req: ", res.Request)
	data, err := ioutil.ReadAll(res.Body) //取出主体的内容
	if err != nil {
		log.Fatal(err)
	}
	_ = data
	fmt.Println("req2 ip: ", req2.RemoteAddr, " request URI: ", req2.RequestURI)
	// fmt.Println("data data: ", string(data)) //打印

	fmt.Println("&&&&&&&&&&&&&&")

	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		fmt.Println("new request is err: ", err)
		return
	}
	trace = &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			log.Printf("resolved to: %s", connInfo.Conn.RemoteAddr())
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	fmt.Println("req ", req.RemoteAddr, " ", req.RequestURI)
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	fmt.Println("ip ", ip)
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://208.139.7.6:80") //根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	// resp, err := client.Get("http://google.com/") //请求并获取到对象,使用代理
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	dataproxy, err := ioutil.ReadAll(resp.Body) //取出主体的内容
	if err != nil {
		log.Fatal(err)
	}
	_ = dataproxy
	fmt.Println("req ip: ", req2.RemoteAddr, " request URI: ", req2.RequestURI)
	fmt.Println("resp req: ", resp.Request)
	fmt.Println("headers: ", resp.Request.Header)

	// sproxy := reg.FindString(string(dataproxy))
	// s := reg.FindString(string(data))
	// res.Body.Close()
	// resp.Body.Close()
	// fmt.Printf("不使用代理:%s", s)     //打印
	// fmt.Printf("使用代理:%s", sproxy) //打印
}

func saySth(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("req header: ", req.Header)
	rw.Write([]byte(`say sth im giving up on you`))
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("client request method: %s host: %s adress: %s \n", method, host, address)
	if hostPortURL.Opaque == "443" { //https访问
		address = hostPortURL.Scheme + ":443"
	} else { //http访问
		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}
	//获得了请求的host和port，就开始拨号吧
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(b[:n])
	}
	fmt.Println(server.LocalAddr())
	fmt.Println(server.RemoteAddr())
	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}

func getProxy() {
	// 是否输出调试信息，true输出，false不输出。默认不输出
	proxy.Setting(true)

	// 自定义ip测试连接，默认 http://www.baidu.com
	proxy.TestUrl = "http://www.baidu.com"

	// 自定义线程数，默认50
	proxy.ProxyProNum = 100

	// 启动程序
	proxy.Start()

	for {
		// GetProxy() 返回 175.171.246.195:8118 格式
		fmt.Println("Get ", proxy.GetProxy())
	}
}

type sourceServerHandler struct{}

func (sourceH *sourceServerHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("source server received req addr: ", req.RemoteAddr)
	rw.Write([]byte(`say sth im giving up on you`))
}

type hostServerHandler struct{}

func (hostH *hostServerHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("host server received req addr: ", req.RemoteAddr)
	fmt.Println("host server received req url: ", req.URL)
	// use a reverse proxy to forward received original req
	rp := httputil.NewSingleHostReverseProxy(req.URL)
	rp.ServeHTTP(rw, req)
}

type proxyServerHandler struct{}

func (proxyH *proxyServerHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("proxy server receive req addr: ", req.RemoteAddr)

	// set proxy, here we use a 'public server' :7777
	proxyURL, err := url.Parse("http://127.0.0.1:7777")
	if err != nil {
		fmt.Println("parse url err: ", err)
		return
	}
	proxyClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	// proxy server know where the crawler want to go, here is :6666
	newProxyReq, err := http.NewRequest(req.Method, "http://127.0.0.1:6666", req.Body)
	if err != nil {
		fmt.Println("new proxy request err: ", err)
		return
	}
	// here's proxyed response
	res, err := proxyClient.Do(newProxyReq)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("status: ", res.Status)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("readall err: ", err)
	fmt.Println("body: ", string(data))
	time.Sleep(2 * time.Second)

	// use a reverse proxy to forward received original req
	// rp := httputil.NewSingleHostReverseProxy(req.URL)
	// rp.ServeHTTP(rw, req)

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println("proxyserver Do req err: ", err)
	// 	return
	// }
	// fmt.Println("proxyserver res status: ", res.Status)
	// // write res out to rw
	// res.Write(rw)
}

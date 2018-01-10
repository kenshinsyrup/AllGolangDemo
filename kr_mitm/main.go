package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kr/mitm"
)

var hostname, _ = os.Hostname()

func main() {
	fmt.Println("hostname: ", hostname)

	certPEM, keyPEM, err := mitm.GenCA(hostname)
	if err != nil {
		fmt.Println("gen ca err: ", err)
		return
	}

	err = ioutil.WriteFile("server.crt", certPEM, 0777)
	if err != nil {
		fmt.Println("write to file with certPEM err: ", err)
		return
	}
	err = ioutil.WriteFile("server.key", keyPEM, 0777)
	if err != nil {
		fmt.Println("write to file with keyPEM err: ", err)
		return
	}

	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		fmt.Println("get cert err: ", err)
		return
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])

	p := &mitm.Proxy{
		CA: &cert,
	}

	http.ListenAndServe(":8888", p)

}

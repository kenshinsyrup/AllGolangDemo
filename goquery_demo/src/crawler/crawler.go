package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Run() {
	// 快代理
	fmt.Println(getAddrsFromKuaidaili())
}

func getBody(dstURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", dstURL, nil)
	if err != nil {
		log.Println("http.NewRequest err: ", err)
		return nil, err
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Fail to do request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Fail to read response body: ", err)
		return nil, err
	}
	return body, nil
}

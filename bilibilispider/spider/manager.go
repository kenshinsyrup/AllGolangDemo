package spider

import (
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

func GetHTML(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		glog.Errorln("http.NewRequest error: ", err)
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36")
	client := http.DefaultClient
	res, e := client.Do(req)
	if e != nil {
		glog.Errorln("client.Do error: ", err)
		panic(err)
	}
	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()
		bodyByte, _ := ioutil.ReadAll(body)
		return string(bodyByte)
	}
	return ""
}

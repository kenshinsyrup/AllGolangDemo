package extract

import (
	"fmt"
	"net/http"
	"strings"

	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

func ParseHTML(res *http.Response) {
	glog.Infoln("ParseHTML start...")
	// doc, err := goquery.NewDocumentFromResponse(res)
	// if err != nil {
	// 	glog.Infoln("NewDocumentFromResponse error: ", err)
	// }
	// glog.Infoln(doc.Text())
	bodyByte, _ := ioutil.ReadAll(res.Body)
	bodyStr := string(bodyByte)
	// glog.Infoln(bodyStr)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(bodyStr))
	if err != nil {
		glog.Infoln("error: ", err)
		return
	}
	s := doc.Find("div.b-page-body").Find("div.video-list.list-c")
	fmt.Println(s.Length())

	s2 := doc.Find("div.vd-list-cnt loaded")
	fmt.Println(s2.Length())
}

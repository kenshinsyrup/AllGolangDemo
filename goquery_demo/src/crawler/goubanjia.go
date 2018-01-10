package crawler

import (
	"bytes"
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

var goubanjiaDisplayNonePattern = regexp.MustCompile(`(?U)(<p.*none.*>.*<\/p>)`)
var goubanjiaTagPattern = regexp.MustCompile(`(?U)<.*>`)

func getAddrsFromGoubanjia() ([]string, error) {
	page := 1
	gobanjiaURL := fmt.Sprintf("http://www.goubanjia.com/index%d.shtml", page)
	body, err := getBody(gobanjiaURL)
	if err != nil {
		log.Println("Fail to parse gobanjia: ", err)
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println("Fail to new goquery document: ", err)
		return nil, err
	}
	tableList := doc.Find("#list")
	table := tableList.Find(".table")
	var addrs []string
	table.Children().Children().Each(func(rowNum int, row *goquery.Selection) {
		// skip the table header
		if rowNum == 0 {
			return
		}
		rh, err := row.Find(".ip").Html()
		if err != nil {
			return
		}
		noNoneContent := goubanjiaDisplayNonePattern.ReplaceAllString(rh, "")
		noTagContent := goubanjiaTagPattern.ReplaceAllString(noNoneContent, "")
		addrs = append(addrs, noTagContent)
	})
	return addrs, nil
}

package crawler

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func getAddrsFromData5u() []string {
	log.Println("getAddrsFromData5u start at: ", time.Now())

	data5uURLs := []string{
		"http://www.data5u.com/free/gngn/index.shtml",
		"http://www.data5u.com/free/gnpt/index.shtml",
		"http://www.data5u.com/free/gwgn/index.shtml",
		"http://www.data5u.com/free/gwpt/index.shtml",
	}

	var addrs []string
	for _, data5uURL := range data5uURLs {
		ips, err := parsedata5uAddrs(data5uURL)
		if err != nil {
			log.Printf("Fail to parse data5u %s: %v\n", data5uURL, err)
			continue
		}
		addrs = append(addrs, ips...)
	}

	log.Printf("getAddrsFromData5u end at: %v, get %d items.\n", time.Now(), len(addrs))
	return addrs
}

func parsedata5uAddrs(data5uURL string) ([]string, error) {
	body, err := getBody(data5uURL)
	if err != nil {
		log.Println("Fail to get data5u body: ", err)
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println("Fail to new goquery document: ", err)
		return nil, err
	}
	table := doc.Find(".wlist ul li[style='text-align:center;']")
	var addrs []string
	table.Children().Each(func(rowNum int, row *goquery.Selection) {
		// skip the table header
		if rowNum == 0 {
			return
		}
		var ip, port string
		row.Children().Each(func(num int, span *goquery.Selection) {
			switch num {
			case 0:
				ip = span.Text()
			case 1:
				port = span.Text()
			default:
				return
			}
		})
		addrs = append(addrs, fmt.Sprintf("%s:%s", ip, port))
	})
	return addrs, nil
}

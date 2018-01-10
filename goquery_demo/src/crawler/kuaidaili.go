package crawler

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// http://www.kuaidaili.com/free/inha/1/
// http://www.kuaidaili.com/free/intr/1/

func getAddrsFromKuaidaili() []string {
	log.Println("getAddrsFromKuaidaili start at: ", time.Now())
	var addrs []string
	kuaidailiURLs := []string{
		"http://www.kuaidaili.com/free/inha/",
		"http://www.kuaidaili.com/free/intr/",
	}
	for _, kuaidailiURL := range kuaidailiURLs {
		var page int
		for {
			page++
			if page > 10 {
				break
			}
			ips, err := parseKuaidailiAddrs(fmt.Sprintf("%s%d/", kuaidailiURL, page))
			if err != nil {
				log.Println("Fail to parse kuaidaili: ", err)
				continue
			}
			addrs = append(addrs, ips...)
			// here sleep 1 second, or kuaidaili will response 502
			time.Sleep(1 * time.Second)
		}
	}

	log.Printf("getAddrsFromKuaidaili end at: %v, get %d items.\n", time.Now(), len(addrs))
	return addrs
}

func parseKuaidailiAddrs(kuaidailiURL string) ([]string, error) {
	body, err := getBody(kuaidailiURL)
	if err != nil {
		log.Println("Fail to get kuaidaili body: ", err)
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Println("Fail to new goquery document: ", err)
		return nil, err
	}
	tableList := doc.Find("#list")
	tBody := tableList.Find("table tbody")
	var addrs []string
	tBody.Children().Each(func(rowNum int, row *goquery.Selection) {
		var IP, Port string
		row.Children().Each(func(tdNum int, td *goquery.Selection) {
			switch tdNum {
			case 0:
				IP = td.Text()
			case 1:
				Port = td.Text()
			default:
				return
			}
		})
		addrs = append(addrs, fmt.Sprintf("%s:%s", IP, Port))
	})
	return addrs, nil
}

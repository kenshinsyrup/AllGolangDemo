package main

import (
	"allgolangdemo/bilibilispider/psql"
	"allgolangdemo/bilibilispider/spider"

	"github.com/jinzhu/gorm"

	"encoding/json"
	"fmt"
	"strings"
)

type PageInfo struct {
	Data PageData `json:"data"`
}

type PageData struct {
	Archives []VideoInfo2 `json:"archives"`
}

type VideoInfo2 struct {
	ID          int64  //主键
	Aid         int    `json:"aid"`         //番号
	Author      string `json:"author"`      //作者
	Description string `json:"description"` //描述
	Create      string `json:"create"`      //创作时间
	Title       string `json:"title"`       //标题
	Play        int    `json:"play"`        //播放数
	Favorites   int    `json:"favorites"`   //收藏数
}

// 暂不使用
type VideoStat struct {
	Coin string `json:"coin"` //投喂硬币数
}

var jQueryStr string = "jQuery17209494025525031935_1485407014321"

func main() {
	// 数据库
	host := "localhost"
	userName := "psqluser"
	dbName := "bilibilidb"
	pwd := "psqlpwd"
	db := psql.SetupDB(host, userName, dbName, pwd)
	if !db.HasTable(&VideoInfo2{}) {
		db.CreateTable(&VideoInfo2{})
	}
	db.AutoMigrate(&VideoInfo2{})
	defer db.Close()

	var pn int
	// for {
	// 	// wg := &sync.WaitGroup{}
	// 	var mutex sync.Mutex
	// 	for i := 0; i < 5; i++ {
	// 		// wg.Add(1)
	// 		go func() {
	// 			// defer wg.Done()
	// 			mutex.Lock()
	// 			pn++
	// 			mutex.Unlock()
	// 			fmt.Println("current page: ", pn)
	// 			crawl(pn, db)
	// 		}()
	// 	}
	// 	// wg.Wait()
	// }

	queue := make(chan int, 5)
	for {
		for i := 0; i < 5; i++ {
			pn++
			go crawl(pn, db, queue)
		}

		for done := range queue {
			fmt.Printf("page %d done", done)
		}
	}
	fmt.Println("all done")
}

func getURL(pn int) string {
	return fmt.Sprintf("https://api.bilibili.com/archive_rank/getarchiverankbypartion?callback=%s&type=jsonp&tid=20&pn=%d&_=1485407014635", jQueryStr, pn)
}

func crawl(pn int, db *gorm.DB, queue chan int) {
	var pageInfo PageInfo
	jsonStr := spider.GetHTML(getURL(pn))
	if strings.Contains(jsonStr, `"code":-40002`) {
		close(queue)
		return
	}
	// remove useless prefix&suffix
	jsonStr = strings.TrimPrefix(jsonStr, jQueryStr+"(")
	jsonStr = strings.TrimSuffix(jsonStr, ");")
	err := json.Unmarshal([]byte(jsonStr), &pageInfo)
	if err != nil {
		fmt.Println("err jsonStr: ", jsonStr)
		fmt.Println("json.Unmarshal error: ", err)
		return
	}

	for _, v := range pageInfo.Data.Archives {
		fmt.Println("av id:", v.ID)
		// if v.Play >= 100000 || v.Favorites > 10000 {
		// 	fmt.Println("******************************")
		// 	fmt.Println(v)
		// 	db.Create(&v)
		// }
	}
	queue <- pn
}

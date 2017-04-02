package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/zipdownload", zipHandler)
	log.Println("Listening...")
	http.ListenAndServe(":9999", nil)
}

func zipHandler(rw http.ResponseWriter, r *http.Request) {
	zipName := "ZipTest.zip"
	// 设置rw的header信息中的ctontent-type，对于zip可选以下两种
	// rw.Header().Set("Content-Type", "application/octet-stream")
	rw.Header().Set("Content-Type", "application/zip")
	// 设置rw的header信息中的Content-Disposition为attachment类型
	rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipName))
	// 向rw中写入zip文件
	err := getZip(rw)
	if err != nil {
		log.Fatal(err)
	}
}

func getZip(w io.Writer) error {
	// 创建zip.Writer
	zipW := zip.NewWriter(w)
	defer zipW.Close()

	for i := 0; i < 5; i++ {
		// 向zip中添加文件
		f, err := zipW.Create(strconv.Itoa(i) + ".txt")
		if err != nil {
			return err
		}
		// 向文件中写入文件内容
		_, err = f.Write([]byte(fmt.Sprintf("Hello file %d", i)))
		if err != nil {
			return err
		}
	}
	return nil
}

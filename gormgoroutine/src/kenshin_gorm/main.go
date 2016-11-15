package main

import (
	"log"
	"net/http"

	"fmt"

	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var count int

func main() {
	fmt.Println("main")
	var err error
	sqlDriver := "postgres"
	sqlParams := "user=Kenshin password=psql dbname=postgres host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(sqlDriver, sqlParams)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Data{})

	m := NewManager(db)

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		data := Data{
			ID:      count,
			Content: "test",
		}
		fmt.Println("data: ", data)

		go func(data *Data) {
			time.Sleep(time.Second)
			fmt.Println("go")
			if err = m.insert(data); err != nil {
				fmt.Println(err)
			}
		}(&data)

		fmt.Println("hello")
		_, err = w.Write([]byte("hello"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("should return")
		return
		fmt.Println("already return")
	}))
}

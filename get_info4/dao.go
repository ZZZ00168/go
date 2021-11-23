package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
 
var db *sql.DB

func insertDB(v_url string, p_url string, title string) {
 
	connStr := "root:123456@tcp(119.91.214.40)/gotest?charset=utf8"
	if db == nil {
		db, _ = sql.Open("mysql", connStr)
	}
	
	_, err2 := db.Exec("INSERT INTO url(video_url,pic_url,title) VALUES(?,?,?)", v_url, p_url, title)

	if err2 != nil {
		log.Fatal(err2.Error)
		return
	}
	log.Printf("入库数据" + v_url)

}

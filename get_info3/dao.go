package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
 

func insertDB(v_url string, p_url string, title string) {
 
	connStr := "root:123456@tcp(119.91.214.40)/gotest?charset=utf8"
	db, err := sql.Open("mysql", connStr)
	
	if err != nil {
		log.Fatal(err.Error)
		return
	}

	_, err2 := db.Exec("INSERT INTO url(video_url,pic_url,title) VALUES(?,?,?)", v_url, p_url, title)

	if err2 != nil {
		log.Printf("入库失败")
		return
	}
	log.Printf("入库数据" + v_url)

}

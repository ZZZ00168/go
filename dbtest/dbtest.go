package main
 
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
 
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var PAGE_LIMIT int
 
type UrlData struct {
	Video_url      string    `json:"Video_url"`
	Pic_url    string `json:"Pic_url"`
	Title     string    `json:"Title"`
}

func main() {
	PAGE_LIMIT=10

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.LoadHTMLGlob("index.html")
	r.GET("/", Index)
	r.Run(":9090")
}

func Index(c *gin.Context) {
 
	connStr := "root:123456@tcp(119.91.214.40)/gotest?charset=utf8"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error)
		return
	}

	page := c.DefaultQuery("page", "0")
	pageIndex, paserErr := strconv.Atoi(page)
	if paserErr != nil{
		log.Fatal(paserErr.Error)
	}

	limitStart := pageIndex * PAGE_LIMIT
	limitEnd := (pageIndex + 1) * PAGE_LIMIT 
	rows, errq := db.Query("select video_url,pic_url,title from url limit ?,?", limitStart, limitEnd)

	if errq != nil {
		log.Printf("error:")
		log.Fatal(errq.Error)
		return
	}
 
	var urls []UrlData
	
	for rows.Next() {
		var u UrlData
        
		errn := rows.Scan(&u.Video_url, &u.Pic_url, &u.Title)
		if errn != nil {
			fmt.Printf("%v", errn)
		}
		
		urls = append(urls, u)
	}
 	
	c.HTML(http.StatusOK, "index.html", gin.H{"res": urls, "pageIndex": pageIndex})
 
}


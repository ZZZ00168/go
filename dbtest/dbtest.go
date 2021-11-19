package main
 
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
 
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)
 
type UrlData struct {
	Video_url      string    `json:"Video_url"`
	Pic_url    string `json:"Pic_url"`
	Title     string    `json:"Title"`
}
 
func Index(c *gin.Context) {
 
	connStr := "root:123456@tcp(119.91.214.40)/gotest?charset=utf8"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error)
		return
	}

	rows, errq := db.Query("select video_url,pic_url,title from url")
	if errq != nil {
		log.Printf("error:")
		log.Fatal(errq.Error)
		return
	}
	log.Printf("run2")
 
	var urls []UrlData
	

	for rows.Next() {
		var u UrlData
        
		errn := rows.Scan(&u.Video_url, &u.Pic_url, &u.Title)
		if errn != nil {
			fmt.Printf("%v", errn)
		}
		
		urls = append(urls, u)
	}
 	
	log.Printf("run3")
	c.HTML(http.StatusOK, "index.html", gin.H{"res": urls})
 
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("index.html")
	r.GET("/", Index)
	r.Run(":9090")
}

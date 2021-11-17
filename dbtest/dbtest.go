package main
 
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
 
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)
 
type User struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Addtime string `json:"Addtime`
}
 
func Index(c *gin.Context) {
 
	connStr := "root:123456@tcp(127.0.0.1)/gotest?charset=utf8"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error)
		return
	}
 	log.Printf("run1")

	rows, errq := db.Query("select id,name,age,addtime from go_users")
	if errq != nil {
		log.Printf("error:")
		log.Fatal(errq.Error)
		return
	}
	log.Printf("run2")
 
	var users []User
	

	for rows.Next() {
		var u User
        
		errn := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Addtime)
		if errn != nil {
			fmt.Printf("%v", errn)
		}
		
		users = append(users, u)
	}
 	
	log.Printf("run3")
	c.HTML(http.StatusOK, "index.html", gin.H{"res": users})
 
}
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("index.html")
	r.GET("/", Index)
	r.Run(":9090")
}

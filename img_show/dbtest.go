package main
 
import (
	"net/http"
 
	"github.com/gin-gonic/gin"
)
 
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "img_show.html", gin.H{})
}

func main() {
	r := gin.Default()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.Static("/img", "./img")
	r.StaticFS("/static", http.Dir("./static"))
	r.LoadHTMLGlob("img_show.html")
	r.GET("/", Index)
	r.Run(":9090")
}

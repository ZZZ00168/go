package main

import (
		"fmt"
		"time"
		"strconv"
)

var mainUrl string

func main() {
        start := time.Now()
        mainUrl = "http://a.0516meiya.com"
        
        parseUrls(mainUrl + "/type/1"+ ".html")
        for i:=2 ; i<10 ; i++ {
   				parseUrls(mainUrl + "/type/1-" + strconv.Itoa(i) + ".html")
        }
        
        elapsed := time.Since(start)
        fmt.Printf("Took %s", elapsed)
}
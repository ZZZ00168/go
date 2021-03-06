package main

import (
		"fmt"
		"time"
		"strconv"
)
func main() {
        start := time.Now()
        for i := 0; i < 10; i++ {
                parseUrls("https://movie.douban.com/top250?start=" + strconv.Itoa(25 * i))
        }
        elapsed := time.Since(start)
        fmt.Printf("Took %s", elapsed)
}
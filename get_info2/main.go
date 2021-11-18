package main

import (
		"fmt"
		"time"
)
func main() {
        start := time.Now()
        parseUrls("http://a.0516meiya.com")
        elapsed := time.Since(start)
        fmt.Printf("Took %s", elapsed)
}
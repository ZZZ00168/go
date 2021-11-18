package main

import (
		"fmt"
		"time"
)
func main() {
        start := time.Now()
        parseUrls("http://www.pornhub.com")
        elapsed := time.Since(start)
        fmt.Printf("Took %s", elapsed)
}
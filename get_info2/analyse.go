package main

import (
		"fmt"
		
		"github.com/axgle/mahonia"
)

func parseUrls(url string) {
	body := fetch(url)
	fmt.Println(body)
	
	bodystr := mahonia.NewDecoder("utf-8").ConvertString(string(body))
  fmt.Println(bodystr)
    
}
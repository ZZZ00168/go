package main

import (
		"fmt"
		"regexp"


		"github.com/axgle/mahonia"
)

func parseUrls(url string) {
	body := fetch(url)
	fmt.Println(body)
	fmt.Println("----------------\n\n\n\n")
	bodystr := mahonia.NewDecoder("utf-8").ConvertString(string(body))
	fmt.Println(bodystr)
	
	urls := regexp.MustCompile(`(I\("(.*)")`)
	items := urls.FindAllStringSubmatch(body, -1)
	for _, item := range items {
		fmt.Println(item[2])
	}
	
	for _, item := range items {
		fmt.Println(decode(item[2]))
	}

}

func decode(s string) string{
	n := ""
	for _, c := range s{
		n += string(128^c);
	}
	return n
}
package main

import (
		"fmt"
		"regexp"


		
)

func parseUrls(url string) {
	body := fetch(url)
	fmt.Println(url)
// "github.com/axgle/mahonia"
//	fmt.Println("----------------\n\n\n\n")
//	bodystr := mahonia.NewDecoder("utf-8").ConvertString(string(body))
//	fmt.Println(bodystr)

//	fmt.Println("----------------\n\n\n\n TEST")
	divmatch := regexp.MustCompile(`<a class="vodbox"[\s\S]*?</a>`)



	items := divmatch.FindAllStringSubmatch(body, -1)
	for _, item := range items {
		// fmt.Println(item[0])

		videoUrl := getUrl(item[0])
		title := getTitle(item[0])

		fmt.Println(videoUrl)
		fmt.Println(title)
		fmt.Println("\n")
	}


}

func getUrl(s string) string{
	// .any char *multi ?min match
	hrefMatch := regexp.MustCompile(`href=".*?"`)

	href := hrefMatch.FindAllStringSubmatch(s,-1)
	hrefStr := string(href[0][0])
	subUrl := hrefStr[6:len(hrefStr)-1]
	url := mainUrl + subUrl
	return url
}

func getTitle(s string) string{
	// .any char *multi ?min match
	titleMatch := regexp.MustCompile(`document.write\(I\(".*?"`)

	matches := titleMatch.FindAllStringSubmatch(s,-1)
	matchStr := string(matches[0][0])
	title := matchStr[18:len(matchStr)-1]
	title = decode(title)
	return title
}


func decode(s string) string{
	n := ""
	for _, c := range s{
		n += string(128^c);
	}
	return n
}
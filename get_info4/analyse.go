package main

import (
		"fmt"
		"regexp"
)

func parseUrls(url string) {
	body := fetch(url)
	fmt.Println(url)

	divmatch := regexp.MustCompile(`<a class="vodbox"[\s\S]*?</a>`)

	items := divmatch.FindAllStringSubmatch(body, -1)
	for _, item := range items {
		// fmt.Println(item[0])

		videoUrl := getUrl(item[0])
		title := getTitle(item[0])
		picUrl := getPicUrl(item[0])

		fmt.Println(getRealVideoUrl(videoUrl))
		fmt.Println(picUrl)
		fmt.Println(title)
		fmt.Println("\n")
		insertDB(videoUrl, picUrl, title)
	}
}

func getRealVideoUrl(s string) string{
	body := fetch(s)
	reg := regexp.MustCompile(`src=".*?index.m3u8"`)
	matchs := reg.FindAllStringSubmatch(body,-1)
	match := string(matchs[0][0])
	url := match[5:len(match)-1]
	return url
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

func getPicUrl(s string) string{
	titleMatch := regexp.MustCompile(`"https://.*.jpg"`)
	matchs := titleMatch.FindAllStringSubmatch(s,-1)
	match := string(matchs[0][0])
	url := match[1:len(match)-1]
	return url
}

func decode(s string) string{
	n := ""
	for _, c := range s{
		n += string(128^c);
	}
	return n
}
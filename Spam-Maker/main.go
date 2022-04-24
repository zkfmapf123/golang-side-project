package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var MASK rune = '*'

func main() {
	bytes, err := ioutil.ReadFile("spam.txt")

	if err != nil {
		panic(err)
	}

	urls := strings.Split(string(bytes), "\n")

	for _, url := range urls {
		printMaskUrl(url)

	}
}

func printMaskUrl(url string) {

	rs := []rune(url)
	startIndex, endIndex := getMaskStart(rs)
	urls := make([]rune, endIndex-startIndex)

	for i, r := range rs[startIndex:endIndex] {
		urls[i] = r
	}

	splitUrls := strings.Split(string(urls), ".")
	newEndIndex := len(splitUrls[0]) + len(splitUrls[1])

	// fmt.Println(startIndex+len(splitUrls[0]), startIndex+newEndIndex)

	for i, r := range rs {

		if i > startIndex+len(splitUrls[0]) && i <= startIndex+newEndIndex {
			fmt.Printf("%c", MASK)
		} else {
			fmt.Printf("%c", r)
		}

	}
	fmt.Println()
}

func getMaskStart(urls []rune) (int, int) {

	var count int = 0
	var index int
	for i, url := range urls {

		if url == '/' {
			count += 1
		}

		if count == 2 {
			index = i
			break
		}
	}

	return index + 1, len(urls)
}

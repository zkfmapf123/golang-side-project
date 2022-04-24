package main

import (
	"fmt"
	"io/ioutil"
)

const MAX_CHAR = 40

func main() {
	rs := readFile("text.txt")

	for i, r := range rs {
		if i%MAX_CHAR == 0 {
			fmt.Println()
		}

		fmt.Printf("%c", r)
	}
}

func readFile(fileName string) []rune {
	buf, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	r := []rune(string(buf))
	return r
}

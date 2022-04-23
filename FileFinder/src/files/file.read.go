package files

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

const FILE_SIZE = 256

func ReadFile(args []string) {

	if len(args) == 0 {
		fmt.Println("Does Not Exist")
		os.Exit(0)
	}

	// Return []fs.FileInfo
	// go doc fs.FileInfo
	files, err := ioutil.ReadDir(args[0])

	// names := getFileSpaceToHuge(files)
	names := getFileSpaceToProfit(files)

	if err != nil {
		panic(err)
	}

	// Read All Files in Path
	// only txt
	// go doc strings
	for _, file := range files {
		fileName := file.Name()
		if strings.Contains(fileName, "txt") {
			fmt.Printf("name : %s\t size : %dKB\n", fileName, file.Size())

			// slice
			// string -> ... -> byte[]
			names = append(names, fileName...)
			writeFile(fileName, names)
		}
	}
}

// https://chmod-calculator.com/
func writeFile(fileName string, writeData []byte) {
	err := ioutil.WriteFile(fileName, writeData, 666)

	if err != nil {
		panic(err)
	}
}

// 1. refactoring -> too much space
func getFileSpaceToHuge(files []fs.FileInfo) []byte {
	fileLengthTotal := len(files) * FILE_SIZE
	fmt.Printf("Total required space : %d bytes\n", fileLengthTotal)
	names := make([]byte, 0, fileLengthTotal)
	return names
}

// 2. refactoring -> profit space
func getFileSpaceToProfit(files []fs.FileInfo) []byte {
	var count = 0
	for _, file := range files {
		if strings.Contains(file.Name(), "txt") {
			count = count + 1
		}
	}

	fmt.Printf("Total required space : %d bytes\n", count*FILE_SIZE)

	return make([]byte, 0, count*FILE_SIZE)
}

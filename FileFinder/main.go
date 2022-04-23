package main

import (
	"os"
	"zkfmapf123/src/files"
)

func main() {
	files.ReadFile(os.Args[1:])
}

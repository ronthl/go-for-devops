package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fileName := "/abc/123/main.go"

	absPath, err := filepath.Abs(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Absolute path: %s\n", absPath) // absPath = /abc/123/main.go

	fileName = "abc/123/main.go"
	absPath, err = filepath.Abs(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Absolute path: %s\n", absPath) // absPath = /your/working/directory/path/abc/123/main.go

	basePath := "/a/b/c"
	targPath := "/a/b/c/d/e/1.txt"

	relPath, err := filepath.Rel(basePath, targPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Relative path: %s\n", relPath) // relPath = d/e/1.txt
}

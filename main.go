package main

import (
	"io"
	"os"
	"path/filepath"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	filePath := filepath.Join(workingDir, "config", "config.json")

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Stream the file content to stdout to show it worked
	io.Copy(os.Stdout, file)
}

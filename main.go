package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./index.html")
	if err != nil {
		fmt.Printf("Got error with detail: %v\n", err)
		return
	}

	fmt.Printf("File content:\n%s\n", string(data))
}

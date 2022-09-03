package main

import (
	"fmt"
	"os"
)

func main() {
	data := "Hello, World!\nThis is the demo for writing local files."
	err := os.WriteFile("./index.txt", []byte(data), 0774)
	if err != nil {
		fmt.Printf("Error occurs with detail:\n%v\n", err)
	}
}

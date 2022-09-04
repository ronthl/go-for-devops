package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	const url = "https://martinfowler.com/"
	const filePath = "./website.html"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := client.Do(req)

	// resp contains an io.ReadCloser that we can read as a file.
	// Let's use io.ReadAll() to read the entire content to data.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Write the data into the file
	err = os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		panic(err)
	}

	// Open the file and write its content to stdout to show it worked
	fileData, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, fileData)
}

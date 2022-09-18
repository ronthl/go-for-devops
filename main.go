package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type record []string

func (rec record) validate() error {
	if len(rec) != 2 {
		return errors.New("data format is incorrect")
	}
	return nil
}

func (rec record) first() string {
	return rec[0]
}

func (rec record) last() string {
	return rec[1]
}

func main() {
	records, err := readRecords()
	if err != nil {
		panic(err)
	}

	for i, rec := range records {
		if i != 0 {
			fmt.Println("*******")
		}
		fmt.Printf("First Name: %s\nLast Name: %s\n", rec.first(), rec.last())
	}
}

func readRecords() ([]record, error) {
	byteData, err := os.ReadFile("data.csv")
	if err != nil {
		return nil, err
	}

	content := string(byteData)
	lines := strings.Split(content, "\n")

	var records []record

	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var rec record = strings.Split(line, ",")
		if err := rec.validate(); err != nil {
			return nil, fmt.Errorf("entry at line %d was invalid: %w", i, err)
		}

		records = append(records, rec)
	}
	return records, nil
}

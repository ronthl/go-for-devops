package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// record represents a record containing first name and last name that are store in a csv.
type record []string

// validate validates if the csv line was had the correct number of entries.
func (rec record) validate() error {
	if len(rec) != 2 {
		return errors.New("data format is incorrect")
	}
	return nil
}

// first returns the record's first name.
func (rec record) first() string {
	return rec[0]
}

// last returns the record's last name.
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

// readRecords stream a csv file and manipulates it to return the records.
// This will skip any blank lines and stop on the first error encountered.
func readRecords() ([]record, error) {
	file, err := os.Open("data.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var records []record

	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		var rec record = strings.Split(line, ",")
		if err = rec.validate(); err != nil {
			return nil, fmt.Errorf("entry at line %d was invalid: %w", lineNum, err)
		}

		records = append(records, rec)
	}

	return records, nil
}

package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

const FLAG = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
const CSV_SORTED = "data-sorted.csv"

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

// csv outputs the data in csv format.
func (rec record) csv() []byte {
	b := bytes.Buffer{}
	for _, field := range rec {
		b.WriteString(field + ",")
	}

	b.WriteString("\n")
	return b.Bytes()
}

func main() {
	records, err := readRecords("data.csv")
	if err != nil {
		panic(err)
	}

	if err = writeRecords(records); err != nil {
		panic(err)
	}

	data, err := os.ReadFile(CSV_SORTED)
	if err != nil {
		panic(err)
	}
	fmt.Printf("*** SORTED ***\n%s", data)
}

// readRecords stream a csv file and manipulates it to return the records.
// This will skip any blank lines and stop on the first error encountered.
func readRecords(filename string) ([]record, error) {
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

func writeRecords(records []record) error {
	file, err := os.OpenFile(CSV_SORTED, FLAG, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Sort by last name
	sort.Slice(
		records,
		func(i, j int) bool {
			return records[i].last() < records[j].last()
		},
	)

	for _, rec := range records {
		_, err := file.Write(rec.csv())
		if err != nil {
			return err
		}
	}

	return nil
}

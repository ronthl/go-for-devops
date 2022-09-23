package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
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
	records := []record{
		{"Johnny", "Depp"},
		{"Emma", "Watson"},
		{"Sasuke", "Uchiha"},
	}

	if err := writeRecords(records); err != nil {
		panic(err)
	}

	data, err := os.ReadFile(CSV_SORTED)
	if err != nil {
		panic(err)
	}
	fmt.Printf("*** SORTED ***\n%s", data)
}

// writeRecords writes the given records to a csv file.
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

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, rec := range records {
		if err = writer.Write(rec); err != nil {
			return err
		}
	}

	return nil
}

package csvutils

import (
	"encoding/csv"
	"os"
)

type CSVSimpleHandler interface {
	NewCSV() *csvObj
	Init([][]string) *csvObj
	Load(string) *csvObj
	Dump(string) bool
	GetAsSlice() [][]string
}

type csvObj struct {
	csvData [][]string
}

// Returns pointer to new empty CSV object.
func NewCSV() *csvObj {
	return &csvObj{}
}

// Populates object with values from data.
// Returns "this".
func (c *csvObj) Init(data [][]string) *csvObj {
	(*c).csvData = data

	return c
}

// Simple dump data to a file. If file exists it will be overwritten.
func (c *csvObj) Dump(fileName string) bool {

	file, err := os.OpenFile(fileName, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, os.FileMode(0644))

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if err = csv.NewWriter(file).WriteAll(c.csvData); err != nil {
		panic(err)
	}

	return true
}


// Load data from CSV file to CSV object and returns "this" (CSV object itself).
// Empty strings will be ignored.
// Use .GetAsSlice method to get CSV data as 2D slice.
func (c *csvObj) Load(fileName string) *csvObj {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		panic(err)
	}

	(*c).csvData = lines

	return c
}

// Returns CSV data in useful format as simple 2D slice.
func (c *csvObj) GetAsSlice() [][]string {
	return (*c).csvData
}

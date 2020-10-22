package csvutils

import (
	"encoding/csv"
	"os"
)

type CSVReaderWriter interface {
	NewCSV() *CSV
	Init([][]string) *CSV
	Load(string) *CSV
	Dump(string) bool
	GetAsSlice() [][]string
}

type CSV struct {
	csvData [][]string
}

// Returns pointer to new empty CSV object.
func NewCSV() *CSV {
	return &CSV{}
}

// Populates object with values from data.
// Returns "this".
func (c *CSV) Init(data [][]string) *CSV {
	(*c).csvData = data

	return c
}

// Simple dump data to a file. If file exists it will be overwritten.
func (c *CSV) Dump(fileName string) bool {

	file, err := os.OpenFile(fileName, os.O_WRONLY, os.FileMode(0644))

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
func (c *CSV) Load(fileName string) *CSV {
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
func (c *CSV) GetAsSlice() [][]string {
	return (*c).csvData
}

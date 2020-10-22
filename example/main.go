package main

import (
	"fmt"
	"github.com/anxp/csvutils"
)

func main() {
	// Creates empty CSV object:
	csvObject := csvutils.NewCSV()

	// Let's add some normal data and an empty string to CSV file, to see how it will be handled:
	csvObject.Init([][]string{{"111","222","333","444"}, {"AAA","BBB","CCC","DDD"}, {""}, {"333","888","000","222"}, {"44.1","22.8","999.99","11.2"}})

	// Write file to disc:
	csvObject.Dump("test.txt")

	// Load file from disc and returns it in convenient format (2D slice):
	fileContentAsSlice := csvObject.Load("test.txt").GetAsSlice()

	fmt.Print(fileContentAsSlice)
	return
}
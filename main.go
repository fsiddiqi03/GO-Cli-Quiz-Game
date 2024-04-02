package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of question, answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		// sends message as a string
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
		os.Exit(1)
	}

	r := csv.NewReader(file)
	line, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provide csv file.")
	}
	fmt.Print(line)
}

func parseLines(lines [][]string) []problem {
	// create a slice with the problem struct, set len to len of slices
	ret := make([]problem, len(lines))
	for i, line := range lines { // loop through the 2D slices and creates a problem, while adding them to res
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

// create a problem object
type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Print(msg)
	os.Exit(1)
}

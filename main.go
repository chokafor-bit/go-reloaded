package main

import (
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		return
	}

	text := string(data)

	// Apply all transformations
	text = ProcessText(text)

	// Write result to output file
	os.WriteFile(outputFile, []byte(text), 0644)
}

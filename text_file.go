package main

import "testing"

func TestProcessText(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "simple cap",
			input:  "It (cap) was the best",
			output: "It Was the best",
		},
		{
			name:   "multiple cap",
			input:  "it was the age of foolishness (cap, 6)",
			output: "It Was The Age Of Foolishness",
		},
		{
			name:   "uppercase",
			input:  "hello world (up)",
			output: "hello WORLD",
		},
		{
			name:   "lowercase",
			input:  "HELLO WORLD (low)",
			output: "HELLO world",
		},
		{
			name:   "multiple words up",
			input:  "one two three (up, 2)",
			output: "one TWO THREE",
		},
		{
			name:   "no command",
			input:  "hello world",
			output: "hello world",
		},
		{
			name:   "empty input",
			input:  "",
			output: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ProcessText(test.input)
			if result != test.output {
				t.Errorf("Input: %q\nExpected: %q\nGot: %q",
					test.input, test.output, result)
			}
		})
	}
}

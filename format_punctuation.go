package main

import (
	"strings"
)

// FormatPunctuation fixes punctuation spacing
func FormatPunctuation(text string) string {
	// Handle punctuation marks: ., ,, !, ?, :, ;
	punctuationMarks := []string{".", ",", "!", "?", ":", ";"}

	// Process each punctuation mark
	result := text
	for _, mark := range punctuationMarks {
		// Remove spaces before punctuation
		result = strings.ReplaceAll(result, " "+mark, mark)
	}

	return result
}

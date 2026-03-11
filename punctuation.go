package main

import "strings"

func FormatPunctuation(text string) string {
	// text = strings.ReplaceAll(orig, old, new)
	text = strings.ReplaceAll(text, " ... ", "... ")

	text = strings.ReplaceAll(text, " .", ". ")

	text = strings.ReplaceAll(text, " ,", ", ")

	text = strings.ReplaceAll(text, " !", "! ")

	text = strings.ReplaceAll(text, " ?", "?")

	text = strings.ReplaceAll(text, " :", ": ")

	text = strings.ReplaceAll(text, " ;", ";")

	text = strings.ReplaceAll(text, " ' ", "'")

	return text
}

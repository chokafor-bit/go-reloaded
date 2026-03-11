package main

import "strings"

func FormatQuotes(text string) string {
	parts := strings.Split(text, "'")
	result := ""
	for i := 0; i < len(parts); i++ {
		if i%2 == 1 {
			result += "'" + strings.TrimSpace(parts[i]) + "'"
		} else {
			result += parts[i]
		}

	}
	return result
}

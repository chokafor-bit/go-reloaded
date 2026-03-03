package main

import (
	"regexp"
	"strings"
)

// FormatQuotes handles single quotes
func FormatQuotes(text string) string {
	re := regexp.MustCompile(`'([^']*)'`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		content := strings.TrimSpace(match[1 : len(match)-1])
		return "'" + content + "'"
	})
}

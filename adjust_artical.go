package main

import (
	"strings"
)

// AdjustArticle changes "a" to "an" before vowels or h
func AdjustArticle(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {

		if strings.ToLower(words[i]) == "a" {

			if len(words[i+1]) > 0 {

				first := strings.ToLower(string(words[i+1][0]))

				if first == "a" || first == "e" || first == "i" ||
					first == "o" || first == "u" || first == "h" {

					if words[i] == "A" {
						words[i] = "An"
					} else {
						words[i] = "an"
					}
				}
			}
		}
	}

	return strings.Join(words, " ")
}

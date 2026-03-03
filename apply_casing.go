package main

import (
	"strconv"
	"strings"
)

func ApplyCasing(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		// Simple commands
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = remove(words, i)
			i--
		}

		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = remove(words, i)
			i--
		}

		if words[i] == "(cap)" && i > 0 {
			words[i-1] = capitalize(words[i-1])
			words = remove(words, i)
			i--
		}

		// Commands with number: (up, 2)
		if (words[i] == "(up," || words[i] == "(low," || words[i] == "(cap,") && i+1 < len(words) {

			numStr := strings.TrimRight(words[i+1], ")")
			num, err := strconv.Atoi(numStr)

			if err == nil {

				for j := 1; j <= num && i-j >= 0; j++ {

					if words[i] == "(up," {
						words[i-j] = strings.ToUpper(words[i-j])
					}

					if words[i] == "(low," {
						words[i-j] = strings.ToLower(words[i-j])
					}

					if words[i] == "(cap," {
						words[i-j] = capitalize(words[i-j])
					}
				}
			}

			// remove "(up," and number part
			words = remove(words, i)
			words = remove(words, i)
			i--
		}
	}

	return strings.Join(words, " ")
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}

	word = strings.ToLower(word)
	return strings.ToUpper(string(word[0])) + word[1:]
}

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

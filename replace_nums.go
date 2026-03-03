package main

import (
	"strconv"
	"strings"
)

// ReplaceNums converts (hex) and (bin) markers
func ReplaceNums(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		if i > 0 && words[i] == "(hex)" {

			hexNum := words[i-1]

			if num, err := strconv.ParseInt(hexNum, 16, 64); err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}

			words = append(words[:i], words[i+1:]...)
			i--
		}

		if i > 0 && words[i] == "(bin)" {

			binNum := words[i-1]

			if num, err := strconv.ParseInt(binNum, 2, 64); err == nil {
				words[i-1] = strconv.FormatInt(num, 10)
			}

			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return strings.Join(words, " ")
}

package main

import (
	"strconv"
	"strings"
)

func ReplaceNums(text string) string {
	words := strings.Split(text, " ")

	for i := 0; i < len(words)-1; i++ {
		base := 0
		if words[i+1] == "(hex)" {
			base = 16
		}
		if words[i+1] == "(bin)" {
			base = 2
		}
		if base != 0 {
			num, err := strconv.ParseInt(words[i], base, 6)
			if err == nil {
				words[i] = strconv.FormatInt(num, 10)
			}
			words[i+1] = ""
		}
	}
	return strings.Join(words, " ")
}

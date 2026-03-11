package main

import "strings"

func FixAtoAn(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		first := strings.ToLower(words[i])
		next := strings.ToLower(string(words[i+1][0]))
		if strings.ContainsAny(next, "aeiouh") {
			if first == "a" {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}

		} else {
			if first == "an" {
				if words[i] == "An" {
					words[i] = "A"
				} else {
					words[i] = "a"
				}
			}
		}
	}
	return strings.Join(words, " ")
}

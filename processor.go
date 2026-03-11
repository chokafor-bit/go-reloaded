package main

func ProcessText(text string) string {

	text = ReplaceNums(text)

	text = ApplyCasing(text)

	text = FormatPunctuation(text)

	text = FormatQuotes(text)

	text = FixAtoAn(text)

	
	return text
}

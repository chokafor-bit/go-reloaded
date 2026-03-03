package main

func ProcessText(text string) string {

	text = ReplaceNums(text)
	text = ApplyCasing(text)
	text = FormatPunctuation(text)
	text = FormatQuotes(text)
	text = AdjustArticle(text)

	return text
}

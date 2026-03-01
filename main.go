package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("sample.text input.text")
		return
	}
	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	// Attempt to open input file for reading
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		// If there's an error, display it and return
		fmt.Println("Error opening input file:")
		return
	}
	// Schedule the input file to be closed when the function exits
	defer inputFile.Close()

	// Create a new Scanner to read from inputFile
	scanner := bufio.NewScanner(inputFile)
	// Use strings.Builder for efficient string concatenation
	var builder strings.Builder
	// Loop over all scan events
	for scanner.Scan() {
		// Write each line of text and a newline character
		builder.WriteString(scanner.Text())
		builder.WriteByte('\n')
	}
	text := builder.String()

	// Check for any errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		// Print any scanning errors
		fmt.Println("Error reading input file")
		// Exit with an error status
		return
	}

	// Process the text to apply transformations
	processedText := ProcessText(text)

	// Attempt to create the output file
	outputFile, err := os.Create(outputFileName)
	// Check for errors while creating the file
	if err != nil {
		// Print an error message and exit if a file create error occurs
		fmt.Println("Error creating output file:")
		return
	}
	// Schedule the output file to be closed when the function exits
	defer outputFile.Close()

	// Create a buffered writer for the output file
	writer := bufio.NewWriter(outputFile)
	// Write the processed text to the output file buffer
	_, err = writer.WriteString(processedText)
	// Check for an error while writing
	if err != nil {
		// Print write error and exit if one occurs
		fmt.Println("Error writing to output file:")
		return
	}
	// Flush the buffer to ensure all content is written to the file
	writer.Flush()
}

// ProcessText performs all the text transformations.
func ProcessText(text string) string {

	text = ReplaceNums(text)

	text = ApplyCasing(text)
	text = FormatPunctuation(text)

	text = FormatQuotes(text)
	text = AdjustArticle(text)

	return text
}

// ReplaceNums converts (hex) and (bin) markers
func ReplaceNums(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if i > 0 {
			if words[i] == "(hex)" {
				// Convert the previous word from hex to decimal
				hexNum := words[i-1]
				if num, err := strconv.ParseInt(hexNum, 16, 64); err == nil {
					words[i-1] = strconv.FormatInt(num, 10)
				}
				// Remove the (hex) marker
				words = append(words[:i], words[i+1:]...)
				i--
			} else if words[i] == "(bin)" {
				// Convert the previous word from binary to decimal
				binNum := words[i-1]
				if num, err := strconv.ParseInt(binNum, 2, 68); err == nil {
					words[i-1] = strconv.FormatInt(num, 10)
				}
				// Remove the (bin) marker and add to it with append
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(words, " ")
}

// ApplyCasing handles (up), (low), (cap) commands
func ApplyCasing(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {
			cmd := strings.Trim(words[i], "()")

			// Handle commands with numbers: (up, 2)
			if strings.Contains(cmd, ",") {
				parts := strings.Split(cmd, ",")
				command := strings.TrimSpace(parts[0])
				numStr := strings.TrimSpace(parts[1])
				num, err := strconv.Atoi(numStr)
				if err != nil {
					i++
					continue
				}

				// Apply command to previous 'num' words
				switch command {
				case "up":
					for j := 1; j <= num && i-j >= 0; j++ {
						words[i-j] = strings.ToUpper(words[i-j])
					}
				case "low":
					for j := 1; j <= num && i-j >= 0; j++ {
						words[i-j] = strings.ToLower(words[i-j])
					}
				case "cap":
					for j := 1; j <= num && i-j >= 0; j++ {
						words[i-j] = strings.Title(strings.ToLower(words[i-j]))
					}
				}
				// Remove the command
				words = append(words[:i], words[i+1:]...)
				i--
			} else {
				// Handle simple commands: (up), (low), (cap)
				switch cmd {
				case "up":
					if i >= 0 {
						words[i-1] = strings.ToUpper(words[i-1])
					}
					words = append(words[:i], words[i+1:]...)
					i--
				case "low":
					if i >= 0 {
						words[i-1] = strings.ToLower(words[i-1])
					}
					words = append(words[:i], words[i+1:]...)
					i--
				case "cap":
					if i >= 0 {
						words[i-1] = strings.Title(strings.ToLower(words[i-1]))
					}
					words = append(words[:i], words[i+1:]...)
					i--
				}
			}
		}
	}
	return strings.Join(words, " ")
}

// FormatPunctuation fixes punctuation spacing
func FormatPunctuation(text string) string {
	// Handle punctuation marks: ., ,, !, ?, :, ;
	punctuationMarks := []string{".", ",", "!", "?", ":", ";"}

	for _, mark := range punctuationMarks {
		// Remove spaces before punctuation
		reBefore := regexp.MustCompile(`\s+` + regexp.QuoteMeta(mark))
		text = reBefore.ReplaceAllString(text, mark)

		// Add space after punctuation if needed
		reAfter := regexp.MustCompile(regexp.QuoteMeta(mark) + `([^\s])`)
		text = reAfter.ReplaceAllString(text, mark+" $1")
	}

	// Handle groups of punctuation like ... or !?
	reEllipsis := regexp.MustCompile(`\.\s*\.\s*\.`)
	text = reEllipsis.ReplaceAllString(text, "...")

	reExclQuestion := regexp.MustCompile(`\!\s*\?`)
	text = reExclQuestion.ReplaceAllString(text, "!?")

	return text
}

// FormatQuotes handles single quotes
func FormatQuotes(text string) string {
	// Find quoted sections and remove spaces inside quotes
	re := regexp.MustCompile(`[.,!?:;]`)
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		// Extract content between quotes
		content := match[1 : len(match)-1]
		content = strings.TrimSpace(content)
		return "'" + content + "'"
	})
	return text
}

// AdjustArticle handles a/an conversion
func AdjustArticle(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		// Check if current word is "a" or "A"
		if strings.ToLower(words[i]) == "a" {
			// Check if next word starts with a vowel or 'h'
			nextWord := words[i+1]
			if len(nextWord) > 0 {
				firstChar := strings.ToLower(string(nextWord[0]))
				if firstChar == "a" || firstChar == "e" || firstChar == "i" || firstChar == "o" || firstChar == "u" || firstChar == "h" {
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

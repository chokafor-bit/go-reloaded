Text Modification Tool
Description

This Go program is designed to modify text files based on specific patterns and rules defined within the code. It performs various modifications such as converting hexadecimal and binary strings to decimal, changing letter cases, handling punctuation, correcting vowel usage, etc.
Features

    Modify an input file from the command line.
    Custom commands to modify text files
    Conversion of hexadecimal and binary strings to decimal.
    Uppercase, lowercase, and capitalize words according to predefined patterns.
    Handling of punctuation marks and spaces.
    Correction of vowel usage in specific contexts.

Functions

    modifyFile(inputFilename string, outputFilename string) error: Reads an input text file, applies modifications, and writes the modified text to an output file.
    hexToDec(hexString string) (string, error): Converts a hexadecimal string to a decimal string.
    binToDec(binString string) (string, error): Converts a binary string to a decimal string.
    processWordModifications(line string) string: Processes word modifications such as uppercase, lowercase, capitalize, punctuation, etc.

Custom Commands for Modification
Hexadecimal Conversion

    (hex): Converts the preceding hexadecimal string to its decimal equivalent.

Binary Conversion

    (bin): Converts the preceding binary string to its decimal equivalent.

Uppercase

    (up): Converts the preceding word to uppercase.
    (up, n): Converts the preceding n words to uppercase.

Capitalize

    (cap): Capitalizes the first letter of the preceding word.
    (cap, n): Capitalizes the first letter of the preceding n words.

Lowercase

    (low): Converts the preceding word to lowercase.
    (low, n): Converts the preceding n words to lowercase.

Punctuation Modification

    ,: Removes spaces before commas.
    .: Removes spaces before periods.
    ...: Converts multiple consecutive periods to a single period.
    ;: Removes spaces before semicolons.
    :: Removes spaces before colons.
    !: Removes spaces before exclamation marks.
    ?: Removes spaces before question marks.

Single Quotation Mark Modification

    ': Removes spaces before single quotation marks.

Vowel Correction

    a followed by a word starting with a vowel: Changes "a" to "an".
    A followed by a word starting with a vowel: Changes "A" to "An".

Usage

To use this tool, follow these steps:

    Clone Repository: Clone this repository to your local machine.
    Navigate to Directory: Open a terminal and navigate to the directory where the code is located.
    Run the Program: Execute the program by providing input and output file paths as command-line arguments. For example:

    go run . sample.txt result.txt

    Replace sample.txt with the path to your input text file and result.txt with the desired output file path.
    Check Output: Once the program finishes execution, check the specified output file for the modified text.

Input File Format

The input file should contain the text to be modified. Each line of the file will be processed independently.
Output File Format

The output file will contain the modified text based on the rules defined in the program. Each line of the input file will correspond to a line in the output file with the modifications applied.
Commands

    go run . sample.txt result.txt: Runs the program with the specified input and output file paths.

Notes

    Ensure that the input file exists and is accessible.
    The program may terminate with an error if the input file cannot be read or if there are issues during processing.

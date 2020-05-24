package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var words string

	for scanner.Scan() {
		word := scanner.Text()
		words = words + word + "\n"
	}

	// Replace words that only consist of digits
	regex := regexp.MustCompile(`(?m)^[\d]*$`)
	words = regex.ReplaceAllString(words, "")

	// Replace words that contain of at least 2 digits
	regex = regexp.MustCompile(`(?m)^.*[\d]{2}.*$`)
	words = regex.ReplaceAllString(words, "")

	// Replace words that only contain at least 3 consecutive vowels
	regex = regexp.MustCompile(`(?m)^[euioa]{3,}$`)
	words = regex.ReplaceAllString(words, "")

	// Replace words that only contain non vowels and at least 4 characters long
	regex = regexp.MustCompile(`(?m)^[^euioa]{4,}$`)
	words = regex.ReplaceAllString(words, "")

	// Replace words that are very long
	regex = regexp.MustCompile(`(?m)^.{15,}$`)
	words = regex.ReplaceAllString(words, "")

	// Replace words that contain non-ascii chars
	regex = regexp.MustCompile(`(?m)^.*[^\x00-\x7F]+.*$`)
	words = regex.ReplaceAllString(words, "")

	// Replace words that start with a digit
	regex = regexp.MustCompile(`(?m)^[\d]+.*$`)
	words = regex.ReplaceAllString(words, "")

	// Replace words that end with a digit
	regex = regexp.MustCompile(`(?m)^.*[\d]+$`)
	words = regex.ReplaceAllString(words, "")

	// Remove all empty lines
	regex = regexp.MustCompile(`(?m)^\n`)
	words = regex.ReplaceAllString(words, "")

	fmt.Println(words)
}

package main

import (
	"strconv"
	"strings"
)

func processText(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		if words[i] == "(hex)" && i > 0 {
			value, err := strconv.ParseInt(words[i-1], 16, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(value, 10)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}

		if words[i] == "(bin)" && i > 0 {
			value, err := strconv.ParseInt(words[i-1], 2, 64)
			if err == nil {
				words[i-1] = strconv.FormatInt(value, 10)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}

		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}

		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}

		if words[i] == "(cap)" && i > 0 {
			words[i-1] = capitalize(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}

		if i < len(words)-1 &&
			(strings.HasPrefix(words[i], "(up,") ||
				strings.HasPrefix(words[i], "(low,") ||
				strings.HasPrefix(words[i], "(cap,")) {

			command := words[i]
			numberToken := words[i+1]

			numberToken = strings.TrimSuffix(numberToken, ")")
			n, err := strconv.Atoi(numberToken)
			if err == nil && i > 0 {

				for j := 0; j < n && i-1-j >= 0; j++ {
					index := i - 1 - j

					if strings.HasPrefix(command, "(up,") {
						words[index] = strings.ToUpper(words[index])
					}

					if strings.HasPrefix(command, "(low,") {
						words[index] = strings.ToLower(words[index])
					}

					if strings.HasPrefix(command, "(cap,") {
						words[index] = capitalize(words[index])
					}
				}

				words = append(words[:i], words[i+2:]...)
				i--
			}
		}
	}

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" && startsWithVowelOrH(words[i+1]) {
			words[i] = "an"
		}
		if words[i] == "A" && startsWithVowelOrH(words[i+1]) {
			words[i] = "An"
		}
	}
	result := strings.Join(words, " ")
	result = formatPunctuation(result)
	result = formatQuotes(result)
	return result
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func startsWithVowelOrH(word string) bool {
	if len(word) == 0 {
		return false
	}

	first := strings.ToLower(string(word[0]))

	return first == "a" ||
		first == "e" ||
		first == "i" ||
		first == "o" ||
		first == "u" ||
		first == "h"
}

func formatPunctuation(text string) string {
	var result []rune
	runes := []rune(text)

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if isPunctuation(char) {
			if len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}

			result = append(result, char)

			if i+1 < len(runes) && runes[i+1] == char {
				continue
			}

			if i+1 < len(runes) && runes[i+1] != ' ' {
				result = append(result, ' ')
			}

		} else {
			result = append(result, char)
		}
	}
	return strings.TrimSpace(string(result))
}

func isPunctuation(r rune) bool {
	return r == '.' ||
		r == ',' ||
		r == '!' ||
		r == '?' ||
		r == ':' ||
		r == ';'
}

func formatQuotes(text string) string {
	runes := []rune(text)
	var result []rune

	insideQuote := false

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if char == '\'' {

			if insideQuote {
				if len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
			}
			result = append(result, char)
			insideQuote = !insideQuote

			if insideQuote && i+1 < len(runes) && runes[i+1] == ' ' {
				i++
			}

		} else {
			result = append(result, char)
		}
	}

	return string(result)
}

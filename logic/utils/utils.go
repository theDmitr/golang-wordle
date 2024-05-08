package utils

import (
	"os"
	"strings"
)

func LoadWordsFromFile(filePath string) *[]string {
	var words []string
	file, err := os.ReadFile(filePath)

	if err != nil {
		return &words
	}

	data := string(file)
	words = append(words, strings.Split(data, "\n")...)

	for i := range words {
		words[i] = strings.TrimRight(words[i], string(rune(13)))
	}

	return &words
}

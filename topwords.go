package topwords

import (
	"regexp"
	"sort"
	"strings"
)

// Sort map of words and return slice of words
func sortMap(list map[string]int) []string {
	wordList := make([]string, 0, len(list))

	for word := range list {
		wordList = append(wordList, word)
	}

	sort.Slice(wordList, func(i, j int) bool {
		return list[wordList[i]] > list[wordList[j]]
	})

	return wordList
}

// Return list of most popular words limited by max count
func GetTopWords(text string, maxCount int) []string {
	words := make(map[string]int)
	for _, word := range chunk(text) {
		word = strings.ToLower(word)
		words[word]++
	}

	result := sortMap(words)
	if totalWords := len(result); totalWords < maxCount || maxCount < 0 {
		maxCount = totalWords
	}

	return result[:maxCount]
}

// Cleans up text of special characters and multiple space then return slice of words
func chunk(text string) []string {
	replacer := regexp.MustCompile(`[[:space:][:punct:]]+`)
	text = replacer.ReplaceAllString(text, " ")
	result := strings.Split(strings.TrimSpace(text), " ")

	return result
}

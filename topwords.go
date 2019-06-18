package topwords

import (
	"regexp"
	"strings"
)

// Sort map of words and return slice of words
func sortMap(list map[string]int, revert bool) []string {
	idxList := make([]int, 0, len(list))
	wordList := make([]string, 0, len(list))

	for word, count := range list {
		idxList = append(idxList, count)
		wordList = append(wordList, word)
	}

	for i := 0; i < len(wordList); i++ {
		for j := i; j > 0; j-- {
			if (!revert && idxList[j-1] > idxList[j]) || (revert && idxList[j-1] < idxList[j]) {
				idxList[j-1], idxList[j] = idxList[j], idxList[j-1]
				wordList[j-1], wordList[j] = wordList[j], wordList[j-1]
			}
		}
	}

	return wordList
}

// Return list of most popular words limited by max count
func GetTopWords(text string, maxCount int) []string {
	words := make(map[string]int)
	for _, word := range chunk(text) {
		word = strings.ToLower(word)
		words[word]++
	}

	result := sortMap(words, true)
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

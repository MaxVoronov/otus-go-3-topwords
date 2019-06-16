package topwords

import (
	"regexp"
	"sort"
	"strings"
)

type Word struct {
	Name  string
	Count uint
}

type WordList []Word

// Return length of word list.
// Used for implementation sort.Interface
func (words WordList) Len() int {
	return len(words)
}

// Compare two words by number of occurrences in the text.
// Used for implementation sort.Interface
func (words WordList) Less(i, j int) bool {
	return words[i].Count < words[j].Count
}

// Swaps the elements with indexes.
// Used for implementation sort.Interface
func (words WordList) Swap(i, j int) {
	words[i], words[j] = words[j], words[i]
}

// Try to find word in list.
// Return index of word position in slice and does the word exist
func (words WordList) findWord(word string) (int, bool) {
	for idx, value := range words {
		if value.Name == word {
			return idx, true
		}
	}

	return -1, false
}

// Return list of most popular words limited by max count
func GetTopWords(words []string, maxCount uint) WordList {
	result := make(WordList, 0, len(words))

	for _, word := range words {
		word = strings.ToLower(word)

		if idx, ok := result.findWord(word); ok {
			result[idx].Count++
		} else {
			result = append(result, Word{word, 1})
		}
	}

	sort.Sort(sort.Reverse(result))

	return result[:maxCount]
}

// Cleans up text of special characters and multiple space then return slice of words
func chunk(text string) []string {
	replacer := regexp.MustCompile(`[[:space:][:punct:]]+`)
	text = replacer.ReplaceAllString(text, " ")
	result := strings.Split(strings.TrimSpace(text), " ")

	return result
}

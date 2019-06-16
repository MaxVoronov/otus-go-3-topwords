package topwords

import "testing"

const ShortTextSample = "Test text! Simple text, just for test? More Text"
const TextSample = `
	This is the house that Jack built.
    This is the malt that lay in the house that Jack built.
    This is the rat that ate the malt
    That lay in the house that Jack built.
    This is the cat
    That killed the rat that ate the malt
    That lay in the house that Jack built.
    This is the dog that worried the cat
    That killed the rat that ate the malt
    That lay in the house that Jack built.`

func TestTextChunking(t *testing.T) {
	expectedChunksCount := 9
	parsedChunks := chunk(ShortTextSample)
	if chunks := len(parsedChunks); chunks != 9 {
		t.Errorf("Not corrects chunks count: expect %d, got %d", expectedChunksCount, chunks)
	}
}

func TestFindWord(t *testing.T) {
	searchWord := "text"
	words := WordList{
		Word{"test", 2},
		Word{searchWord, 3},
		Word{"simple", 1},
	}

	if idx, ok := words.findWord(searchWord); !ok || idx != 1 {
		t.Errorf("Test word \"%s\" not found", searchWord)
	}
}

func TestGetTopWords(t *testing.T) {
	var maxCount uint = 10
	chunks := chunk(TextSample)
	topWords := GetTopWords(chunks, maxCount)

	if wordsCount := len(topWords); uint(wordsCount) != maxCount {
		t.Errorf("Not correct total words count: expect %d, got %d", maxCount, wordsCount)
	}

	searchWord := "the"
	if idx, ok := topWords.findWord(searchWord); !ok || idx != 0 {
		t.Errorf("Test word \"%s\" not found", searchWord)
	}
}

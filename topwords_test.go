package topwords

import "testing"

const ShortTextSample = "Test text! Simple text, just for test? More Text"
const TextSample = `The
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

func TestGetTopWords(t *testing.T) {
	var maxCount = 10
	topWords := GetTopWords(TextSample, maxCount)

	if wordsCount := len(topWords); wordsCount != maxCount {
		t.Errorf("Not correct total words count: expect %d, got %d", maxCount, wordsCount)
	}

	searchWord := "the"
	if resultWord := topWords[0]; searchWord != resultWord {
		t.Errorf("Not correct result word: expect \"%s\", got \"%s\"", searchWord, resultWord)
	}
}

// Read and parse given text
package reader

import (
	"regexp"
	"strings"
	"sync"
	"unicode/utf8"
)

// textStructure internal structure of reader
type textStructure struct {
	itemsMutex    *sync.Mutex
	mapper        mapper
	minWordLength int
	rgp           *regexp.Regexp
	excludeWords  map[string]bool
}

// mapper internal interface of mapperClass
type mapper interface {
	Insert([]string, uint8)
	Remove(string)
}

// New reader Constructor
func New(mapper mapper, minWordLength int) *textStructure {
	return &textStructure{
		mapper:        mapper,
		minWordLength: minWordLength,
		rgp:           regexp.MustCompile(`[^a-zA-Z\s.0-9]+`),
		excludeWords:  make(map[string]bool),
		itemsMutex:    new(sync.Mutex),
	}
}

// Read read and parse each line from the text
func (t *textStructure) Read(content string, paragraphNumber uint8) {
	line := t.rgp.ReplaceAllString(content, "")

	if line = strings.TrimSpace(line); line != "" {
		for _, line := range strings.Split(line, ".") {
			t.parseLine(strings.ToLower(strings.TrimSpace(line)), paragraphNumber)
		}
	}

}

// parseLine split a line to word
func (t textStructure) parseLine(line string, paragraphNumber uint8) {
	words := strings.Split(line, " ")
	var resultWords []string

	wordsCount := len(words)
	for index, word := range words {
		if index == 0 || index == wordsCount-1 {
			t.updateExcludeList(word)
			t.mapper.Remove(word)
			continue
		}

		if t.isWordValid(word) {
			resultWords = append(resultWords, word)
		}
	}

	t.mapper.Insert(resultWords, paragraphNumber)
}

// updateExcludeList Update exclude list with words
func (t *textStructure) updateExcludeList(word string) {
	t.itemsMutex.Lock()
	defer t.itemsMutex.Unlock()

	if _, ok := t.excludeWords[word]; !ok {
		t.excludeWords[word] = true
	}
}

// isWordValid returns true if word is valid
func (t *textStructure) isWordValid(word string) bool {
	return utf8.RuneCountInString(word) > t.minWordLength && !t.isWordExcluded(word)
}

// isWordExcluded returns if given word in exclude list
func (t *textStructure) isWordExcluded(word string) bool {
	t.itemsMutex.Lock()
	defer t.itemsMutex.Unlock()

	_, ok := t.excludeWords[word]

	return ok
}

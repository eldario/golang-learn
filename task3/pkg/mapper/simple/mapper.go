// package for collect words
package mapper

import (
	"sort"
	"sync"
)

// sortedMap internal structure for class
type sortedMap struct {
	itemsMutex sync.Mutex
	items      []string
	words      map[string]int
	topCount   int
}

// wordItem structure of word object
type wordItem struct {
	Word  string
	Count int
	order int
}

// New Structure constructor
func New(topCount int) *sortedMap {
	return &sortedMap{words: make(map[string]int), topCount: topCount}
}

// Insert a new word in words map
func (s *sortedMap) Insert(word string) {
	s.itemsMutex.Lock()
	if _, ok := s.words[word]; !ok {
		s.items = append(s.items, word)
		s.words[word] = 0
	}

	s.words[word]++
	s.itemsMutex.Unlock()
}

// Remove a word from words map
func (s *sortedMap) Remove(word string) {
	if _, ok := s.words[word]; ok {
		s.itemsMutex.Lock()
		delete(s.words, word)
		s.itemsMutex.Unlock()
	}
}

// GetResults Get frequently used words in text
func (s *sortedMap) GetResults() []wordItem {
	var sortedResult []wordItem

	for index, word := range s.items {
		count := s.words[word]
		sortedResult = append(sortedResult, wordItem{word, count, index})
	}

	sort.Slice(sortedResult, func(i, j int) bool {
		if sortedResult[i].Count == sortedResult[j].Count {
			return sortedResult[i].order < sortedResult[j].order
		}
		return sortedResult[i].Count > sortedResult[j].Count
	})

	if len(sortedResult) >= s.topCount {
		sortedResult = sortedResult[:s.topCount]
	}

	sort.Slice(sortedResult, func(i, j int) bool {
		return sortedResult[i].order < sortedResult[j].order
	})

	return sortedResult
}

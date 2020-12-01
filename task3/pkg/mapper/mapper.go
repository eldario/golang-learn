// package for collect words
package mapper

import (
	"sort"
	"sync"
)

// sortedMap internal structure for class
type sortedMap struct {
	itemsMutex *sync.Mutex
	items      map[string]uint8
	words      map[string]int
	topCount   int
}

// wordItem structure of word object
type wordItem struct {
	Word  string
	Count int
	Order uint8
}

// New Structure constructor
func New(topCount int) *sortedMap {
	return &sortedMap{
		items:      make(map[string]uint8),
		words:      make(map[string]int),
		topCount:   topCount,
		itemsMutex: new(sync.Mutex),
	}
}

// Insert a new word in words map
func (s *sortedMap) Insert(word string, position uint8) {
	s.itemsMutex.Lock()
	defer s.itemsMutex.Unlock()

	if _, ok := s.words[word]; !ok {
		s.words[word] = 0
	}


	if _, ok := s.items[word]; !ok {
		s.items[word] = 0
	}
	s.items[word]++

	if order := s.items[word]; order > position {
		s.items[word] = position
	}

	s.words[word]++
}

// Remove a word from words map
func (s *sortedMap) Remove(word string) {
	s.itemsMutex.Lock()
	defer s.itemsMutex.Unlock()

	if _, ok := s.words[word]; ok {
		delete(s.words, word)
	}
}

// GetResults Get frequently used words in text
func (s *sortedMap) GetResults() []wordItem {
	var sortedResult []wordItem
	for word, index := range s.items {
		count := s.words[word]
		sortedResult = append(sortedResult, wordItem{word, count, index})
	}

	sort.Slice(sortedResult, func(i, j int) bool {
		if sortedResult[i].Count == sortedResult[j].Count {
			return sortedResult[i].Order < sortedResult[j].Order
		}
		return sortedResult[i].Count > sortedResult[j].Count
	})

	if len(sortedResult) >= s.topCount {
		sortedResult = sortedResult[:s.topCount]
	}

	sort.Slice(sortedResult, func(i, j int) bool {
		return sortedResult[i].Order < sortedResult[j].Order
	})

	return sortedResult
}

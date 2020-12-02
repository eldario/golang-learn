// package for collect words
package mapper

import (
	"sort"
	"sync"
)

// sortedMap internal structure for class
type sortedMap struct {
	itemsMutex *sync.Mutex
	items      map[uint8][]string
	words      map[string]wordItem
	topCount   int
}

// wordItem structure of word object
type wordItem struct {
	Word  string
	Count int
	Score float32
}

// New Structure constructor
func New(topCount int) *sortedMap {
	return &sortedMap{
		words:      make(map[string]wordItem),
		topCount:   topCount,
		itemsMutex: new(sync.Mutex),
	}
}

// Insert a new word in words map
func (s *sortedMap) Insert(words []string, position uint8) {
	s.itemsMutex.Lock()
	defer s.itemsMutex.Unlock()

	for index, word := range words {
		score := float32(position) + float32(index+1)/1000
		if _, ok := s.words[word]; !ok {
			s.words[word] = wordItem{word, 0, score}
		}
		w := s.words[word]
		w.Count++
		if w.Score > score {
			w.Score = score
		}

		s.words[word] = w
	}

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

	for _, word := range s.words {
		sortedResult = append(sortedResult, word)
	}

	sort.Slice(sortedResult, func(i, j int) bool {
		if sortedResult[i].Count == sortedResult[j].Count {
			return sortedResult[i].Score < sortedResult[j].Score
		}
		return sortedResult[i].Count > sortedResult[j].Count
	})

	if len(sortedResult) >= s.topCount {
		sortedResult = sortedResult[:s.topCount]
	}

	sort.Slice(sortedResult, func(i, j int) bool {
		return sortedResult[i].Score < sortedResult[j].Score
	})

	return sortedResult
}

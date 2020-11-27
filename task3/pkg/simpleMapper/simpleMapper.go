package simpleMapper

import (
	"fmt"
	"sort"
)

type SortedMap struct {
	items    []string
	words    map[string]int
	topCount int
}

type WordItem struct {
	word  string
	count int
	order int
}

/**
 * Structure constructor.
 */
func New() *SortedMap {
	return &SortedMap{words: make(map[string]int)}
}

/**
 * Setter of top count list.
 */
func (s *SortedMap) SetTopCountElements(count int) {
	s.topCount = count
}

/**
 * Insert a new word in words map.
 */
func (s *SortedMap) Insert(word string) {
	if _, ok := s.words[word]; ok {
		s.words[word]++
		return
	}

	s.items = append(s.items, word)
	s.words[word] = 1
}

/**
 * Remove word from list.
 */
func (s *SortedMap) Remove(word string) {
	if _, ok := s.words[word]; ok {
		delete(s.words, word)
	}
}

/**
 * Get frequently used words in text.
 */
func (s *SortedMap) GetFrequentUses() []string {
	var (
		sortedResult []WordItem
		result       []string
	)

	for word, count := range s.words {
		index := Find(s.items, word)
		sortedResult = append(sortedResult, WordItem{word, count, index})
	}

	sort.Slice(sortedResult, func(i, j int) bool {
		if sortedResult[i].count == sortedResult[j].count {
			return sortedResult[i].order < sortedResult[j].order
		}
		return sortedResult[i].count > sortedResult[j].count
	})

	if len(sortedResult) >= s.topCount {
		sortedResult = sortedResult[0:s.topCount]
	}

	sort.Slice(sortedResult, func(i, j int) bool {
		return sortedResult[i].order < sortedResult[j].order
	})

	for index, word := range sortedResult {
		if index+1 > s.topCount {
			break
		}
		result = append(result, fmt.Sprintf("%s;%d;%d", word.word, word.count, word.order))

	}

	return result
}

func Find(words []string, word string) int {
	for i, n := range words {
		if word == n {
			return i
		}
	}
	return len(words)
}

package main

import (
	"fmt"
	"sort"
)

type SortedMap struct {
	items []string
	words map[int][]string
}

/**
 * Structure constructor.
 */
func New() *SortedMap {
	return &SortedMap{words: make(map[int][]string)}
}

/**
 * Insert a new word in words map.
 */
func (s *SortedMap) Insert(word string) {
	if index := s.findElementIndex(word); index != -1 {
		s.increment(word, index)
		return
	}

	s.items = append(s.items, word)
	s.increment(word, 0)
}

/**
 * Exists given key in map.
 */
func (s *SortedMap) findElementIndex(word string) int {
	for count, words := range s.words {
		for _, value := range words {
			if value == word {
				return count
			}
		}
	}

	return -1
}

/**
 * Clear word from map element.
 */
func (s *SortedMap) clear(word string) {
	for count, words := range s.words {
		for index, value := range words {
			if word == value {
				copy(words[index:], words[index+1:])
				s.words[count] = words[:len(words)-1]
				if len(words) == 0 {
					delete(s.words, count)
				}
				break
			}
		}
	}
}

/**
 * Increment word from map element.
 */
func (s *SortedMap) increment(word string, currentCount int) {
	s.clear(word)
	s.words[currentCount+1] = append(s.words[currentCount+1], word)
}

/**
 * Get frequently used words in text.
 */
func (s *SortedMap) GetFrequentUses() []string {
	var values []string
	topList := s.getTopList()
	for _, word := range s.items {
		for _, topWord := range topList {
			if word == topWord {
				values = append(values, word)
			}
		}
	}

	return values
}

/**
 * Get top list of 10 elements from map.
 */
func (s *SortedMap) getTopList() []string {
	var (
		keys   []int
		values []string
	)

	for key, _ := range s.words {
		keys = append(keys, key)
	}

	sort.Sort(sortDesc(keys))

	for _, k := range keys {
		for _, word := range s.words[k] {
			values = append(values, word)
			if len(values) == 10 {
				return values
			}
		}
	}

	return values
}

/**
 * {@inheritDoc}
 */
type sortDesc []int

/**
 * {@inheritDoc}
 */
func (v sortDesc) Len() int {
	return len(v)
}

/**
 * {@inheritDoc}
 */
func (v sortDesc) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

/**
 * {@inheritDoc}
 */
func (v sortDesc) Less(i, j int) bool {
	return v[i] > v[j]
}

/**
 * Handle method.
 */
func main() {
	someList := New()

	text := "o  a b c d e f  h i  a b c d e f  h i  a a z z z x h g g g o o o"

	for _, word := range []rune(text) {
		if word != ' ' {
			someList.Insert(string(word))
		}
	}

	fmt.Println(someList.GetFrequentUses()) //[o a b c d e f h z g]
}

package reader

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
	"tasks/task3/pkg/simpleMapper"
	"unicode/utf8"
)

type TextStructure struct {
	text          io.Reader
	minWordLength int
	excludeWords  map[string]int
}

/**
 * Constructor.
 */
func New(reader io.Reader) *TextStructure {
	return &TextStructure{text: reader, minWordLength: 3, excludeWords: make(map[string]int)}
}

/**
 * Work with all word from Reader.
 */
func (t *TextStructure) GetWords(count int) []string {
	s := bufio.NewScanner(t.text)
	lines := getPreparedLines(s)

	someList := simpleMapper.New()
	for _, line := range lines {
		words := strings.Split(line, " ")
		wordsCount := len(words)

		for index, word := range words {
			// Remove word if it at the end of sentence
			if index == 0 || index == wordsCount-1 {
				t.UpdateExcludeList(word)
				someList.Remove(word)
				continue
			}

			if !t.isWordValid(word) {
				continue
			}

			someList.Insert(word)
		}

	}

	someList.SetTopCountElements(count)
	return someList.GetFrequentUses()
}

/**
 * Minimal word length setter.
 */
func (t *TextStructure) SetMinWordLength(wordLength int) {
	t.minWordLength = wordLength
}

/**
 * Small validation for given word.
 */
func (t *TextStructure) isWordValid(word string) bool {
	if utf8.RuneCountInString(word) <= t.minWordLength { // if length word less than 3 symbols
		return false
	}

	if _, err := strconv.Atoi(word); err == nil {
		return false
	}

	return true
}

/**
 * Update exclude list with words.
 */
func (t *TextStructure) UpdateExcludeList(word string) {
	if _, ok := t.excludeWords[word]; !ok {
		t.excludeWords[word] = 1
	}
}

/**
 * Preparing sentences from read lines.
 */
func getPreparedLines(s *bufio.Scanner) []string {
	var lines []string
	for s.Scan() {
		line := strings.TrimSpace(regexp.MustCompile("[^a-zA-Z .]+").ReplaceAllString(s.Text(), ""))

		if line == "" {
			continue
		}

		for _, splitLine := range strings.Split(line, ". ") {
			if strings.HasSuffix(splitLine, ".") {
				splitLine = strings.Replace(splitLine, ".", "", 1)
			}

			lines = append(lines, strings.ToLower(splitLine))
		}
	}

	return lines
}

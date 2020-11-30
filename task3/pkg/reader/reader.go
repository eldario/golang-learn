package reader

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type TextStructure struct {
	mapperType
	text          io.Reader
	minWordLength int
}

type mapperType interface {
	IsWordExcluded(word string) bool
	UpdateExcludeList(word string)
	Insert(word string)
	GetFrequentUses() []string
}

/**
 * Constructor.
 */
func New(reader io.Reader, mapper mapperType, minWordLength int) *TextStructure {
	return &TextStructure{
		text:          reader,
		mapperType:    mapper,
		minWordLength: minWordLength,
	}
}

/**
 * Work with all word from Reader.
 */
func (t *TextStructure) GetWords() []string {
	s := bufio.NewScanner(t.text)
	lines := getPreparedLines(s)

	someList := t.mapperType
	for _, line := range lines {
		words := strings.Split(line, " ")
		wordsCount := len(words)

		for index, word := range words {
			// Remove word if it at the end of sentence
			if index == 0 || index == wordsCount-1 {
				someList.UpdateExcludeList(word)
				continue
			}

			if someList.IsWordExcluded(word) || !t.isWordValid(word) {
				continue
			}

			someList.Insert(word)
		}
	}

	return someList.GetFrequentUses()
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
 * Preparing sentences from read lines.
 */
func getPreparedLines(s *bufio.Scanner) []string {
	var lines []string
	for s.Scan() {
		line := strings.TrimSpace(regexp.MustCompile("[^a-zA-Z .0-9]+").ReplaceAllString(s.Text(), ""))

		if line == "" {
			continue
		}

		for _, splitLine := range strings.Split(line, ". ") {
			if strings.HasSuffix(splitLine, ".") {
				splitLine = strings.Replace(splitLine, ".", "", 1)
			}

			lines = append(lines, strings.TrimSpace(strings.ToLower(splitLine)))
		}
	}

	return lines
}

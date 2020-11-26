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
}

/**
 * Constructor.
 */
func New(reader io.Reader) *TextStructure {
	return &TextStructure{text: reader, minWordLength: 3}
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
			if index == 0 || index == wordsCount-1 || !t.isWordValid(word) {
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

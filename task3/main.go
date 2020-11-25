package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"tasks/task3/pkg/hardMapper"
	"unicode/utf8"
)

/**
 * Handle method.
 */
func main() {
	file, err := os.Open("./task3/files/some_text.txt")
	if err != nil {
		fmt.Println("Panic")
		return
	}

	defer file.Close()

	fmt.Println(GetWords(file))
}

/**
 * Work with all word from Reader.
 */
func GetWords(file io.Reader) []string {
	someList := hardMapper.New()

	s := bufio.NewScanner(file)
	lines := getPreparedLines(s)

	for _, line := range lines {
		words := strings.Split(line, " ")
		wordsCount := len(words)

		for index, word := range words {
			if index == 0 || index == wordsCount-1 || !isWordValid(word) {
				continue
			}
			someList.Insert(word)
		}

	}

	return someList.GetFrequentUses()
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
func isWordValid(word string) bool {
	if utf8.RuneCountInString(word) < 4 { // if length word less than 3 symbols
		return false
	}

	if _, err := strconv.Atoi(word); err == nil {
		return false
	}

	return true
}

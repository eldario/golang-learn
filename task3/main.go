package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
	"tasks/task3/pkg/mapper"
	"tasks/task3/pkg/reader"
)

type readLiner interface {
	Read(string, uint8)
}

// main Handle method
func main() {
	count, filePath, wordLength := parseFlags()

	file, err := os.Open(filePath)

	if err != nil {
		panic(err.Error())
	}

	defer file.Close()

	sortedMap := mapper.New(count)
	lineReader := reader.New(sortedMap, wordLength)

	content := bufio.NewScanner(file)
	var paragraphNumber uint8 = 1

	waitGroup := new(sync.WaitGroup)

	for content.Scan() {
		waitGroup.Add(1)
		go func(liner readLiner, content string, paragraphNumber uint8, wg *sync.WaitGroup) {
			defer wg.Done()

			liner.Read(content, paragraphNumber)
		}(lineReader, content.Text(), paragraphNumber, waitGroup)

		paragraphNumber++
	}

	waitGroup.Wait()

	for _, word := range sortedMap.GetResults() {
		fmt.Println(word.Word, word.Count, word.Score)
	}
}

// parseFlags Get parsed flags
func parseFlags() (int, string, int) {
	count := flag.Int("count", 10, "an int")
	filePath := flag.String("filepath", "files/some_text.txt", "File name")
	wordLength := flag.Int("minlength", 3, "File name")

	flag.Parse()
	return *count, *filePath, *wordLength
}

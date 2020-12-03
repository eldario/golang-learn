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
	Read(string, uint32)
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
	var paragraphNumber uint32

	out := make(chan struct{}, 3)

	var waitGroup sync.WaitGroup

	for content.Scan() {
		waitGroup.Add(1)
		out <- struct{}{}

		paragraphNumber++

		go func(liner readLiner, content string, paragraphNumber uint32) {

			defer waitGroup.Done()

			liner.Read(content, paragraphNumber)

			<-out
		}(lineReader, content.Text(), paragraphNumber)
	}

	waitGroup.Wait()

	for _, word := range sortedMap.GetResults() {
		fmt.Println(word.Word, word.Count)
	}
}

// parseFlags Get parsed flags
func parseFlags() (uint, string, int) {
	count := flag.Uint("count", 10, "an int")
	filePath := flag.String("filepath", "files/some_text.txt", "File name")
	wordLength := flag.Int("minlength", 3, "File name")

	flag.Parse()
	return *count, *filePath, *wordLength
}

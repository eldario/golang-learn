package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"tasks/task3/pkg/mapper/simple"
	"tasks/task3/pkg/reader"
)

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
	for content.Scan() {
		lineReader.Read(content.Text())
	}

	for _, word := range sortedMap.GetResults() {
		fmt.Println(word.Word, word.Count)
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

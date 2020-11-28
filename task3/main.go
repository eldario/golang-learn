package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
	"strings"
	"tasks/task3/pkg/reader"
)

/**
 * Handle method.
 */
func main() {
	// Get parsed flags
	count, filePath, wordLength := parseFlags()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Panic", err)
		return
	}

	defer file.Close()

	fileReader := reader.New(file).SetMinWordLength(wordLength)

	printInTable(fileReader.GetWords(count))

}

/**
 * Return flags from cli.
 */
func parseFlags() (int, string, int) {
	count := flag.Int("count", 10, "an int")
	filePath := flag.String("filepath", "files/some_text.txt", "File name")
	wordLength := flag.Int("minlength", 3, "File name")

	flag.Parse()
	return *count, *filePath, *wordLength
}

/**
 * Beauty print.
 */
func printInTable(result []string) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Word", "Count", "Order")

	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for index, word := range result {
		words := strings.Split(word, ";")

		tbl.AddRow(index+1, words[0], words[1], words[2])
	}

	tbl.Print()
}

package main

import (
	"fmt"
	"tasks/task4/pkg/pipe"
)

// main Handle
func main() {
	runPipePacket()
}

// runPipePacket try to run pipe packet
func runPipePacket() {
	firstChannel := make(chan string)

	pipe.FillValues([]string{"foo", "fooBar", "test", "Hello", "Winter"}, firstChannel)

	format := func(str string) string {
		return fmt.Sprintf("Length of the word [%s] is %d", str, len(str))
	}

	secondChannel := make(chan string)
	go func() {
		defer close(secondChannel)

		pipe.NewPipe(firstChannel, secondChannel, format)
	}()

	for value := range secondChannel {
		fmt.Println("Response:", value)
	}
}

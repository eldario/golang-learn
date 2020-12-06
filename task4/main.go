package main

import (
	"fmt"
	"tasks/task4/pkg/buffered"
	"tasks/task4/pkg/fan/in"
	"tasks/task4/pkg/fan/out"
	"tasks/task4/pkg/pipe"
)

// main Handle
func main() {
	runPipePacket()
	runFanInPacket()
	runFanOutPacket()
	runBufferedChan()
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

// runFanInPacket try to run fanIn packet
func runFanInPacket() {
	var channelList []<-chan string

	wordsList := [][]string{
		{"first", "second", "third", "fourth"},
		{"fifth", "sixth", "seventh", "eighth", "ninth", "tenth"},
		{"eleventh", "twelfth"},
		{"thirteenth", "fourteenth", "fifteenth"},
	}

	for _, words := range wordsList {
		channel := in.GenerateChannel(words)
		channelList = append(channelList, channel)
	}

	secondChannel := make(chan string)
	go func() {
		defer close(secondChannel)
		in.NewFanIn(channelList, secondChannel)
	}()

	for value := range secondChannel {
		fmt.Println("Value:", value)
	}
}

// runFanOutPacket try to run fanOut packet
func runFanOutPacket() {
	var channelList []chan string
	for i := 0; i < 3; i++ {
		channelList = append(channelList, make(chan string))
	}

	inChannel := make(chan string)
	out.InsertWordInChannel([]string{"foo", "fooBar", "test", "Hello", "Winter"}, inChannel)

	go func() {
		defer func() {
			for _, channel := range channelList {
				close(channel)
			}
		}()

		out.NewFanOut(inChannel, channelList)
	}()

	for {
		for index, channel := range channelList {
			select {
			case value, ok := <-channel:
				if !ok {
					return
				}
				fmt.Printf("Values [%s] from channel %d\n", value, index)
			}
		}
	}
}

// runBufferedChan try to run buffered packet
func runBufferedChan() {
	inChannel := make(chan string)

	go func() {
		defer close(inChannel)

		for i := 0; i < 10; i++ {
			inChannel <- fmt.Sprintf("IndexWord[%d]", i)
		}
	}()

	outChannel := buffered.NewChan(inChannel, 4)

	for value := range outChannel {
		fmt.Println("Value:", value)
	}

}

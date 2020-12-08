package main

import (
	"context"
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
	inChannel := make(chan string)
	outChannel := make(chan string)
	format := func(str string) string {
		return fmt.Sprintf("Length of the word [%s] is %d", str, len(str))
	}
	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	defer func() {
		fmt.Println("Call cancel function")
		cancelFunction()
	}()

	p := pipe.New(ctxWithCancel, inChannel, outChannel, format)

	p.FillValues([]string{"foo", "fooBar", "test", "Hello", "Winter"})

	go func() {
		defer close(outChannel)

		p.Run()
	}()

	for value := range outChannel {
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

	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	defer func() {
		fmt.Println("Call cancel function")
		cancelFunction()
	}()

	outChannel := make(chan string)
	fanIn := in.New(ctxWithCancel, channelList, outChannel)

	for _, words := range wordsList {
		channel := in.GenerateChannel(words)
		fanIn.Add(channel)
	}

	go func() {
		defer close(outChannel)
		fanIn.Run()
	}()

	for value := range outChannel {
		fmt.Println("Value:", value)
	}
}

// runFanOutPacket try to run fanOut packet
func runFanOutPacket() {
	var channelList []chan string

	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	defer func() {
		fmt.Println("Call cancel function")
		cancelFunction()
	}()

	inChannel := make(chan string)
	fanOut := out.New(ctxWithCancel, inChannel, channelList)

	for i := 0; i < 3; i++ {
		fanOut.Add(make(chan string))
	}

	out.InsertWordInChannel([]string{"foo", "fooBar", "test", "Hello", "Winter"}, inChannel)

	go func() {
		defer func() {
			for _, channel := range fanOut.OutChannels {
				close(channel)
			}
		}()

		fanOut.Run()
	}()

	for {
		for index, channel := range fanOut.OutChannels {
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

	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	defer func() {
		fmt.Println("Call cancel function")
		cancelFunction()
	}()

	bufferedChan := buffered.New(ctxWithCancel, inChannel, 1)

	go func() {
		bufferedChan.Run()
	}()

	for value := range bufferedChan.OutChannel {
		fmt.Println("Value:", value)
	}

}

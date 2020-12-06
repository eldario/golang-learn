package out

import "sync"

// NewFanOut get from one channel and send it to channels list
func NewFanOut(inChannel <-chan string, outChannels []chan string) {
	var waitGroup sync.WaitGroup

	for value := range inChannel {
		waitGroup.Add(1)
		go func(value string, outChannels []chan string) {
			defer waitGroup.Done()

			for _, channel := range outChannels {
				channel <- value
			}
		}(value, outChannels)
	}

	waitGroup.Wait()
}

// InsertWordInChannel insert a new word in channel
func InsertWordInChannel(words []string, inChannel chan<- string) {
	go func() {
		defer close(inChannel)

		for _, word := range words {
			inChannel <- word
		}
	}()
}

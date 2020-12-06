package in

import "sync"

// NewFanIn get all list of channels, collect data and send it out
func NewFanIn(channels []<-chan string, out chan string) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(channels))
	for _, channel := range channels {
		go func(channel <-chan string) {
			defer waitGroup.Done()

			for word := range channel {
				out <- word
			}
		}(channel)
	}

	waitGroup.Wait()
}

// GenerateChannel put words in generated channel and return it
func GenerateChannel(words []string) <-chan string {
	outChannel := make(chan string)

	go func() {
		defer close(outChannel)
		
		for _, word := range words {
			outChannel <- word
		}
	}()

	return outChannel

}

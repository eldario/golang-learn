package buffered

// NewChan put values into the channel that fit
func NewChan(inChannel <-chan string, bufferSize int) <-chan string {
	outChannel := make(chan string, bufferSize)

	go func(bufferSize int) {
		defer close(outChannel)

		var count = 0
		for word := range inChannel {
			if count == bufferSize {
				return
			}
			outChannel <- word
			count++
		}
	}(bufferSize)

	return outChannel
}

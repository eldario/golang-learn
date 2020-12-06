package pipe

// msgFunc type of format function
type msgFunc func(str string) string

// NewPipe re-sender method
func NewPipe(inChannel <-chan string, outChannel chan<- string, format msgFunc) {
	for value := range inChannel {
		outChannel <- format(value)
	}
}

// FillValues insert given list of words in channel
func FillValues(words []string, inChannel chan<- string) {
	go func() {
		defer close(inChannel)

		for _, word := range words {
			inChannel <- word
		}
	}()
}

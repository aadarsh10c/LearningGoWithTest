package iteration

//Takes a charachter and returns a 5 times repeated charachter as string
func Iterator(letter string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated = repeated + letter
	}
	return
}

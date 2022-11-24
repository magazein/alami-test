package helper

func AddToChan(channel chan []string, data [][]string) {
	// send each data to channel
	for i := 0; i < len(data); i++ {
		channel <- data[i]
	}

	// close channel
	close(channel)
}

func Average(data ...int) int {
	var count int

	// count each element
	for _, x := range data {
		count += x
	}

	// get average
	return count / len(data)
}

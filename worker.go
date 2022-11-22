package main

import "sync"

func proceedAsync(
	concurrency int,
	callback func(i int, data []string, target [][]string),
	source [][]string,
	target [][]string,
) {
	// count data
	numberOfData := len(source)

	// create buffered channel
	var avgChan = make(chan []string, numberOfData)
	go addToChan(avgChan, source)

	// create waitgroup
	var wg sync.WaitGroup
	wg.Add(concurrency)

	// Read channel and run the job based on max concurrency
	for i := 0; i < concurrency; i++ {
		go func(in int) {
			defer wg.Done()
			for each := range avgChan {
				callback(in, each, target)
			}
		}(i)
	}

	// Wait until all job done
	wg.Wait()
}

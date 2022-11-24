package helper

import (
	"sync"
)

func ProceedAsync(
	concurrency int,
	callback func(i int, data []string) error,
	source [][]string,
) []error {
	// count data
	numberOfData := len(source)

	// create buffered channel
	var avgChan = make(chan []string, numberOfData)
	go AddToChan(avgChan, source)

	// create waitgroup
	var wg sync.WaitGroup
	wg.Add(concurrency)

	var errs []error

	// Read channel and run the job based on max concurrency
	for i := 0; i < concurrency; i++ {
		go func(in int) {
			defer wg.Done()
			for each := range avgChan {
				err := callback(in, each)
				if err != nil {
					errs = append(errs, err)
				}
			}
		}(i)
	}

	// Wait until all job done
	wg.Wait()

	return errs
}

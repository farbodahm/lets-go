package main

import "testing"

// Test to see if workers are producing the expected number of result
func TestWorkersProduceExpectedNumberOfMessages(t *testing.T) {
	const n = 2    // Number of workers
	const itr = 10 // Number of jobs per worker

	ch := make(chan string)
	Wg.Add(n)
	go CloseWorkersChan(ch)

	// Running workers
	for i := 1; i <= n; i++ {
		go Worker(i, itr, ch)
	}

	recievedNum := 0
	// Receiving messages produced by workers
	for range ch {
		recievedNum++
	}

	if recievedNum != n*itr {
		t.Error(
			"Number of recieved messages should be exactly",
			n*itr, ", But recieved", recievedNum,
		)
	}
}

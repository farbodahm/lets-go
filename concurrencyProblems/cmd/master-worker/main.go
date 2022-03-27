package main

// Main consumer functionality
func main() {
	const n = 2    // Number of workers
	const itr = 10 // Number of jobs per worker

	ch := make(chan string)
	Wg.Add(n)
	go CloseWorkersChan(ch)

	// Running workers
	for i := 1; i <= n; i++ {
		go Worker(i, itr, ch)
	}

	// Receiving messages produced by workers
	for mesg := range ch {
		println(mesg)
	}
}

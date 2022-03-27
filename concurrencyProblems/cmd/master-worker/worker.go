package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Wg sync.WaitGroup // Used for `signaling` the consumer when workers finished their work and chan should be closed

// worker is creating a message and putting it on the given channel (Producer)
// it will generate exactly itr messages
func Worker(id int, itr int, ch chan<- string) {
	defer Wg.Done()

	for i := 0; i < itr; i++ {
		d := time.Duration(rand.Int63n(int64(time.Second))) // Random wait time up to 1 Second
		time.Sleep(d)

		ch <- fmt.Sprintf("Id %d: is on %d iteration.", id, i) // Message
	}
}

// closeWorkersChan will close the given channel when all of the workers have finished their job
func CloseWorkersChan(ch chan string) {
	Wg.Wait()
	close(ch)
}

// My Solution

package main

import (
	"sync"
)

var wg sync.WaitGroup

func closeChannel(ch chan []string) {
	wg.Wait()
	close(ch)
}

func getLinks(ch chan []string, path string) {
	defer wg.Done()
	tempResult := Retrieve(path)
	ch <- tempResult

	for _, v := range tempResult {
		wg.Add(1)
		go getLinks(ch, v)
	}
}

func Hack(start string) []string {
	ch := make(chan []string, 1000)
	finalResult := []string{}

	tempResult := Retrieve(start)
	finalResult = append(finalResult, tempResult...)

	for _, v := range tempResult {
		wg.Add(1)
		go getLinks(ch, v)
	}

	go closeChannel(ch)

	for v := range ch {
		finalResult = append(finalResult, v...)
	}
	return finalResult
}

package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	timeout := time.After(5 * time.Second)
	startPoint := "snapp.ir/"
	expected := []string{
		"snapp.ir/blog",
		"snapp.ir/blog/post1",
		"snapp.ir/blog/post2",
		"snapp.ir/contact",
		"snapp.ir/about",
	}

	done := make(chan bool)
	go func() {
		actual := Hack(startPoint)
		assert.ElementsMatch(t, expected, actual)
		done <- true
	}()
	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case <-done:
	}
}

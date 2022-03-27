package main

import (
	"log"
	"sync"
	"time"
)

var n int                   // Number of available breads
var student_lock sync.Mutex // For `synchronizing` students (At a time, just one student should pick a bread)

// baker charges the breads(n) so that students can continue eating.
// When it bought the breads, it will signal on use_bread channel.
func baker(get_bread, use_bread chan int) {
	for {
		<-get_bread
		log.Println("I'm buying bread!")
		time.Sleep(5 * time.Second)
		n = 10
		use_bread <- 1
	}
}

// student is the main worker. It can study only when we have enough breads.
func student(student_id int, get_bread, use_bread chan int) {
	for {
		student_lock.Lock()
		if n > 0 { // We have bread...
			log.Println(student_id, ": Eating bread. Bread=", n)
			n--
		} else {
			log.Println(student_id, ": Should wait for bread! Bread=", n)
			get_bread <- 1
			<-use_bread
			log.Println(student_id, ": Yay! We have bread to eat. Bread=", n)
			n--
		}
		student_lock.Unlock()

		log.Println(student_id, ": I'm going to study...")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	log.Println("Starting...")
	get_bread := make(chan int) // Used by students to `signal` baker to buy bread
	use_bread := make(chan int) // Used by baker to `signal` students to eat bread
	n = 10                      // Initial number of breads

	go baker(get_bread, use_bread)
	for i := 0; i < 10; i++ {
		go student(i, get_bread, use_bread)
	}

	time.Sleep(200 * time.Second)
}

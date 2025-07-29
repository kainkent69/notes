package main

import (
	"log"
	"sync"
	"time"
)

// todo
var mux sync.Mutex

func work(id int, trigger chan struct{}) {
	mux.Lock()
	log.Printf("Work Is Being Processed ID: %v\n", id)
	start := time.Now()
	time.Sleep(500 * time.Millisecond)
	end := time.Since(start)
	log.Printf("Work ID: %v is Finished in {{ %v }}\n", id, end)
	time.Sleep(2000 * time.Millisecond)
	mux.Unlock()

	if id == 9 {
		trigger <- struct{}{}
	} else {
		log.Printf("Id %d is initialized", id)
	}
}

func main() {
	log.Printf("Using Goroutines as the control flow\n ")
	data := make(chan struct{})
	
	for i := range 10 {
		log.Printf("A job is being created %d\n", i)
		time.Sleep(100 * time.Millisecond)
		go work(i, data)
	}
	<- data
	
}

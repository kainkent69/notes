package main

import (
	"log"
	"sync"
	"time"
)

// todo
var mux sync.Mutex

func work(id int, trigger chan struct{}) {
	<-trigger
	mux.Lock()
	log.Printf("Work Is Being Processed ID: %v\n", id)
	start := time.Now()
	time.Sleep(500 * time.Millisecond)
	end := time.Since(start)
	log.Printf("Work ID: %v is Finished in {{ %v }}\n", id, end)
	time.Sleep(2000 * time.Millisecond)
	mux.Unlock()
	if id < 9 {
		trigger <- struct{}{}
	}

}

func main() {
	log.Printf("Using Goroutines as the control flow ")
	data := make(chan struct{})
	for i := range 10 {
		time.Sleep(100 * time.Millisecond)
		go work(i, data)
	}
	data <- struct{}{}
}

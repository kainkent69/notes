package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)
// create WaitGroup
var wg sync.WaitGroup

// calculate the sum of two
func sumof(from, to int) int {
	data := 0
	wg.Add(1)
	for i := from; i < to; i++ {
		data += i
	}
	log.Printf("New some is added[%v][%v] => %d\n", from, to, data)
	defer wg.Done()
	return data
}

// create random  numbers to be used as sample data
func makeRand() (a, b int) {
	for {
		a = rand.Intn(100)
		b = rand.Intn(100)
		if a != b && (a > 0 || b > 0) {
			break
		}
	}

	if a > b {
		return b, a
	}
	return a, b
}


func main() {
	started := time.Now()
	wg = sync.WaitGroup{}
	for range 10 {
		a, b := makeRand()
		go sumof(a, b)
	}

	go func() {
		wg.Wait()
		log.Println("The Concurrent calls suceeded")
		return
	}()

	log.Printf("Finished... {{%v}}\n", time.Since(started)/time.Microsecond)
}

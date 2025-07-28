package main

import (
	"log"
	"time"
)

func work(id int, ch chan int) {
	log.Printf("work: %d started\n", id)
	time.Sleep(time.Millisecond * 500)
	log.Printf("work: %d finished\n", id)
	ch <- id // end
	

}
func main() {
	ch := make(chan int, 3)
	for i := range 3 {
		go work(i, ch)
	}
	
	for i:= range 3 {
		<-ch
 		log.Printf("task: %d is done\n", i)
	}
}

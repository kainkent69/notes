package main

import (
	"log"
	"time"
)

// signal to stop.
var signal bool

// pool
type WorkerPool struct {
	Count      uint32
	WorkersCh  []chan int
	DespatchCh chan int
	Notify     chan QueueNotify
	List       QueueList
}

// notifier data
type NotifyData struct {
	Id   int
	Data int
}

// notifier
type QueueNotify struct {
	Type int8 // 0 for add, 1 for fail and 2 for sucesss
	Data NotifyData
}

// the queueu
type QueueList struct {
	List     []int
	Failed   []int
	Finished []int
}

// work state - Finished: bool, result int
type WorkState struct {
	Finished bool
	WorkerId uint32
	Result   int
}

// the work type return
type WorkFunc func(data any) WorkState

// worker
func Worker(id int, pool *WorkerPool, work WorkFunc) {
	busy := false
	for {
		// reqest a data
		if busy == false {
			busy = true
			log.Printf("job %d: requesting\n", id)
			pool.DespatchCh <- id
			log.Printf("job %d: resquested\n", id)

		}

		response := <-pool.WorkersCh[id] // make it block

		log.Printf("job %d: responded\n", id)
		state := work(response)
		// do the state.
		switch {
		case state.Finished:
			pool.Notify <- QueueNotify{Type: 2, Data: struct {
				Id   int
				Data int
			}{id, id}}
			busy = false

		case !state.Finished:
			pool.Notify <- QueueNotify{Type: 2, Data: struct {
				Id   int
				Data int
			}{id, id}}
			busy = false

		}

	}
}

func (w *WorkerPool) New(work WorkFunc, workersCnt uint32) WorkerPool {
	workerCh := make([]chan int, workersCnt)
	// add the channels
	for i := range workersCnt {
		workerCh[i] = make(chan int)
	}

	// queue
	queueList := QueueList{
		List:     make([]int, 1),
		Failed:   make([]int, 1),
		Finished: make([]int, 1),
	}
	Pool := WorkerPool{
		Count:      workersCnt,
		WorkersCh:  workerCh,
		Notify:     make(chan QueueNotify),
		DespatchCh: make(chan int, workersCnt),
		List:       queueList,
	}

	log.Println("launching list go routine")
	go Pool.addRoutine()
	log.Println("launching workers")
	for i := range workersCnt {
		go Worker(int(i), &Pool, work)
	}
	log.Println("workers: launched")
	return Pool
}

// method for adding a new list
func (pool *WorkerPool) addRoutine() {
	signal = false
	list := pool.List

	end := func(notif QueueNotify) {
		// to end the loop
		if notif.Data.Id == 3 {
			signal = true
		}
	}
	for {
		select {
		// on requesting a new
		case id := <-pool.DespatchCh:
			log.Printf("recieved a request %d\n", id)
			pool.WorkersCh[id] <- id

		case notif := <-pool.Notify:
			log.Printf("has notification id(%d) = %d \n", notif.Data.Id, notif.Data.Data)
			// 			state = false
			current := list.List[notif.Data.Id]
			if list.List == nil {
				log.Panicln("List is nil")
			}
			// wrap as function to not exit on return
			func() {
				log.Println("processs notif")
				switch notif.Type {
				case 0:
					log.Println("new task to add")
					list.List = append(list.List, notif.Data.Data)
					return

				case 1:
					log.Println("a job has failed")
					list.Failed = append(list.Failed, current)
					end(notif)
					return

				case 2:
					log.Println("a job has suceed")
					list.Finished = append(list.Finished, current)
					end(notif)
					return

				default:
					log.Println("invalid notification")
					return

				}
			}()
		}
	}

}

func main() {
	// create a pool
	log.Println("creating worker pool")
	pool := (&WorkerPool{}).New(func(_data any) WorkState {
		data := _data.(int)
		state := WorkState{}
		time.Sleep(time.Second / 2)
		log.Printf("a work done: %d\n", data)
		state.Finished = true
		return state
	}, uint32(4))

	// create a new task
	log.Println("creating tasks")
	for i := range 4 {
		pool.Notify <- QueueNotify{
			Type: 0,
			Data: struct {
				Id   int
				Data int
			}{Id: i, Data: i}}
	}

	// an endless loop to keep main alive
	for !signal {
		log.Println("Stopping....")
		time.Sleep(time.Second * 1)
		log.Println("Stopped")

	}
}

// 1. Find all available tasked that isnt busy.
// 2. make a vallid job status {busy, data, id, Options}

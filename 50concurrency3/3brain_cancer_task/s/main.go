package main

import (
	"sync"
	"time"
)

func Process(i int) {
	println("processing ", i)

	time.Sleep(time.Second) // simulating a workload
}

func main() {
	var input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := &sync.Mutex{}

	for _, i := range input {
		m.Lock()
		go func(inter int) {
			// m.Unlock - ???
			Process(inter)
			m.Unlock()
		}(i)

	}

	time.Sleep(time.Second)
}

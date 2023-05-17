package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

func main() {
	raceUpdate()
	// 1 naive implementation with mutex
	// 2 optimal impl with swap
}

// some times we want any of concurrent values
func raceUpdate() {
	var value int64
	var wg sync.WaitGroup

	setRandomValue := func() {
		// it's not optimal to use defer in such simple func btw. It adds ~40ns to exec time
		defer wg.Done()
		atomic.SwapInt64(&value, rand.Int63())
		// atomic.CompareAndSwapInt64()
		// value = rand.Int63()
	}

	cnt := 100
	wg.Add(cnt)

	for i := 0; i < cnt; i++ {
		go setRandomValue()
	}

	wg.Wait()
	fmt.Println("value:", value)
}

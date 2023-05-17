package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

func poorDecision() {
	var value int64
	var wg sync.WaitGroup
	var mu sync.Mutex

	setRandomValue := func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		value = rand.Int63()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go setRandomValue()
	}

	wg.Wait()
	fmt.Println("Value:", value)

}

func swap() {
	var value int64
	var wg sync.WaitGroup

	setRandomValue := func() {
		newValue := rand.Int63()
		atomic.SwapInt64(&value, newValue)
		wg.Done()
	}

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go setRandomValue()
	}

	wg.Wait()
	fmt.Println("Value:", value)
}

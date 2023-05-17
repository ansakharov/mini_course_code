package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	storage := make(map[int]string)
	wg := sync.WaitGroup{}

	// pass mutex and waitgroup by pointer, otherwise they will be copied!
	mu := sync.RWMutex{}

	/*	go func() {
		mu.Lock()
		mu.Unlock()
	}()*/

	go writeToMap(storage, &wg, &mu)
	result := readFromMap(storage, &wg, &mu)

	wg.Wait()
	fmt.Println(result)
}

func writeToMap(in map[int]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	wg.Add(1)
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
		mu.Lock()
		in[i] = fmt.Sprintf("%d", i^2)
		mu.Unlock()
	}
}

func readFromMap(in map[int]string, wg *sync.WaitGroup, mu *sync.RWMutex) []string {
	result := make([]string, 0, 1000)

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		i := i
		go func() {
			defer wg.Done()

			time.Sleep(time.Millisecond)

			mu.RLock()
			temp, ok := in[i]
			mu.RUnlock()

			if ok {
				result = append(result, temp)
			}
		}()
	}

	return result
}

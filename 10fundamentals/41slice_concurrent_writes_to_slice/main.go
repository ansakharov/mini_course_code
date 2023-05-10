package main

import (
	"fmt"
	"sync"
)

func main() {
	unsafeWrites()
	correctWrites()
}

func unsafeWrites() {
	slice := []int{0, 0, 0, 0}
	// wg - brief introduction
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		for index := range slice {
			index := index
			wg.Add(1)

			go func() {
				defer wg.Done()
				slice[index]++
			}()
		}
	}

	wg.Wait()

	fmt.Println(slice)
}

func correctWrites() {
	slice := []int{0, 0, 0, 0}
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		for index := range slice {
			index := index
			wg.Add(1)

			go func() {
				defer wg.Done()
				mu.Lock()
				defer mu.Unlock()
				slice[index]++
			}()
		}
	}

	wg.Wait()

	fmt.Println(slice)
}

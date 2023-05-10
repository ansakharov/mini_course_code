package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		for key := range m {
			key := key
			wg.Add(1)
			go func() {
				defer wg.Done()
				m[key]++
			}()
		}
	}

	wg.Wait()

	fmt.Println("Map after concurrent writes:", m)
}

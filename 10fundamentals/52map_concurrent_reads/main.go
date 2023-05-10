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

	for key := range m {
		key := key
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%s: %d\n", key, m[key])
		}()
	}

	wg.Wait()
}

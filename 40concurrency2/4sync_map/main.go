package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup
	count := 100

	for i := 0; i < count; i++ {
		wg.Add(1)

		go func(key, value int) {
			defer wg.Done()
			m.Store(key, value)
		}(i, i*i)
	}

	wg.Wait()
	// len(m)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			value, ok := m.Load(key)
			if ok {
				fmt.Printf("k: %v, v: %v\n", key, value)
			}
		}(i)
	}

	wg.Wait()
	time.Sleep(time.Second * 5)

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			m.Delete(key)
		}(i)
	}

	wg.Wait()

	m.Range(func(key, value any) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
}

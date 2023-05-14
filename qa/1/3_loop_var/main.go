package main

import (
	"fmt"
	"sync"
)

// Можете чуть подробнее объяснить пример с указателями. Как value := value помогло?
func main() {
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		i := i
		go func(i int) {
			//i := 666
			defer wg.Done()

			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}

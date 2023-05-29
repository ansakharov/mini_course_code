package main

import (
	"fmt"
	"sync"
)

// Можно ли подробно разобрать отличие блокировки на мьютексе от блокировки на канале (P.S. Был такой вопрос на собесе)

//Мьютексы для shared memory

// channels для синхронизации горутин
func main() {
	var mu sync.Mutex
	storage := make(map[int]int)

	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		i := i

		// go store(i, storage, &group, &mu)

		/*
			go func() {
				defer group.Done()

				mu.Lock()
				defer mu.Unlock()
				storage[i] = i * i
			}()
		*/

		go func(idx int, in map[int]int, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			in[idx] = idx * idx

			return
		}(i, storage, &group, &mu)
	}

	group.Wait()
	fmt.Println(storage)

}

func store(idx int, in map[int]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()
	in[idx] = idx * idx
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Мы пытаемся подсчитать количество выполненных параллельно операций,
//что может пойти не так?

var callCounter uint64

func main() {
	group := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			// Ходим в базу, делаем долгую работу
			//time.Sleep(time.Second)
			//Увеличиваем счетчик
			atomic.AddUint64(&callCounter, 1)
			//callCounter++
		}()
	}

	group.Wait()
	fmt.Println("Call counter value = ", callCounter)
}

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, cancelFunc context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	if id == 1 {
		fmt.Printf("Worker %d finished with error\n", id)
		cancelFunc()

		return
	}

	select {
	case <-ctx.Done():
		fmt.Println("ctx done")
		return
	case <-time.After(time.Hour * 10):
		return
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	workers := 3
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		/*		if i == 1 {
				continue
			}*/
		wg.Add(1)

		go worker(ctx, i, cancel, &wg)
	}

	// Ждем завершения всех горутин
	wg.Wait()
}

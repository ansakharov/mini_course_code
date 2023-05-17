package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	wg := sync.WaitGroup{}
	wg.Add(math.MaxUint16)
	//mu := sync.Mutex{}

	// math.MaxUnit16 = 65535
	for i := 0; i < math.MaxUint16; i++ {
		go func() {
			defer wg.Done()
			//mu.Lock()
			//defer mu.Unlock()

			atomic.AddInt64(&counter, 1)
			// counter++
			go func() {
				_ = atomic.LoadInt64(&counter)
			}()
		}()
	}

	wg.Wait()
	fmt.Println(atomic.LoadInt64(&counter))
}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count := 100000
	var wg sync.WaitGroup

	var memStat runtime.MemStats
	runtime.ReadMemStats(&memStat)

	wg.Add(count)
	for i := 0; i < count; i++ {
		i := i
		go func() {
			defer wg.Done()
			_ = i * i
		}()
	}

	wg.Wait()
	var memStatNow runtime.MemStats

	runtime.ReadMemStats(&memStatNow)

	memConsumed := float64(memStatNow.Sys-memStat.Sys) / 1024 / 1024
	runtime.GC()

	fmt.Printf("Memory consumed by %d goroutines: %.2f MB\n", count, memConsumed)
}

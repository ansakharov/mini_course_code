package main

import (
	"sync"
	"testing"
)

/* В чем особенности, и когда стоит применять Mutex и RWMutex? */
func BenchmarkMutexEqual(b *testing.B) {
	storage := make(map[int]int)
	mu := sync.Mutex{}

	for i := 0; i < b.N; i++ {
		i := i
		go func() {
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		}()
		go func() {
			mu.Lock()
			_ = storage[i]
			mu.Unlock()
		}()
	}
}

func BenchmarkRWMutexEqual(b *testing.B) {
	storage := make(map[int]int)
	rwMu := sync.RWMutex{}

	for i := 0; i < b.N; i++ {
		i := i
		go func() {
			rwMu.Lock()
			storage[i] = i
			rwMu.Unlock()
		}()
		go func() {
			rwMu.RLock()
			_ = storage[i]
			rwMu.RUnlock()
		}()
	}
}

func BenchmarkMoreReads(b *testing.B) {
	storage := make(map[int]int)
	mu := sync.Mutex{}

	for i := 0; i < b.N; i++ {
		i := i

		if i%1000 == 0 {
			go func() {
				mu.Lock()
				storage[i] = i
				mu.Unlock()
			}()
		}

		go func() {
			mu.Lock()
			_ = storage[i]
			mu.Unlock()
		}()
	}
}

func BenchmarkRWMoreReads(b *testing.B) {
	storage := make(map[int]int)
	rwMu := sync.RWMutex{}

	for i := 0; i < b.N; i++ {
		i := i
		if i%1000 == 0 {
			go func() {
				rwMu.Lock()
				storage[i] = i
				rwMu.Unlock()
			}()
		}

		go func() {
			rwMu.RLock()
			_ = storage[i]
			rwMu.RUnlock()
		}()
	}
}

func BenchmarkMoreWrites(b *testing.B) {
	storage := make(map[int]int)
	mu := sync.Mutex{}

	for i := 0; i < b.N; i++ {
		i := i

		go func() {
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		}()

		if i%1000 == 0 {
			go func() {
				mu.Lock()
				_ = storage[i]
				mu.Unlock()
			}()
		}
	}
}

func BenchmarkRWMutexMoreWrites(b *testing.B) {
	storage := make(map[int]int)
	rwMu := sync.RWMutex{}

	for i := 0; i < b.N; i++ {
		i := i

		go func() {
			rwMu.Lock()
			storage[i] = i
			rwMu.Unlock()
		}()

		if i%1000 == 0 {
			go func() {
				rwMu.RLock()
				_ = storage[i]
				rwMu.RUnlock()
			}()
		}
	}
}

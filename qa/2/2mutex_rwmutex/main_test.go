package main

import (
	"sync"
	"testing"
)

var m = make(map[int]int)
var mutex = sync.Mutex{}
var rwMutex = sync.RWMutex{}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			mutex.Lock()
			m[i] = i
			mutex.Unlock()
		}()
		go func() {
			mutex.Lock()
			_ = m[i]
			mutex.Unlock()
		}()
	}
}

func BenchmarkRWMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			rwMutex.Lock()
			m[i] = i
			rwMutex.Unlock()
		}()
		go func() {
			rwMutex.RLock()
			_ = m[i]
			rwMutex.RUnlock()
		}()
	}
}

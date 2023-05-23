package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

func main() {
	var mem runtime.MemStats

	for {
		runtime.GC()
		res, err := fetch()
		if err != nil {
			_ = err
		}

		_ = res

		runtime.ReadMemStats(&mem)

		fmt.Printf("Allocated memory: %d bytes\n", mem.Alloc)
		fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
		time.Sleep(time.Microsecond)
	}
}

type result struct {
	value string
	err   error
}

func fetch() ([]string, error) {
	var results []string
	var count = 10

	data := make(chan result)
	sem := make(chan struct{}, 3)

	for i := 0; i < count; i++ {
		d := time.Now()
		go func(date time.Time) {
			sem <- struct{}{}
			{
				val, err := do(date)
				data <- result{val, err}
			}
			<-sem
		}(d)
	}

	for count > 0 {
		count--
		r := <-data

		if r.err != nil {
			return results, r.err
		}

		results = append(results, r.value)
	}
	return results, nil
}

// todo mvp implemetation, i'll write code late
func do(t time.Time) (string, error) {
	return t.String(), errors.New("err1")
}

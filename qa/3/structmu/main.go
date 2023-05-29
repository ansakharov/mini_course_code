package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"

	"golang.org/x/sync/errgroup"
)

// если саму структуру передаем по указателю, то мьютекс в структуре не нужно передавать по указателю
type test struct {
	id      int
	storage map[int]struct{}
	name    string
	mu      sync.Mutex
}

func t(in *test) {
	in.mu.Lock()
	defer in.mu.Unlock()
	//fmt.Println(in.id)
	in.storage[rand.Intn(math.MaxInt64)] = struct{}{}
}

func main() {
	group := errgroup.Group{}
	testStr := test{
		id:      1,
		name:    "qwe",
		storage: make(map[int]struct{}),
	}

	for i := 0; i < 100; i++ {
		group.Go(func() error {
			t(&testStr)

			return nil
		})
	}

	fmt.Println(group.Wait())
}

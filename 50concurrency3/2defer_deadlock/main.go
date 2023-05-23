package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sync"
)

type inType int64

const (
	invalid  inType = 0
	base     inType = 1
	extended inType = 2
	bonus    inType = 3
)

func main() {
	mu := &sync.Mutex{}

	fmt.Println(trickyFunc(invalid, mu))
	fmt.Println(trickyFunc(base, mu))
	fmt.Println(trickyFunc(extended, mu))
	fmt.Println(trickyFunc(inType(99), mu))
	fmt.Println(trickyFunc(bonus, mu))
	fmt.Println(trickyFunc(base, mu))
}

func trickyFunc(in inType, mu *sync.Mutex) (int64, error) {
	mu.Lock()

	switch in {
	case invalid:
		mu.Unlock()
		return 0, errors.New("invalid type")
	case base:
		res := goToService2()
		mu.Unlock()
		return res, nil
	case extended:
		res := goToService1()
		mu.Unlock()
		return res, nil
	case bonus:
		res, err := goToService3()
		if err != nil {
			return 0, fmt.Errorf("goToService3 err: %w", err)
		}
		mu.Unlock()
		return res, nil
	default:
		mu.Unlock()
		return 0, errors.New("unsupported type")
	}

}

func goToService1() int64 {
	return rand.Int63n(math.MaxInt64)
}

func goToService2() int64 {
	return rand.Int63n(math.MaxInt64)
}

func goToService3() (int64, error) {
	return rand.Int63n(math.MaxInt64), errors.New("server went down")
}

package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// don't update TestRandFunc, it's for tests
var TestRandFunc func(n int64) int64 = rand.Int63n

var ErrService = errors.New("service error")
var ErrTimeout = errors.New("timeout ended")

type order struct {
	value   int64
	id      int64
	fee     int64
	itemIDs []int64
}

func (o order) String() string {
	return fmt.Sprintf(" id: %d, value: %d, fee: %d, itemIDs: %v", o.id, o.value, o.fee, o.itemIDs)
}

type Service struct {
	idsToFail map[int64]struct{}
}

func New(idsThatReturnsError map[int64]struct{}) *Service {
	if idsThatReturnsError == nil {
		idsThatReturnsError = make(map[int64]struct{})
	}
	return &Service{
		idsToFail: idsThatReturnsError,
	}
}

func (s *Service) getOrderByID(id int64) (*order, error) {
	time.Sleep(time.Duration(TestRandFunc(2000)) * time.Millisecond)

	// some user ids fails
	if _, ok := s.idsToFail[id]; ok {
		return nil, ErrService
	}

	return &order{
		value: int64(rand.Intn(1000)),
		id:    id,
		fee:   int64(rand.Intn(200)),
		itemIDs: func() []int64 {
			res := make([]int64, 10)
			for i := 0; i < 10; i++ {
				res[i] = int64(i * rand.Intn(30))
			}
			return res
		}(),
	}, nil
}

func (s *Service) getOrderByIDWrapper(ctx context.Context, id int64) (*order, error) {
	return s.getOrderByID(id)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	usersWithErr := map[int64]struct{}{
		1: {},
		2: {},
		3: {},
	}

	service := New(usersWithErr)

	// user with err
	fmt.Println(service.getOrderByID(1))
	// user with err
	fmt.Println(service.getOrderByID(2))
	// user with err
	fmt.Println(service.getOrderByID(3))

	fmt.Println(service.getOrderByID(4))
	fmt.Println(service.getOrderByID(5))

	for i := 0; i < 100; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
		result, err := service.getOrderByIDWrapper(ctx, 5)
		
		fmt.Println(result, err)

		cancel()
	}

}

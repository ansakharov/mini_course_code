package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
	state State
}

type State int

const (
	InitialState State = iota
	FirstStage
	SecondStage
	FinishedStage
)

const numWorkers = 40

func FirstProcessing(in <-chan job, out chan<- job, wg *sync.WaitGroup) {
	for j := range in {
		j.value = int64(float64(j.value) * math.Pi)
		j.state = FirstStage
		out <- j
	}
	wg.Done()
}

func SecondProcessing(in <-chan job, out chan<- job, wg *sync.WaitGroup) {
	for j := range in {
		j.value = int64(float64(j.value) * math.E)
		j.state = SecondStage
		out <- j
	}
	wg.Done()
}

func LastProcessing(in <-chan job, out chan<- job, wg *sync.WaitGroup) {
	for j := range in {
		j.value = int64(float64(j.value) / float64(rand.Intn(10)))
		j.state = FinishedStage

		out <- j
	}
	wg.Done()
}

func main() {
	length := 50_000_00
	jobs := make([]job, length)
	in := make(chan job, len(jobs))
	for i := 0; i < length; i++ {
		jobs[i].value = int64(i)
		in <- jobs[i]
	}
	close(in)

	start := time.Now()

	stage1 := make(chan job, 10000)
	stage2 := make(chan job, 10000)
	stage3 := make(chan job, 10000)

	var wg sync.WaitGroup

	// First stage
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go FirstProcessing(in, stage1, &wg)
	}
	go func() {
		wg.Wait()
		close(stage1)
	}()

	// Second stage
	var wg2 sync.WaitGroup
	wg2.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go SecondProcessing(stage1, stage2, &wg2)
	}
	go func() {
		wg2.Wait()
		close(stage2)
	}()

	// Third stage
	var wg3 sync.WaitGroup
	wg3.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go LastProcessing(stage2, stage3, &wg3)
	}
	go func() {
		wg3.Wait()
		close(stage3)
	}()

	for _ = range stage3 {
		// This is just to consume the output and not block stage3.
	}

	finished := time.Since(start)

	fmt.Println(finished)
}

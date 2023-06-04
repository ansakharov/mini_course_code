package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

// replace batch processing with stream processing
// now program waits until all process functions completes
// it should return jobs concurrently, after processing each job
// use channels
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

func FirstProcessingAsync(jobs chan job) chan job {
	result := make(chan job, 8096)

	go func() {
		for job := range jobs {
			job.value = int64(float64(job.value) * math.Pi)
			job.state = FirstStage
			result <- job
		}
		close(result)
	}()

	return result
}

func SecondProcessingAsync(jobs chan job) chan job {
	result := make(chan job, 8096)

	go func() {
		for job := range jobs {
			job.value = int64(float64(job.value) * math.E)
			job.state = SecondStage
			result <- job
		}
		close(result)
	}()

	return result
}

func LastProcessingAsync(jobs chan job) chan job {
	result := make(chan job, 8096)

	go func() {
		for job := range jobs {
			job.value = int64(float64(job.value) / float64(rand.Intn(10)+1))
			job.state = FinishedStage
			result <- job
		}
		close(result)
	}()

	return result
}

func FirstProcessing(jobs []job) []job {
	var result []job
	for _, job := range jobs {
		job.value = int64(float64(job.value) * math.Pi)
		job.state = FirstStage
		result = append(result, job)
	}

	return result
}

func SecondProcessing(jobs []job) []job {
	var result []job
	for _, job := range jobs {
		job.value = int64(float64(job.value) * math.E)
		job.state = SecondStage
		result = append(result, job)
	}

	return result
}

func LastProcessing(jobs []job) []job {
	var result []job
	for _, job := range jobs {
		job.value = int64(float64(job.value) / float64(rand.Intn(10)+1))
		job.state = FinishedStage
		result = append(result, job)
	}

	return result
}

func main() {
	length := 5_000_000
	jobs := make([]job, length)
	jobsChan := make(chan job, length)
	for i := 0; i < length; i++ {
		jobs[i].value = int64(i)
		jobsChan <- jobs[i]
	}

	close(jobsChan)

	start := time.Now()
	jobs = LastProcessing(
		SecondProcessing(
			FirstProcessing(jobs),
		),
	)
	for idx, value := range jobs {
		_, _ = idx, value
	}
	finished := time.Since(start)
	fmt.Println(finished)

	startAsync := time.Now()
	result := LastProcessingAsync(SecondProcessingAsync(FirstProcessingAsync(jobsChan)))

	fmt.Println(time.Since(startAsync))

	g := errgroup.Group{}

	g.SetLimit(100)
	for i := 0; i < 150; i++ {
		g.Go(func() error {
			for value := range result {
				_ = value
			}

			return nil
		})
	}

	_ = g.Wait()

	fmt.Println(time.Since(startAsync))
}

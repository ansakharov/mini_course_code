package src

import (
	"context"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	// you can limit concurrent net request. It's optional
	MaxGoroutines = 1
	// timeout for net requests
	Timeout = 2 * time.Second
)

type SiteStatus struct {
	Name          string
	StatusCode    int
	TimeOfRequest time.Time
}

type Monitor struct {
	StatusMap        map[string]SiteStatus
	Mtx              *sync.Mutex
	G                errgroup.Group
	Sites            []string
	RequestFrequency time.Duration
}

func NewMonitor(sites []string, requestFrequency time.Duration) *Monitor {
	return &Monitor{
		StatusMap:        make(map[string]SiteStatus),
		Mtx:              &sync.Mutex{},
		Sites:            sites,
		RequestFrequency: requestFrequency,
	}
}

func (m *Monitor) Run(ctx context.Context) error {
	// run printStatuses and checkSite in different goroutines
	return nil
}

func (m *Monitor) checkSite(ctx context.Context, site string) {
	// check site and write result to StatusMap
}
func (m *Monitor) printStatuses(ctx context.Context) error {
	// iterate iver map and print results
	return nil
}

package src

import (
	"context"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	// You can limit concurrent net request. It's optional
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
	// Run printStatuses(ctx) and checkSite(ctx) for m.Sites
	// Renew sites requests to map every m.RequestFrequency
	// Return if context closed
	return nil
}

func (m *Monitor) checkSite(ctx context.Context, site string) {
	// with http client go through site and write result to m.StatusMap
}
func (m *Monitor) printStatuses(ctx context.Context) error {
	// print results of m.Status every second of until ctx cancelled

	return nil
}

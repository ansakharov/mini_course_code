package src

import (
	"context"
	"fmt"
	"net/http"
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
	m.G.Go(func() error {
		return m.printStatuses(ctx)
	})

	for _, site := range m.Sites {
		site := site
		m.G.Go(func() error {
			ticker := time.NewTicker(m.RequestFrequency)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-ticker.C:
					m.checkSite(ctx, site)
				}
			}
		})
	}

	if err := m.G.Wait(); err != nil && err != context.Canceled {
		return err
	}

	return nil
}

func (m *Monitor) checkSite(ctx context.Context, site string) {
	client := http.Client{
		Timeout: Timeout,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, site, nil)
	if err != nil {
		fmt.Printf("Ошибка при создании запроса для %s: %s\n", site, err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Ошибка при получении данных с %s: %s\n", site, err)
		return
	}
	defer resp.Body.Close()

	m.Mtx.Lock()
	m.StatusMap[site] = SiteStatus{Name: site, StatusCode: resp.StatusCode, TimeOfRequest: time.Now().Truncate(time.Second)}
	m.Mtx.Unlock()
}
func (m *Monitor) printStatuses(ctx context.Context) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			m.Mtx.Lock()
			for _, status := range m.StatusMap {
				fmt.Printf("%s %d %v\n", status.Name, status.StatusCode, status.TimeOfRequest)
			}
			fmt.Println()
			m.Mtx.Unlock()
		}
	}
}

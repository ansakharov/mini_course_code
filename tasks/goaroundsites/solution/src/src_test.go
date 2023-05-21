package src

import (
	"context"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestMonitor_Run(t *testing.T) {
	type fields struct {
		StatusMap        map[string]SiteStatus
		Mtx              *sync.Mutex
		Sites            []string
		RequestFrequency time.Duration
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test Run",
			fields: fields{
				StatusMap: make(map[string]SiteStatus),
				Mtx:       &sync.Mutex{},
				Sites: []string{
					"https://www.google.com",
					"https://www.ya.ru",
				},
				RequestFrequency: time.Millisecond * 500,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Monitor{
				StatusMap:        tt.fields.StatusMap,
				Mtx:              tt.fields.Mtx,
				Sites:            tt.fields.Sites,
				RequestFrequency: tt.fields.RequestFrequency,
				G:                errgroup.Group{},
			}
			ctx, cancel := context.WithTimeout(tt.args.ctx, 10*time.Second)
			defer cancel()

			err := m.Run(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && err != context.DeadlineExceeded {
				t.Errorf("Expected error context deadline exceeded, but got %v", err)
			}

			// Check whether the map is populated
			if len(m.StatusMap) == 0 {
				t.Errorf("Status map is empty")
			}

			for _, site := range tt.fields.Sites {
				status, ok := m.StatusMap[site]
				if !ok {
					t.Errorf("No status found for site %s", site)
					continue
				}

				if status.StatusCode == 0 {
					t.Errorf("Expected valid status  for site %s, but got %d", site, status.StatusCode)
				}

				if time.Since(status.TimeOfRequest) > 3*time.Second {
					t.Errorf("Status for site %s is older than 3 seconds", site)
				}
			}
		})
	}
}

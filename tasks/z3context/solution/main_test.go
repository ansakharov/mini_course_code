package main

import (
	"context"
	"math/rand"
	"testing"
	"time"
)

// Run test from console with: go test -timeout 10
func TestService_getOrderByIDWrapper(t *testing.T) {
	type fields struct {
		idsToFail map[int64]struct{}
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     *order
		wantErr  bool
		randFunc func(n int64) int64
	}{
		{
			name: "test with error",
			fields: fields{
				idsToFail: map[int64]struct{}{1: {}},
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr:  true,
			randFunc: rand.Int63n,
		},
		{
			name: "test with no error",
			fields: fields{
				idsToFail: map[int64]struct{}{},
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: false,
			randFunc: func(n int64) int64 {
				// server responds veeeeery fast
				return 0
			},
		},
		{
			name: "test with timeout",
			fields: fields{
				idsToFail: map[int64]struct{}{},
			},
			args: args{
				ctx: func() context.Context {
					ctx, _ := context.WithTimeout(context.Background(), time.Second)

					return ctx
				}(),
				id: 99,
			},
			wantErr: true,
			randFunc: func(n int64) int64 {
				// server responds veeeeery slow
				return 36000000
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmp := TestRandFunc
			TestRandFunc = tt.randFunc

			defer func() {
				TestRandFunc = tmp
			}()
			s := &Service{
				idsToFail: tt.fields.idsToFail,
			}
			_, err := s.getOrderByIDWrapper(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOrderByIDWrapper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

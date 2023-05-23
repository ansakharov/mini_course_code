package main

import (
	"reflect"
	"testing"
)

func Test_ringBuffer_Run(t *testing.T) {
	type args struct {
		inputs      []int
		outChanSize int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want2 []int
	}{
		{
			name: "Test case 1: Buffer size 1, input 0-9",
			args: args{
				inputs:      []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				outChanSize: 1,
			},
			want:  []int{9},
			want2: []int{8, 9},
		},
		{
			name: "Test case 2: Buffer size 4, input 0-9",
			args: args{
				inputs:      []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				outChanSize: 4,
			},
			want:  []int{6, 7, 8, 9},
			want2: []int{5, 6, 7, 8, 9},
		},
		{
			name: "Test case 3: Buffer size 1, single input",
			args: args{
				inputs:      []int{0},
				outChanSize: 1,
			},
			want:  []int{0},
			want2: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inCh := make(chan int)
			outCh := make(chan int, tt.args.outChanSize)
			rb := NewRingBuffer(inCh, outCh)
			go rb.Run()

			for _, v := range tt.args.inputs {
				inCh <- v
			}
			close(inCh)

			got := make([]int, 0)
			for res := range outCh {
				got = append(got, res)
			}

			if !reflect.DeepEqual(got, tt.want) && !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("ringBuffer.Run() = %v, want %v or at least %v", got, tt.want, tt.want2)
			}
		})
	}
}

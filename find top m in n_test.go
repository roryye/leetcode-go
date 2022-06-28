package main

import "testing"

func Test_findTopMInN(t *testing.T) {
	type args struct {
		n []int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: []int{3, 1, 0, 2},
				m: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTopMInN(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("findTopMInN() = %v, want %v", got, tt.want)
			}
		})
	}
}

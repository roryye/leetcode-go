package main

import "testing"

func TestFibonacciIteration(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 6,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibonacciIteration(tt.args.n); got != tt.want {
				t.Errorf("FibonacciIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

package sort

import (
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				a: []int{4, 2, 1, 5, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapSort(tt.args.a)
			fmt.Println(tt.args.a)
		})
	}
}

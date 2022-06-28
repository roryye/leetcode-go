package sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	type args struct {
		n     []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				n:     []int{2, 1, 4, 3, 5},
				left:  0,
				right: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.n, tt.args.left, tt.args.right)
			fmt.Println(tt.args.n)
		})
	}
}

package sort

import (
	"testing"
)

func Test_linkQuickSort(t *testing.T) {
	type args struct {
		start *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				start: &Node{
					Val: 4,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val: 5,
							Next: &Node{
								Val: 1,
								Next: &Node{
									Val: 3,
								},
							},
						},
					},
				},
			},
		},
		{

			args: args{
				start: &Node{
					Val: 2,
					Next: &Node{
						Val: 1,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := linkQuickSort(tt.args.start)
			printLink(res)
		})
	}
}

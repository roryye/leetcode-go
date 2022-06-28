package sort

import "testing"

func Test_quickSelect(t *testing.T) {
	type args struct {
		n     []int
		left  int
		right int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: []int{2, 1, 4, 3, 5},
				k: 2,
			},
			want: 4,
		},
		{
			args: args{
				n: []int{2, 1, 4, 3, 5},
				k: 1,
			},
			want: 5,
		},
		{
			args: args{
				n: []int{2, 1, 4, 3, 5},
				k: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 第 k 大：k = len(tt.args.n)-tt.args.k)
			// 第 k 小：k = tt.args.k - 1
			if got := quickSelect(tt.args.n, 0, len(tt.args.n)-1, len(tt.args.n)-tt.args.k); got != tt.want {
				t.Errorf("quickSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}

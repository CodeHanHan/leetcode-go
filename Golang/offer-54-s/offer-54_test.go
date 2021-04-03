package offer_54_s_

import (
	"testing"
)

func Test_kthLargest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				root: buildTree([]int{3, 1, 2, 4}, []int{1, 2, 3, 4}),
				k:    1,
			},
			want:4,
		},
		{
			name: "1",
			args: args{
				root: buildTree([]int{5, 3, 2, 1, 4, 6}, []int{1, 2, 3, 4, 5, 6}),
				k:    3,
			},
			want:4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"test 1",
			args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val:   7,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   2,
								Left:  nil,
								Right: nil,
							},
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val:   13,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val:   5,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
				sum: 22,
			},
			[][]int{{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

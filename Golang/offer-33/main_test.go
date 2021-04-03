package main

import "testing"

func Test_verifyPostorder(t *testing.T) {
	type args struct {
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"1 test， 没有左子树",
			args{[]int{6, 5}},
			true,
		},
		{
			"2 test",
			args{[]int{1, 3, 2, 6, 5}},
			true,
		},
		{
			"3 test",
			args{[]int{1, 6, 3, 2, 5}},
			false,
		},
		{
			"4 test",
			args{[]int{1, 3, 2, 7, 6, 5}},
			true,
		},
		{
			"5 test， 没有右子树",
			args{[]int{2, 5}},
			true,
		},
		{
			"6 test，空",
			args{[]int{}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPostorder(tt.args.postorder); got != tt.want {
				t.Errorf("verifyPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

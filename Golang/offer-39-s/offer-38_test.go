package offer_39_s_

import "testing"

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{nums: []int{1, 2, 3, 2, 2, 2, 5, 4, 2}},
			want: 2,
		},
		{
			args: args{nums: []int{3, 2, 3}},
			want: 3,
		},
		{
			args: args{nums: []int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

package offer_53_II_s_

import "testing"

func Test_missingNumber(t *testing.T) {
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
			name: "1",
			args: args{nums: []int{0, 1, 3}},
			want: 2,
		},
		{
			name: "2",
			args: args{nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 9}},
			want: 8,
		},
		{
			name: "3",
			args: args{nums: []int{1}},
			want: 0,
		},
		{
			name: "4",
			args: args{nums: []int{0}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

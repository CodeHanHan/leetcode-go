package offer_51_d_

import "testing"

func Test_reversePairs(t *testing.T) {
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
			name: "[7,5,6,4]",
			args: args{nums: []int{7, 5, 6, 4}},
			want: 5,
		},
		{
			name: "[7,6,5,4,3,2,1]",
			args: args{nums: []int{7, 6, 5, 4, 3, 2, 1}},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePairs(tt.args.nums); got != tt.want {
				t.Errorf("reversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

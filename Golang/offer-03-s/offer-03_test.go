package offer_03_s

import "testing"

func Test_findRepeatNumber(t *testing.T) {
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
			args: args{nums: []int{2, 3, 1, 0, 2, 5, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

package offer_57_I_s

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{2, 7},
		},
		{
			name: "1",
			args: args{
				nums:   []int{10, 26, 30, 31, 47, 60},
				target: 40,
			},
			want: []int{10, 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

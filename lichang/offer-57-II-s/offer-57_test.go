package offer_57_II_s

import (
	"reflect"
	"testing"
)

func Test_findContinuousSequence(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{target: 9},
			want: [][]int{{2, 3, 4}, {4, 5}},
		},
		{
			name: "2",
			args: args{target: 15},
			want: [][]int{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContinuousSequence(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findContinuousSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

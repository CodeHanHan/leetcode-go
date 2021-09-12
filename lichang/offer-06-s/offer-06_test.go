package offer_06_s

import (
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{head: BuildList([]int{1, 2, 3, 4, 5})},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "2",
			args: args{head: BuildList([]int{1, 3, 2})},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}

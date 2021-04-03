package offer_38

import (
	"reflect"
	"testing"
)

func Test_permutation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "2 char",
			args: args{s: "ab"},
			want: []string{"ab", "ba"},
		},
		{
			name: "3 char",
			args: args{s: "abc"},
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name: "包含重复",
			args: args{s: "abb"},
			want: []string{"abb", "bab", "bba"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

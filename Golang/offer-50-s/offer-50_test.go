package offer_50_s_

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{
			name: "empty",
			args: args{s: ""},
			want: ' ',
		},
		{
			name: "abaccdeff",
			args: args{s: "abaccdeff"},
			want: 'b',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

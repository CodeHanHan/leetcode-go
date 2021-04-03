package offer_10_II_s

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "7",
			args: args{n: 7},
			want: 21,
		},
		{
			name: "0",
			args: args{n: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

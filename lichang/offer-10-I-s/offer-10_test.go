package offer_10_I_s

import "testing"

func Test_fib(t *testing.T) {
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
			name: "1",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "2",
			args: args{n: 5},
			want: 5,
		},
		{
			name: "3",
			args: args{n: 99},
			want: 94208912,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

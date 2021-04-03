package offer_53_s_

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: 2,
		},
		{
			name: "1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: 2,
		},
		{
			name: "2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				nums:   []int{2, 2},
				target: 3,
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				nums:   []int{2},
				target: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search1() = %v, want %v", got, tt.want)
			}
		})
	}
}

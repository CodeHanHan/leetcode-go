package offer_46

import "testing"

func Test_translateNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "12258",
			args: args{num: 12258},
			want: 5,
		},
		{
			name: "506",
			args: args{num: 506},
			want: 1,
		},
		{
			name: "25",
			args: args{num: 25},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNum(tt.args.num); got != tt.want {
				t.Errorf("translateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateNumDP(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "12258",
			args: args{num: 12258},
			want: 5,
		},
		{
			name: "506",
			args: args{num: 506},
			want: 1,
		},
		{
			name: "25",
			args: args{num: 25},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNumDP(tt.args.num); got != tt.want {
				t.Errorf("translateNumDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

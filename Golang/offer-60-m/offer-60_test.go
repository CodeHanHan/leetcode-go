package offer_60_m

import (
	"reflect"
	"testing"
)

func Test_dicesProbability(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
		//{
		//	name: "1",
		//	args: args{n: 1},
		//	want: []float64{0.16667,0.16667,0.16667,0.16667,0.16667,0.16667},
		//},
		{
			name: "2",
			args: args{n: 2},
			want: []float64{0.02778, 0.05556, 0.08333, 0.11111, 0.13889, 0.16667, 0.13889, 0.11111, 0.08333, 0.05556, 0.02778},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dicesProbability1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dicesProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}

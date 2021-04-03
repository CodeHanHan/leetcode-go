package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MinStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_Min(t *testing.T) {
	type fields struct {
		data []int
		min  []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				data: tt.fields.data,
				min:  tt.fields.min,
			}
			if got := this.Min(); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_Pop(t *testing.T) {
	type fields struct {
		data []int
		min  []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				data: tt.fields.data,
				min:  tt.fields.min,
			}
		})
	}
}

func TestMinStack_Push(t *testing.T) {
	type fields struct {
		data []int
		min  []int
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				data: tt.fields.data,
				min:  tt.fields.min,
			}
		})
	}
}

func TestMinStack_Top(t *testing.T) {
	type fields struct {
		data []int
		min  []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{
				data: tt.fields.data,
				min:  tt.fields.min,
			}
			if got := this.Top(); got != tt.want {
				t.Errorf("Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			"test_1",
			args{&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil},
			}},
			[][]int{{1}, {2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestQueue_clear(t *testing.T) {
//    type fields struct {
//        val  []*TreeNode
//        tail int
//        head int
//    }
//    tests := []struct {
//        name   string
//        fields fields
//    }{
//        // TODO: Add test cases.
//    }
//    for _, tt := range tests {
//        t.Run(tt.name, func(t *testing.T) {
//            queue := &Queue{
//                val:  tt.fields.val,
//                tail: tt.fields.tail,
//                head: tt.fields.head,
//            }
//        })
//    }
//}

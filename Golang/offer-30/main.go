package main

import "fmt"

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{}, //指向最小值
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.min) <= 0 {
		this.min = append(this.min, 0)
		return
	}
	if this.data[this.min[len(this.min)-1]] >= x {
		this.min = append(this.min, len(this.data)-1)
		return
	}
}

func (this *MinStack) Pop() { //min需要更新
	if len(this.data) <= 0 {
		panic("empty stack")
	}
	if this.min[len(this.min)-1] == len(this.data)-1 { //如果最小值是需被弹出的元素
		this.min = this.min[0 : len(this.min)-1] //使用上一个最小值
	}
	this.data = this.data[0 : len(this.data)-1] //弹出
	return
}

func (this *MinStack) Top() int {
	if len(this.data) <= 0 {
		panic("empty stack")
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	if len(this.data) <= 0 {
		panic("empty stack")
	}
	return this.data[this.min[len(this.min)-1]]
}

func main() {
	minstack := Constructor()
	minstack.Push(5)
	minstack.Push(3)
	minstack.Push(4)
	fmt.Println(minstack.Min())
	minstack.Push(1)
	fmt.Println(minstack.Min())
	minstack.Pop()
	fmt.Println(minstack.Min())
}

// https://leetcode.cn/problems/min-stack
// 最小栈
package main

import "math"

type MinStack struct {
	valStack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		[]int{},
		[]int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	this.valStack = append(this.valStack, val)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(val, top))
}

func (this *MinStack) Pop() {
	this.valStack = this.valStack[:len(this.valStack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.valStack[len(this.valStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

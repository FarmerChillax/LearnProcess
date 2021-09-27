package day1

import (
	"math"
)

type MinStack struct {
	stack    []int
	secStack []int
	min      int
}

/** initialize your data structure here. */
func StackConstructor() MinStack {
	return MinStack{min: math.MaxInt64}
}

func (this *MinStack) Push(x int) {
	if x <= this.min {
		this.min = x
		this.secStack = append(this.secStack, x)
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	if this.stack[len(this.stack)-1] == this.min {
		this.secStack = this.secStack[:len(this.secStack)-1]
		if len(this.secStack) == 0 {
			this.min = math.MaxInt64
		} else {
			this.min = this.secStack[len(this.secStack)-1]
		}
	}

	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

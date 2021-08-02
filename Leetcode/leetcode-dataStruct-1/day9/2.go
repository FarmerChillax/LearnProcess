package main

import "fmt"

func main() {
	queue := Constructor()
	queue.Push(233)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Empty())
}

type MyQueue struct {
	in, out []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) in2out() {
	for len(this.in) > 0 {
		this.out = append(this.out, this.in[len(this.in)-1])
		this.in = this.in[:len(this.in)-1]
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		// 将输入栈 弹出到输出栈中
		this.in2out()
	}
	out := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return out
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		// 将输入栈 弹出到输出栈中
		this.in2out()
	}
	return this.out[len(this.out)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

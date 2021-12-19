package day1

import "fmt"

// type CQueue struct {
// 	input, output []int
// }

// func Constructor() CQueue {
// 	return CQueue{}
// }

// func (this *CQueue) in2out() {
// 	for len(this.input) > 0 {
// 		this.output = append(this.output, this.input[len(this.input)-1])
// 		this.input = this.input[:len(this.input)-1]
// 	}
// }

// func (this *CQueue) AppendTail(value int) {
// 	this.input = append(this.input, value)
// }

// func (this *CQueue) DeleteHead() int {
// 	if len(this.input) == 0 && len(this.output) == 0 {
// 		return -1
// 	}
// 	if len(this.output) == 0 {
// 		this.in2out()
// 	}
// 	out := this.output[len(this.output)-1]
// 	this.output = this.output[:len(this.output)-1]
// 	return out
// }

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

// //  version 2
type CQueue struct {
	val      [2][]int
	nowQueue int
}

func (this *CQueue) TestReversal() {
	fmt.Println((this.nowQueue + 1) % len(this.val))
}

// 改变指针，交换输入输出栈
func (this *CQueue) IndexReversal() {
	this.nowQueue = (this.nowQueue + 1) % len(this.val)
}

// 获取output的index
func (this *CQueue) GetOutputIndex() int {
	if this.nowQueue == 0 {
		return 1
	}
	return 0
}

func Constructor() CQueue {
	return CQueue{
		val:      [2][]int{[]int{}, []int{}},
		nowQueue: 0,
	}
}

func (this *CQueue) AppendTail(value int) {
	this.val[this.nowQueue] = append(this.val[this.nowQueue], value)
}

func (this *CQueue) DeleteHead() int {
	outputIndex := this.GetOutputIndex()
	if len(this.val[outputIndex]) == 0 {
		this.IndexReversal()
		outputIndex = this.GetOutputIndex()
		if len(this.val[outputIndex]) == 0 {
			return -1
		}
	}
	result := this.val[outputIndex][0]
	this.val[outputIndex] = this.val[outputIndex][1:]
	return result
}

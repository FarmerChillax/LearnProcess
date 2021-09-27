package day1

type CQueue struct {
	input, output []int
}

func QueueConstructor() CQueue {
	return CQueue{}
}

func (this *CQueue) in2out() {
	for len(this.input) > 0 {
		this.output = append(this.output, this.input[len(this.input)-1])
		this.input = this.input[:len(this.input)-1]
	}
}

func (this *CQueue) AppendTail(value int) {
	this.input = append(this.input, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.input) == 0 && len(this.output) == 0 {
		return -1
	}
	if len(this.output) == 0 {
		this.in2out()
	}
	out := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]
	return out
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

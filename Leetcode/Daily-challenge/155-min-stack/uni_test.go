package minstack

import (
	"fmt"
	"testing"
)

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(123)
	obj.Push(233)
	param_3 := obj.Top() // 233
	fmt.Println(param_3)
	obj.Pop()
	obj.Pop()
	param_4 := obj.GetMin()
	fmt.Println(obj, param_4)
}

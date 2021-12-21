package test

import (
	"offer/day1"
	"testing"
)

func TestSub1(t *testing.T) {
	expect := 3
	queue := day1.Constructor()
	queue.AppendTail(expect)
	res := queue.DeleteHead()
	if res != expect {
		t.Errorf("return value is: %d\n", res)
	}
}

func TestSub2(t *testing.T) {
	stack := day1.StackConstructor()
	stack.Push(0)
	stack.Push(1)
	stack.Push(0)
	res := stack.Min()
	if res != 0 {
		t.Errorf("res value: %d\n", res)
	}
	stack.Pop()

	res = stack.Min()
	if res != 0 {
		t.Errorf("res value: %d\n", res)
	}
}

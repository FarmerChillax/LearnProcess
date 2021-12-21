package day1

import (
	"log"
	"testing"
)

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func TestStackQueue(t *testing.T) {
	queue := Constructor()
	// 1
	queue.AppendTail(233)
	result := queue.DeleteHead()
	expected := 233
	handleErr(result, expected)
	// 2
	result = queue.DeleteHead()
	expected = -1
	handleErr(result, expected)
	// 3
	queue.AppendTail(123)
	queue.AppendTail(233)
	queue.AppendTail(666)
	result = queue.DeleteHead()
	expected = 123
	handleErr(result, expected)
}

func handleErr(result, expected interface{}) {
	if result != expected {
		log.Fatalf("not match, result: %v; expected: %v\n", result, expected)
	}
}

package test

import (
	"fmt"
	"offer/day2"
	"reflect"
	"testing"
)

func makeList(lists []int) *day2.ListNode {
	head := &day2.ListNode{}
	node := head
	for _, item := range lists {
		node.Next = &day2.ListNode{Val: item, Next: nil}
		node = node.Next
	}
	return head.Next
}

func makeNodes(lists [][]int) *day2.Node {
	head := &day2.Node{}
	node := head
	for _, tNode := range lists {
		for _, item := range tNode {
			node.Next = &day2.Node{Val: item, Next: nil}
			node = node.Next
		}
	}
	return head.Next
}

func TestDay2Sub1(t *testing.T) {
	lists := []int{1, 3, 2}
	expected := []int{2, 3, 1}
	head := makeList(lists)
	res := day2.ReversePrint(head)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Not equal, res: %v\n", res)
	}
}

func TestDay2Sub1V2(t *testing.T) {
	lists := []int{1, 3, 2}
	expected := []int{2, 3, 1}
	head := makeList(lists)
	res := day2.ReversePrintV2(head)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Not equal, res: %v\n", res)
	}
}
func TestDay2Sub1V3(t *testing.T) {
	lists := []int{1, 3, 2}
	expected := []int{2, 3, 1}
	head := makeList(lists)
	res := day2.ReversePrintV3(head)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Not equal, res: %v\n", res)
	}
}

func TestDay2Sub2(t *testing.T) {
	// todoing
	lists := []int{1, 2, 3, 4, 5}
	fmt.Println(lists)
}

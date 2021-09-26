package sort_test

import (
	"algorithrn/sorts"
	"reflect"
	"sort"
	"testing"
)

func testQuickSort(t *testing.T) {
	arr := []int{12, 87, 1, 66, 38, 213, 126, 1233, 6542, 456, 973, 102, 3}
	var expected []int
	copy(expected, arr)
	result := sorts.QuickSort(arr)
	sort.Ints(expected)
	if reflect.DeepEqual(result, expected) {
		t.Errorf("Sorts not equal\n")
		t.Errorf("result: %v\n", result)
		t.Errorf("expect: %v\n", expected)
	}
}

func testSelectSort(t *testing.T) {
	arr := []int{12, 87, 1, 66, 38, 213, 126, 1233, 6542, 456, 973, 102, 3}
	var expected []int
	copy(expected, arr)
	result := sorts.SelectionSort(arr)
	sort.Ints(expected)
	if reflect.DeepEqual(result, expected) {
		t.Errorf("Sorts not equal\n")
		t.Errorf("result: %v\n", result)
		t.Errorf("expect: %v\n", expected)
	}
}

func testBubbleSort(t *testing.T) {
	arr := []int{12, 87, 1, 66, 38, 213, 126, 1233, 6542, 456, 973, 102, 3}
	var expected []int
	copy(expected, arr)
	sorts.BubbleSort(arr)
	sort.Ints(expected)
	if reflect.DeepEqual(arr, expected) {
		t.Errorf("Sorts not equal\n")
		t.Errorf("result: %v\n", arr)
		t.Errorf("expect: %v\n", expected)
	}
}

func testOptimizationBubbleSort(t *testing.T) {
	arr := []int{12, 87, 1, 66, 38, 213, 126, 1233, 6542, 456, 973, 102, 3}
	var expected []int
	copy(expected, arr)
	sorts.OptimizationBubbleSort(arr)
	sort.Ints(expected)
	if reflect.DeepEqual(arr, expected) {
		t.Errorf("Sorts not equal\n")
		t.Errorf("result: %v\n", arr)
		t.Errorf("expect: %v\n", expected)
	}
}
func TestSub(t *testing.T) {
	t.Run("type=Quick", testQuickSort)
	t.Run("type=select", testSelectSort)
	t.Run("type=Bubble", testBubbleSort)
	t.Run("type=OptBubble", testOptimizationBubbleSort)
}

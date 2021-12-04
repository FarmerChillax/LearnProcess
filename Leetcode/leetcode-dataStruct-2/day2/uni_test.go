package day2

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	intervals := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	expected := [][]int{
		[]int{1, 6},
		[]int{8, 10},
		[]int{15, 18},
	}

	result := merge(intervals)
	if !reflect.DeepEqual(result, expected) {
		log.Fatalln("not match, result:", result)
	}
}

func TestMerge_2(t *testing.T) {
	intervals := [][]int{
		[]int{1, 3},
	}
	expected := [][]int{
		[]int{1, 3},
	}

	result := merge(intervals)
	if !reflect.DeepEqual(result, expected) {
		log.Fatalln("not match, result:", result)
	}
}

func TestMerge_3(t *testing.T) {
	intervals := [][]int{
		[]int{1, 4},
		[]int{4, 5},
	}
	expected := [][]int{
		[]int{1, 5},
	}

	result := merge(intervals)
	if !reflect.DeepEqual(result, expected) {
		log.Fatalln("not match, result:", result)
	}
}

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func TestConstructor(t *testing.T) {
	myHashMap := Constructor()
	myHashMap.Put(1, 2)
	fmt.Println(myHashMap)
	myHashMap.Put(1, 3)
	fmt.Println(myHashMap)
	fmt.Println(myHashMap.Get(1))
	myHashMap.Remove(1)
	fmt.Println(myHashMap)
	fmt.Println(myHashMap.Get(1))
}

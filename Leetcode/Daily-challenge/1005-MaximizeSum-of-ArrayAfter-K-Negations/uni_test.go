package maximizesumofarrayafterknegations

import (
	"log"
	"reflect"
	"testing"
)

func Test1005_1(t *testing.T) {
	nums := []int{4, 2, 3}
	k := 1
	expected := 5
	result := largestSumAfterKNegations(nums, k)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("not match. result:", result)
	}
}

func Test1005_2(t *testing.T) {
	nums := []int{3, -1, 0, 2}
	k := 3
	expected := 6
	result := largestSumAfterKNegations(nums, k)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("not match. result:", result)
	}
}

func Test1005_3(t *testing.T) {
	nums := []int{2, -3, -1, 5, -4}
	k := 2
	expected := 13
	result := largestSumAfterKNegations(nums, k)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("not match. result:", result)
	}
}

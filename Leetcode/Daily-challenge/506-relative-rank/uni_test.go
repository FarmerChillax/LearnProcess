package relativerank

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestMaxPower(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	expected := []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}
	result := findRelativeRanks(input)
	fmt.Println("result:", result)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("Not Match,", result)
	}
}

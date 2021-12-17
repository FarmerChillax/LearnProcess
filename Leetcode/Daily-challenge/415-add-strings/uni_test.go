package addstrings

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestAddStrings(t *testing.T) {
	num1 := "11"
	num2 := "123"
	excepted := "134"
	result := addStrings(num1, num2)
	fmt.Println(result)
	if !reflect.DeepEqual(result, excepted) {
		log.Fatalln("not match, result:", result)
	}
}

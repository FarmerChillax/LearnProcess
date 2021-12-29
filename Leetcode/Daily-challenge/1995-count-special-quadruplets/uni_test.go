package countspecialquadruplets

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestCountQuadruplets(t *testing.T) {
	input := "tourist"
	expected := 1
	result := countQuadruplets()
	fmt.Println("result:", result)
	if !reflect.DeepEqual(expected, result) {
		log.Fatalln("Not Match,", result)
	}
}

package utils

import (
	"log"
	"reflect"
)

func HandleResult(res, expcet interface{}) {
	if !reflect.DeepEqual(res, expcet) {
		log.Fatalf("expcet: %d; result: %d\n", expcet, res)
	}
}

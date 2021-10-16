package plugin

import (
	"reflect"
	"strings"
)

var filterNames = make(map[string]reflect.Type)

type UpperFilter struct{}

func (u *UpperFilter) Process(msg string) string {
	return strings.ToUpper(msg)
}

func init() {
	filterNames["upper"] = reflect.TypeOf(UpperFilter{})
}

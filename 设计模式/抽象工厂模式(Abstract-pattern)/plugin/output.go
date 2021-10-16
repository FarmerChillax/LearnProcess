package plugin

import (
	"fmt"
	"reflect"
)

var outputNames = make(map[string]reflect.Type)

type ConsoleOutput struct{}

func (c *ConsoleOutput) Send(msg string) {
	fmt.Println(msg)
}

func init() {
	outputNames["console"] = reflect.TypeOf(ConsoleOutput{})
}

package plugin

import "reflect"

type InputType string

// input插件名称和类型的映射关系
var inputNames = make(map[string]reflect.Type)

// Hello input 插件，接收消息
type HelloInput struct{}

func (h *HelloInput) Receive() string {
	return "Hello World"
}

func init() {
	inputNames["hello"] = reflect.TypeOf(HelloInput{})
}

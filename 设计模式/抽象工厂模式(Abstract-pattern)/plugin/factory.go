package plugin

import "reflect"

type Type uint8

// 插件接口
type Plugin interface{}

// 输入插件
type Input interface {
	Plugin
	Receive() string
}

// 过滤插件
type Filter interface {
	Plugin
	Process(msg string) string
}

// 输出插件
type Output interface {
	Plugin
	Send(msg string)
}

type Config struct {
	Name string
}

// 插件抽象工厂接口
type Factory interface {
	Create(conf Config) Plugin
}

// input插件工厂对象，实现Factory接口
type InputFactory struct{}

// 读取配置，通过反射机制进行对象实例化
func (i *InputFactory) Create(conf Config) Plugin {
	t, _ := inputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

// filter和output插件工厂实现类似
type FilterFactory struct{}

func (f *FilterFactory) Create(conf Config) Plugin {
	t, _ := filterNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type OutputFactory struct{}

func (o *OutputFactory) Create(conf Config) Plugin {
	t, _ := outputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

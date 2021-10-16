package pipe

import plugin "abstract/plugin"

// 保存用于创建Plugin的工厂实例，其中map的key为插件类型，value为抽象工厂接口
var pluginFactories = make(map[plugin.Type]plugin.Factory)

// 根据plugin.Type返回对应Plugin类型的工厂实例
func factoryOf(t plugin.Type) plugin.Factory {
	factory, _ := pluginFactories[t]
	return factory
}

// pipeline工厂方法，根据配置创建一个Pipeline实例
// func Of(conf Config) *Pipeline {
// 	p := &Pipeline{}
// 	p.input = factoryOf(plugin.InputType).Create(conf.Input).(plugin.Input)
// 	p.filter = factoryOf(plugin.FilterType).Create(conf.Filter).(plugin.Filter)
// 	p.output = factoryOf(plugin.OutputType).Create(conf.Output).(plugin.Output)
// 	return p
// }

// // 初始化插件工厂对象
// func init() {
// 	pluginFactories[plugin.InputType] = &plugin.InputFactory{}
// 	pluginFactories[plugin.FilterType] = &plugin.FilterFactory{}
// 	pluginFactories[plugin.OutputType] = &plugin.OutputFactory{}
// }

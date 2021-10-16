package pipe

import "abstract/plugin"

type Pipeline struct {
	input  plugin.Input
	filter plugin.Filter
	output plugin.Output
}

// 消息处理流程
func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}

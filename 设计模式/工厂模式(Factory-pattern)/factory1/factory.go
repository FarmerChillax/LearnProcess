package factory1

type Type uint8

const (
	Start Type = iota
	End
)

// 事件接口
type Event interface {
	EventType() Type
	Content() string
}

type StartEvent struct {
	content string
}

func (e StartEvent) EventType() Type {
	return Start
}

func (e StartEvent) Content() string {
	return e.content
}

type EndEvent struct {
	content string
}

func (e EndEvent) EventType() Type {
	return End
}

func (e EndEvent) Content() string {
	return e.content
}

// 工厂对象
type Factory struct{}

func (e *Factory) Create(etype Type) Event {
	switch etype {
	case Start:
		return &StartEvent{
			content: "start event",
		}
	case End:
		return &EndEvent{
			content: "End event",
		}
	default:
		return nil
	}
}

package factory1

import "testing"

func TestFactory(t *testing.T) {
	factory := Factory{}
	e := factory.Create(Start)
	if e.EventType() != Start {
		t.Errorf("expect event.Start, but actual %v.", e.EventType())
	}
	e = factory.Create(End)
	if e.EventType() != End {
		t.Errorf("expect event.End, but actual %v.", e.EventType())
	}
}

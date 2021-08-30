package factory2

import "testing"

func TestEvent(t *testing.T) {
	e := OfStart()
	if e.EventType() != Start {
		t.Errorf("expect event.Start, but actual %v.", e.EventType())
	}
	e = OfEnd()
	if e.EventType() != End {
		t.Errorf("expect event.End, but actual %v.", e.EventType())
	}
}

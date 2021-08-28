package lazymode_test

import (
	"lazy-mode/msgpool"
	"testing"
)

func TestMessagePool(t *testing.T) {
	msg := msgpool.Instance().GetMsg()
	if msg.Count != 0 {
		t.Errorf("expect msg count %d, but actual %d", 0, msg.Count)
	}
	msg.Count = 1
	msgpool.Instance().AddMsg(msg)
	msg2 := msgpool.Instance().GetMsg()
	if msg.Count != 1 {
		t.Errorf("expect msg count %d, but actual %d.", 1, msg2.Count)
	}
}

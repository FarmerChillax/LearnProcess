package sub_test

import (
	"daily-challenge/sub"
	"testing"
)

func Subject211(t *testing.T) {
	wordDiction := sub.Constructor()

	wordDiction.AddWord("bad")
	wordDiction.AddWord("dad")
	wordDiction.AddWord("mad")

	res := wordDiction.Search("pad")
	if res != false {
		t.Errorf("return: %v\n", res)
	}
	res = wordDiction.Search("bad")
	if res != true {
		t.Errorf("return: %v\n", res)
	}
	wordDiction.Search(".ad")
	if res != true {
		t.Errorf("return: %v\n", res)
	}
	wordDiction.Search("b..")
	if res != true {
		t.Errorf("return: %v\n", res)
	}
}

func TestSub211(t *testing.T) {
	t.Run("sub=211", Subject211)
}

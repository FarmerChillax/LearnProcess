package allone

import (
	"log"
	"testing"
)

func TestAllone(t *testing.T) {
	obj := Constructor()
	obj.Inc("a")
	obj.Inc("b")
	obj.Inc("b")
	obj.Inc("c")
	obj.Inc("c")
	obj.Inc("c")

	obj.Dec("b")
	obj.Dec("b")
	log.Println(obj)
	log.Println(obj.GetMinKey())
	obj.Dec("a")
	log.Println(obj)
	log.Println(obj.GetMaxKey())
	log.Println(obj.GetMinKey())

}

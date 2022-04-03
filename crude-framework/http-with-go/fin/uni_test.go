package fin

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	fmt.Println(parsePattern("/p/*"))
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	fmt.Println(parsePattern("/p/*name/*"))
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	fmt.Println(r.roots["GET"].children[0].children[0])
	fmt.Println("*********")
	fmt.Println(r.handlers)
	n, ps := r.getRoute("GET", "/hello/farmer")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "farmer" {
		t.Fatal("name should be equal to 'farmer'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

}

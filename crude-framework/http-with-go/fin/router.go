package fin

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// 解析路由，将路径切割
func parsePattern(pattern string) []string {
	parts := make([]string, 0)
	for _, item := range strings.Split(pattern, "/") {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
	// log.Printf("Route %5s - %s", method, pattern)
}

func (r *router) getRoute(method, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		// 如果该方法没有挂载到root, 则返回nil
		return nil, nil
	}

	// 开始搜索路由
	node := root.search(searchParts, 0)

	if node != nil {
		parts := parsePattern(node.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				// 如果这一级路由是 : 开头，则表示为路由参数
				// params[注册的变量名] = 变量内容
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				// 如果是*开头，则表明为全匹配(后面都为变量内容)
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return node, params
	}
	return nil, nil

}

func (r *router) handle(c *Context) {

	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		key := c.Method + "-" + node.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
		// r.handlers[key](c)
	} else {
		c.handlers = append(c.handlers, func(ctx *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}

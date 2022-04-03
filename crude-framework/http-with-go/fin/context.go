package fin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// 原生对象
	w http.ResponseWriter
	r *http.Request
	// 请求信息
	Path, Method string
	// 响应信息
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w:      w,
		r:      r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// 从Form表单读取数据
func (c *Context) PostForm(key string) string {
	return c.r.FormValue(key)
}

// 从Query String 读取数据
func (c *Context) Query(key string) string {
	return c.r.URL.Query().Get(key)
}

// set http response status code
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.w.WriteHeader(code)
}

// set http response header with k v
func (c *Context) SetHeader(k, v string) {
	c.w.Header().Set(k, v)
}

// 返回字符串
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.w.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, data interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.w)
	if err := encoder.Encode(data); err != nil {
		http.Error(c.w, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.w.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.w.Write([]byte(html))
}

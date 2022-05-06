package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.farmer233.top",
	"https://blog.farmer233.top",
	"https://blog.xiaotao233.top",
}

func main() {
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			// 当前 goroutine 结束后给 wg 计数减1, wg.Done() 等价于 wg.Add(-1)
			// defer wg.Add(-1)
			defer wg.Done()

			// 发生 HTTP get 请求并打印 HTTP 状态码
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(url, resp.Status)
			} else {
				fmt.Println(err)
			}

		}(url)
	}

	wg.Wait()
}

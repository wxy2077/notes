/*
* @Time    : 2021-02-01 20:20
* @Author  : CoderCharm
* @File    : gee.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	// 存储请求方法-匹配参数 和 处理请求的函数
	router map[string]HandlerFunc
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 把请求路径 和 处理函数放入 router 这个map中
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	// 通过方法和请求路径 匹配到 handler请求处理函数
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		// 没有匹配到则表示 路径和请求方法不存在
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

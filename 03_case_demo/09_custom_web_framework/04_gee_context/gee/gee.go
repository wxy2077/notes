/*
* @Time    : 2021-02-01 20:35
* @Author  : CoderCharm
* @File    : gee.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package gee

import "net/http"

type HandlerFunc func(c *Context)

type Engine struct {
	router *Route
}

// Engine 必须实现 ServeHTTP
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	e.router.handle(c)
}

func New() *Engine {
	return &Engine{router: NewRoute()}
}

// 给router 套一层壳添加 请求方式和请求处理函数
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

// 添加 GET 方法的请求处理函数
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// 添加POST 方法的请求处理函数
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// 启动web服务
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

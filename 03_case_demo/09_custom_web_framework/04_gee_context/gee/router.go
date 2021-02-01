/*
* @Time    : 2021-02-01 20:36
* @Author  : CoderCharm
* @File    : router.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package gee

import "net/http"

type Route struct {
	handlers map[string]HandlerFunc
}

// 初始化路由
func NewRoute() *Route {
	return &Route{
		handlers: make(map[string]HandlerFunc),
	}
}

// 路由内部添加添加方法
func (r *Route) addRoute(method string, pattern string, handler HandlerFunc) {
	// 拼接请求方式 和 路由
	key := method + "-" + pattern
	// 存储到路由处理的映射中 和 请求处理函数 一一对应
	r.handlers[key] = handler
}

// 找到并执行处理请求函数
func (r *Route) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		// 没有找到直接 404
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

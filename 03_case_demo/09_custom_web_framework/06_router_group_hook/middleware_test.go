/*
* @Time    : 2021-02-03 19:55
* @Author  : CoderCharm
* @File    : middleware_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package main

import (
	"fmt"
	"testing"
)

type HandlerFunc func(c *Context)

// HandlersChain defines a HandlerFunc array.
type HandlersChain []HandlerFunc

// 管理中间件的上下文
type Context struct {
	// middleware
	handlers HandlersChain //存储顺序 [中间件1, 中间件2, 中间件3...., 请求处理函数]
	index    int
}

func newContext() *Context {
	return &Context{
		// gin 里面初始化变量
		// https://github.com/gin-gonic/gin/blob/master/context.go#L91
		index: -1,
	}
}

// 下一步调用
func (c *Context) Next() {
	// 如果把 Next() c.index++ 去掉，就会一直卡在第一个中间件上了，必然死循环。
	// 每调用一次 Next()， c.index 必须得 +1，否则 c.index 就会一直是 0
	c.index++            // 0
	s := len(c.handlers) // 3
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

// A 中间件
func middlewareA(c *Context) {
	fmt.Println("A--1")
	c.Next()
	fmt.Println("A--2")
}

// B 中间件
func middlewareB(c *Context) {
	fmt.Println("B-------1")
	c.Next()
	fmt.Println("B-------2")
}

// 请求处理到函数
func handleFunc(c *Context) {
	fmt.Println("请求处理函数")
}

func TestMiddleware(t *testing.T) {

	n := newContext()

	n.handlers = append(n.handlers, middlewareA)
	n.handlers = append(n.handlers, middlewareB)
	n.handlers = append(n.handlers, handleFunc)

	// 启动
	n.Next()

	// A--1
	// B-------1
	// 请求处理函数
	// B-------2
	// A--2

	// 最终执行顺序就是: A1-> B1 -> 处理函数 -> B2 -> A1
}

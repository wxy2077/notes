/*
* @Time    : 2021-02-03 19:46
* @Author  : CoderCharm
* @File    : router_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

简单 写一个 路由分组的例子 来测试一下
**/
package main

import (
	"fmt"
	"testing"
)

// 路由分组
// https://github.com/gin-gonic/gin/blob/master/routergroup.go#L41
type RouterGroup struct {
	prefix string       // 前缀
	engine *Engine      // 内部的Engine始终保证为共享的一个
	parent *RouterGroup // 父路由
}

// 核心
type Engine struct {
	*RouterGroup // 使Engine 也拥有RouterGroup的方法
}

// Get方法 这里没有实现响应处理函数
func (group *RouterGroup) Get(pattern string) {
	fmt.Printf("Get %s%s\n", group.prefix, pattern)
}

// 创建路由分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	return &RouterGroup{
		prefix: group.prefix + prefix, // 前缀为上一个 路由分组前缀 加下一个
		parent: group,                 // 当前路由设置为父路由
	}
}

func New() *Engine {
	engine := &Engine{}
	// RouterGroup里面的 engine属性为 自身的engine  确保所有的engine 为一个
	engine.RouterGroup = &RouterGroup{engine: engine}
	return engine
}

func TestRouter(t *testing.T) {
	// 创建实例
	r := New()

	r.Get("/index")

	// 路由分组
	api := r.Group("/api")
	api.Get("/123")

	// api子路由分组
	v1 := api.Group("/v1")
	v1.Get("/666")

	// 输出
	// Get /index
	// Get /api/123
	// Get /api/v1/666
}

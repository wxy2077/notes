/*
* @Time    : 2021-02-02 19:05
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

这一章 路由知识 没有细究

性能肯定是不如 04_gee_context 因为动态路由匹配这里，不如上一章 map[key] 匹配效率高。

# 路由匹配的第三方库
https://github.com/julienschmidt/httprouter

django 1.x 版本是用正则匹配路由，后面2.x+版本就改了

路由参数这一块 实现起来感觉有点绕，而且也不常用路径参数，一般用Query参数居多

核心就是 insert 和 search 对路由操作的方法。

// 路由结构体改造 两个属性
type router struct {
	Roots    map[string]*node  # 存储节点信息
	Handlers map[string]HandlerFunc	 # 存储所有的路由信息 和 请求处理函数
}


**/
package main

import (
	"log"
	"net/http"
	"practice/03_case_demo/09_custom_web_framework/05_gee_dynamic_route/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=大家好
		c.String(http.StatusOK, "gee %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/你好
		c.String(http.StatusOK, "gee %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.POST("/hello", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"hello": "I'am fine!",
		})
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	// 可以不暴露出 Router 该成小写
	log.Println(r.Router.Roots)
	log.Println(r.Router.Handlers)

	_ = r.Run("127.0.0.1:7054")
}

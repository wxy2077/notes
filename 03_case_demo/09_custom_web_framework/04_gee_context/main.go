/*
* @Time    : 2021-02-01 20:43
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package main

import (
	"practice/03_case_demo/09_custom_web_framework/04_gee_context/gee"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.String(200, "测试首页 你输入的name为%s 路径为 %s", c.Query("name"), c.Path)
	})

	r.GET("/json", func(c *gee.Context) {
		c.JSON(200, gee.H{
			"json": "Value",
		})
	})

	_ = r.Run("127.0.0.1:7053")

}

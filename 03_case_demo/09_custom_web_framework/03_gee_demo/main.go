/*
* @Time    : 2021-02-01 20:25
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

go 不推荐相对导入, 这里直接用的绝对导入

**/
package main

import (
	"fmt"
	"net/http"
	"practice/03_case_demo/09_custom_web_framework/03_gee_demo/gee"
)

func main() {

	r := gee.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	_ = r.Run("127.0.0.1:7052")
}

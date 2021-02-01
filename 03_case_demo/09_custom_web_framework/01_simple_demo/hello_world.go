/*
* @Time    : 2021-02-01 19:46
* @Author  : CoderCharm
* @File    : hello_world.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
一个最基础的返回JSON 数据的web服务器

https://stackoverflow.com/questions/31622052/how-to-serve-up-a-json-response-using-go


connCtx := ctx
if cc := srv.ConnContext; cc != nil {
	connCtx = cc(connCtx, rw)
	if connCtx == nil {
		panic("ConnContext returned nil")
	}
}
tempDelay = 0
c := srv.newConn(rw)
c.setState(c.rwc, StateNew) // before Serve can return
go c.serve(connCtx) // 每新进来一个连接就开一个协程处理

**/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// 获取url参数
	fmt.Println(r.URL.Query().Get("aaa"))

	// 设置编码 一定得要设置字符集编码 否则中文乱码
	w.Header().Set("Content-Type", "text/html;charset=utf-8")

	_, _ = w.Write([]byte("<h2 style='color:blue'>你好！首页</h2>"))

}

type DemoData struct {
	Hello string `json:"Hello"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// 返回json数据
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := DemoData{"World"}
	_ = json.NewEncoder(w).Encode(data)

}

func main() {

	// 设置两个路由
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloWorld)

	addr := "127.0.0.1:7050"
	fmt.Println("启动地址为:", addr)
	server := http.Server{
		Addr: addr,
	}
	// 监听服务
	_ = server.ListenAndServe()

}

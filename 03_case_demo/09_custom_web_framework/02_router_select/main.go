/*
* @Time    : 2021-02-01 19:50
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :


package http

// A Handler responds to an HTTP request.
//
// ServeHTTP should write reply headers and data to the ResponseWriter
// and then return. Returning signals that the request is finished; it
// is not valid to use the ResponseWriter or read from the
// Request.Body after or concurrently with the completion of the
// ServeHTTP call.
//
// Depending on the HTTP client software, HTTP protocol version, and
// any intermediaries between the client and the Go server, it may not
// be possible to read from the Request.Body after writing to the
// ResponseWriter. Cautious handlers should read the Request.Body
// first, and then reply.
//
// Except for reading the body, handlers should not modify the
// provided Request.
//
// If ServeHTTP panics, the server (the caller of ServeHTTP) assumes
// that the effect of the panic was isolated to the active request.
// It recovers the panic, logs a stack trace to the server error log,
// and either closes the network connection or sends an HTTP/2
// RST_STREAM, depending on the HTTP protocol. To abort a handler so
// the client sees an interrupted response but the server doesn't log
// an error, panic with the value ErrAbortHandler.
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

func ListenAndServe(address string, h Handler) error

**/
package main

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router router
}

// 定义请求函数类型
type handler func(w http.ResponseWriter, r *http.Request)

// 定义路由 - 对应 处理请求函数
type router map[string]handler

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 只通过路径 匹配处理请求的方法 不区分 GET or POST
	// r.URL.Path
	if handler, ok := engine.router[r.URL.Path]; ok {
		handler(w, r)
	} else {
		_, _ = fmt.Fprintf(w, "404")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World")
}

func main() {

	r := router{}

	r["/"] = func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "首页")
	}

	r["/hello"] = hello

	//engine := new(Engine)
	engine := &Engine{
		router: r,
	}

	addr := "127.0.0.1:7051"
	fmt.Println("服务启动:", addr)

	_ = http.ListenAndServe(addr, engine)
}

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

以下为Golang http 标准库核心服务

func (srv *Server) Serve(l net.Listener) error {
	if fn := testHookServerServe; fn != nil {
		fn(srv, l) // call hook with unwrapped listener
	}

	origListener := l
	l = &onceCloseListener{Listener: l}
	defer l.Close()

	if err := srv.setupHTTP2_Serve(); err != nil {
		return err
	}

	if !srv.trackListener(&l, true) {
		return ErrServerClosed
	}
	defer srv.trackListener(&l, false)

	baseCtx := context.Background()
	if srv.BaseContext != nil {
		baseCtx = srv.BaseContext(origListener)
		if baseCtx == nil {
			panic("BaseContext returned a nil context")
		}
	}

	var tempDelay time.Duration // how long to sleep on accept failure

	ctx := context.WithValue(baseCtx, ServerContextKey, srv)
	// 最外层死循环
	for {
		// 监听listener的请求
		rw, err := l.Accept()
		if err != nil {
			select {
			case <-srv.getDoneChan():
				return ErrServerClosed
			default:
			}
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				srv.logf("http: Accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return err
		}
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
		go c.serve(connCtx)  // 每新进来一个连接，就开一个 goroutine 处理
	}
}

每个 goroutine 最小只需要2K的内存
https://github.com/golang/go/blob/bbd25d26c0a86660fb3968137f16e74837b7a9c6/src/runtime/stack.go#L72

如果当前 HTTP 服务接收到了海量的请求，会在内部创建大量的 Goroutine，这可能会使整个服务质量明显降低无法处理请求

比 标准库net/http 快10倍的第三方http库
https://github.com/valyala/fasthttp

测试设备以及结果
https://github.com/valyala/fasthttp/issues/4

为什么fasthttp 比 net/http 快10倍
https://stackoverflow.com/questions/41627931/why-is-fasthttp-faster-than-net-http


基于 fasthttp的 web框架 https://github.com/gofiber/fiber

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

type H map[string]interface{}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// 返回json数据
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//data := DemoData{"World"}
	//_ = json.NewEncoder(w).Encode(data)

	encoder := json.NewEncoder(w)
	obj := H{"Hello": "World"}
	if err := encoder.Encode(obj); err != nil {
		http.Error(w, err.Error(), 500)
	}
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

/*
* @Time    : 2021-02-01 20:33
* @Author  : CoderCharm
* @File    : context.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// 存储请求上下文信息
type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter

	// 把一些基础信息单独抽出来
	Path       string
	Method     string
	StatusCode int
}

// 构建上下文实例
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Req:    r,
		Writer: w,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// 获取url的查询参数
func (c *Context) Query(name string) string {
	return c.Req.URL.Query().Get(name)
}

// 获取表单参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// 设置状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 设置header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 快速构建响应
// 返回字符串
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain;charset=utf-8")
	c.Status(code)
	_, _ = c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// 返回json数据
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json;charset=utf-8")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// 返回字节流数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, _ = c.Writer.Write(data)
}

// 返回html数据
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html;charset=utf-8")
	c.Status(code)
	_, _ = c.Writer.Write([]byte(html))
}

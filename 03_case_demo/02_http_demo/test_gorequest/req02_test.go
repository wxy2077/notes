/*
* @Time    : 2020-12-31 16:32
* @Author  : CoderCharm
* @File    : req02_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

使用gorequest clone 模式 重用通用设置
**/
package test_gorequest

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"testing"
	"time"
)

var baseRequest = gorequest.New().Timeout(5*time.Second).Set("User-Agent", UserAgent)

func TestReq2(t *testing.T) {

	getUrl := "https://httpbin.org/get"

	resp, body, errs := baseRequest.Clone().Get(getUrl).End()
	if errs != nil {
		panic(fmt.Sprintf("URL:%s 请求错误: %s", getUrl, errs))
	}
	fmt.Println(resp)
	// 可以使用 _ 接收 resp
	fmt.Printf("url:%s 请求状态 %s \n", getUrl, resp.Status)

	fmt.Println(body)
}

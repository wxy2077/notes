/*
* @Time    : 2020-12-31 17:31
* @Author  : CoderCharm
* @File    : req03_pool_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

创建一个请求对象池, 并发请求

**/
package test_gorequest

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"testing"
	"time"
)

// 定义一个请求对象池
type ReqPool struct {
	poolChan chan *gorequest.SuperAgent
}

// 获取一个请求对象
func (p *ReqPool) FetchReq(timeout time.Duration) (req *gorequest.SuperAgent, err error) {
	select {
	case req = <-p.poolChan:
		return req, nil
	case <-time.After(timeout): //超时控制
		return nil, errors.New("time out")
	}
}

// 释放请求 到请求池
func (p *ReqPool) ReleaseReq(req *gorequest.SuperAgent) error {
	select {
	case p.poolChan <- req:
		return nil
	default:
		return errors.New("request pools overflow")
	}
}

// New请求池
func NewReqPool(poolNum int) *ReqPool {
	// 实例化
	poolObj := ReqPool{}

	// 创建请求对象channel
	poolObj.poolChan = make(chan *gorequest.SuperAgent, poolNum)

	for i := 0; i < poolNum; i++ {
		baseRequest := gorequest.New().Timeout(5*time.Second).Set("User-Agent", UserAgent).
			Set("self_token", "111")
		poolObj.poolChan <- baseRequest
	}
	return &poolObj
}

func TestReqPool(t *testing.T) {

	requestPool := NewReqPool(10)

	req, err := requestPool.FetchReq(time.Second * 5)
	if err != nil {
		panic("获取请求对象错误")
	}

	defer func(requestPool *ReqPool, req *gorequest.SuperAgent) {
		_ = requestPool.ReleaseReq(req)
	}(requestPool, req)

	getUrl := "https://httpbin.org/get"

	resp, body, errs := req.Clone().Get(getUrl).End()

	if errs != nil {
		panic(fmt.Sprintf("URL:%s 请求错误: %s", getUrl, errs))
	}
	fmt.Println(resp)
	// 可以使用 _ 接收 resp
	fmt.Printf("url:%s 请求状态 %s \n", getUrl, resp.Status)

	fmt.Println(body)

	//fmt.Println(runtime.NumGoroutine())

}

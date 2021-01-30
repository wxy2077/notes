/*
* @Time    : 2021-01-26 19:16
* @Author  : CoderCharm
* @File    : server.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	imooc "practice/03_case_demo/05_go_grpc/go_micro/proto"
)

type NemoServer struct{}

// 实现接口
func (c *NemoServer) SayHello(ctx context.Context, req *imooc.SayRequest, res *imooc.SayResponse) error {
	res.Answer = "Just For Fun"
	return nil
}

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("nemo.server"),
	)

	// initialise flags
	service.Init()

	// 注册服务
	_ = imooc.RegisterNemoHandler(service.Server(), &NemoServer{})

	// start the service
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}

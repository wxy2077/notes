/*
* @Time    : 2021-01-26 21:32
* @Author  : CoderCharm
* @File    : client.go
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

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("nemo.client"),
	)

	service.Init()

	capImooc := imooc.NewNemoService("nemo.server", service.Client())
	//
	res, err := capImooc.SayHello(context.TODO(), &imooc.SayRequest{Message: "hhhhhhh"})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}

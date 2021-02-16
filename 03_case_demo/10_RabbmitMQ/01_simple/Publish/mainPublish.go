/*
* @Time    : 2021-02-16 20:23
* @Author  : CoderCharm
* @File    : mainPublish.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/

package main

import (
	"fmt"
	"practice/03_case_demo/10_RabbmitMQ/01_simple/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" + "wxy")
	defer rabbitmq.Destory()

	rabbitmq.PublishSimple("Hello MQ!")

	fmt.Println("发送成功！")

}

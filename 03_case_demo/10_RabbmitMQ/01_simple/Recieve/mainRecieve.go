/*
* @Time    : 2021-02-16 20:24
* @Author  : CoderCharm
* @File    : mainRecieve.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/

package main

import "practice/03_case_demo/10_RabbmitMQ/01_simple/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("wxy")
	rabbitmq.ConsumeSimple()
}

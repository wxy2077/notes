/*
* @Time    : 2020-12-29 09:47
* @Author  : CoderCharm
* @File    : chan_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

channel 基本使用

只能使用 make关键字创建

ch := make(chan type, value)


**/
package _2_channel

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {

	// 声明一个channel Buffer为3 表示可以存储 3个 int 数据
	c := make(chan int, 3)
	//c := make(chan int)

	c <- 1
	c <- 2
	c <- 8

	// 取出一个数据（可以选择不接收） 队列顺序
	<-c
	//n := <-c
	//t.Log(n)

	// 查看channel的 Buffer 容量
	t.Log(cap(c))

	c <- 9

	// 也可以多次取值
	x, y := <-c, <-c

	t.Log(fmt.Sprintf("%d + %d = %d", x, y, x+y))

}

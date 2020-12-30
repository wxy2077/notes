/*
* @Time    : 2020-12-30 21:00
* @Author  : CoderCharm
* @File    : chan5_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

channel 设置只读 和 只写
**/
package _2_channel

import (
	"fmt"
	"testing"
)

func Producer5(writeC chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("生产+++%d\n", i)
		writeC <- i
	}
}

func Producer6(writeC chan<- int) {
	for i := 5; i < 10; i++ {
		fmt.Printf("生产+++%d\n", i)
		writeC <- i
	}
}

func Consumer5(redC <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("-----------------消费 %d \n", <-redC)
	}
}

func TestChannelCommunication(t *testing.T) {

	c := make(chan int, 15)

	// 只读
	var redC <-chan int = c
	// 只写
	var writeC chan<- int = c

	// 生产
	go Producer5(writeC)
	go Producer6(writeC)

	// 消费
	Consumer5(redC)

}

/*
* @Time    : 2020-12-29 11:16
* @Author  : CoderCharm
* @File    : chan2_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

单个 channel 的使用

可以利用 Range 和 Close 多次取 和 存


**/
package _2_channel

import (
	"fmt"
	"testing"
)

func addChan(n int, c chan string) {
	for i := 0; i < n; i++ {
		// 循环添加到
		c <- fmt.Sprintf("--- %d ---", i)
	}

	// 结束添加
	close(c)
}

func TestChan2(t *testing.T) {

	// 声明一个 string chan
	c := make(chan string, 10)

	// 协程异步调用 添加到chan
	go func(n int, c chan string) {
		for i := 0; i < n; i++ {
			// 循环添加到
			c <- fmt.Sprintf("--- %d ---", i)
		}
		// 结束添加
		close(c)
	}(cap(c), c)

	// 循环取出
	for v := range c {
		t.Log(v)
	}

}

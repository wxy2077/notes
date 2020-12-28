/*
* @Time    : 2020-12-28 23:33
* @Author  : CoderCharm
* @File    : async_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

CSP 并发机制
**/
package _2_channel

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("Other Task 111111")
	time.Sleep(time.Millisecond * 50)
	fmt.Println("Task is done")
}

//func TestService(t *testing.T)  {
//	fmt.Println(service())
//	otherTask()
//}

func AsyncService() chan string {
	// 创建一个channel
	retCh := make(chan string)
	//retCh := make(chan string, 1)

	// 协程函数 异步执行
	go func() {
		ret := service()
		fmt.Println("before return result")
		// 往channel 存入值
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()

	otherTask()

	fmt.Println(<-retCh)

	//time.Sleep(1*time.Second)
}

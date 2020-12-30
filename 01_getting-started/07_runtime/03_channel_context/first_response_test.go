/*
* @Time    : 2020-12-30 15:20
* @Author  : CoderCharm
* @File    : first_response_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

并发控制

**/
package _3_channel_context

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {

	numOfRunner := 10

	// 不指定buffer会协程泄漏 函数返回后 还有其他协程存活 耗尽资源
	//ch := make(chan string)

	// numOfRunner 防止协程泄漏
	ch := make(chan string, numOfRunner)

	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
			//ch <- fmt.Sprintf("result is from %d", i)
		}(i)
	}

	// 注意 当channel里面有值的时候 就回立即返回  没有就会阻塞
	return <-ch
}

func TestFirstResponse(t *testing.T) {

	// 会输出当前系统中的协程数
	t.Log("Before:", runtime.NumGoroutine())

	// 响应每次不一致 是由于 channel的阻塞机制
	t.Log(FirstResponse())

	//
	time.Sleep(time.Second * 5)

	t.Log("After:", runtime.NumGoroutine())
}

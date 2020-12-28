/*
* @Time    : 2020-12-28 21:42
* @Author  : CoderCharm
* @File    : goroutine_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _1_goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func task(i int) {
	runtime.Gosched() // runtime.Gosched () 表示让 CPU 把时间片让给别人，下次某个时候继续恢复执行该 goroutine。
	fmt.Println(i)
}

func TestGoRoutine(t *testing.T) {

	for i := 0; i < 10; i++ {
		go task(i)
	}
	task(999)

	//time.Sleep(time.Second*1)
}

/*
* @Time    : 2021-01-13 19:13
* @Author  : CoderCharm
* @File    : demo01_goroutine_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _8_panic_recover

import (
	"testing"
	"time"
)

func TestGoroutineRecover(t *testing.T) {

	// 主线程中的 defer 没有执行
	defer println("in main")

	go func() {
		defer println("in goroutine")
		// goroutine 中的错误 会直接影响主线程
		panic("自定义错误")
	}()

	time.Sleep(1 * time.Second)
}

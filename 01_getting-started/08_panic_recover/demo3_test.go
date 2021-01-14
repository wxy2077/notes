/*
* @Time    : 2021-01-13 19:26
* @Author  : CoderCharm
* @File    : demo3_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _8_panic_recover

import (
	"errors"
	"testing"
	"time"
)

var (
	customErr = errors.New("自定义错误")
)

func TestRecover(t *testing.T) {
	// 主线程中的 defer 没有执行
	defer println("in main")

	go func() {
		// `panic` 只会触发当前 `Goroutine` 的 `defer`
		defer func() {
			if err := recover(); err != nil {
				//恢复错误
				t.Log("恢复错误")

				// 捕获自定义错误
				if err == customErr {
					t.Log(err)
				} else {
					t.Log("其他错误")
				}
			}
		}()

		defer println("in goroutine")
		// 抛出自定义错误
		panic(customErr)
	}()

	time.Sleep(1 * time.Second)
}

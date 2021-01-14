/*
* @Time    : 2021-01-13 19:23
* @Author  : CoderCharm
* @File    : demo2_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _8_panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer println("Finally!")

	defer func() {
		if err := recover(); err != nil {
			t.Log(fmt.Sprintf("recovered from %s", err))
		}
	}()

	t.Log("Start")

	panic(errors.New("Something wrong!"))

	// os.Exit 退出时不会调用 defer 指定的函数
	// os.Exit 退出时不输出当前调用栈信息
	//os.Exit(-1)
	//fmt.Println("End")
}

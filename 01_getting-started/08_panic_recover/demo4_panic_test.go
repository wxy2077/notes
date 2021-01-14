/*
* @Time    : 2021-01-13 19:35
* @Author  : CoderCharm
* @File    : demo4_panic_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

panic 多级嵌套

**/
package _8_panic_recover

import (
	"fmt"
	"testing"
)

func TestMultiPanic(t *testing.T) {
	defer fmt.Println("in main")

	//defer func() {
	//	if err := recover(); err!= nil{
	//		fmt.Println("恢复了")
	//	}
	//}()

	defer func() {
		fmt.Println("第二个 defer func")
		defer func() {
			fmt.Println("第三个 defer func")
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}

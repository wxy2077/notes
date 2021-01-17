/*
* @Time    : 2021-01-15 19:47
* @Author  : CoderCharm
* @File    : 04_array_slice_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

数组 和 切片的区别

array vs slice

数组和切片的最大区别就是
	数组是值传递，相当于复制
	切片是指针传递，可直接修改原来的值

**/
package _3_array

import (
	"fmt"
	"testing"
)

func modifyArray(a [3]int) {
	a[0] = 666
}

func modifySlice(n []int) {
	n[0] = 777
}

func TestArrayVSSlice(t *testing.T) {

	// 声明一个 int类型长度为3的数组  数组是值传递 相当于复制
	intArray := [3]int{1, 2, 3}

	t.Log(fmt.Sprintf("数组修改前: %v", intArray))
	modifyArray(intArray)
	t.Log(fmt.Sprintf("数组修改后: %v", intArray))

	// 声明一个 int类型的切片 切片是指针传递
	intSlice := []int{1, 2, 3}

	t.Log(fmt.Sprintf("切片修改前: %v", intSlice))
	// 切片是指针传递,会修改原来的值
	modifySlice(intSlice)
	t.Log(fmt.Sprintf("切片修改后: %v", intSlice))

}

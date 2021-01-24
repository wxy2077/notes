/*
* @Time    : 2021-01-22 08:54
* @Author  : CoderCharm
* @File    : 03_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _9_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect03(t *testing.T) {
	// Go里面不能 直接比较两个 slice map之类的
	// https://stackoverflow.com/questions/37900696/why-cant-go-slice-be-used-as-keys-in-go-maps-pretty-much-the-same-way-arrays-ca
	// Slice, map, and function values are not comparable
	sliceA := []int{1, 2, 3}
	sliceB := []int{1, 2, 3}
	//t.Log(sliceA == sliceB) // panic

	//mapA := map[int]int{1:2, 2:3}
	//mapB := map[int]int{1:2, 2:3}
	t.Log(reflect.DeepEqual(sliceA, sliceB))
	fmt.Println()

	// 支持指针地址比较 但是这样总是不想等的
	//sliceC := &[]int{1, 2, 3}
	//sliceD := &[]int{1, 2, 3}
	//t.Log(sliceC)

}

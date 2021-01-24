/*
* @Time    : 2021-01-18 18:25
* @Author  : CoderCharm
* @File    : 02_test.go
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

func TestIntReflect(t *testing.T) {
	n := map[int]int{1: 2, 2: 3}

	typ := reflect.TypeOf(n)  // a reflect.Type
	val := reflect.ValueOf(n) // a reflect.Type

	t.Log(typ)          // int
	t.Log(typ.String()) // int
	t.Log(typ.Kind())   // 返回此类型的特定种类

	// 注意 数字字符串 不能获取字段数量
	//num := val.NumField() // 获取值字段的数量
	//t.Log(num)
	t.Log(val.Kind())
	// 数组切片 map
	//t.Log(val.Len())

	//t.Log(val.String())
	//t.Log(val.Kind())
	//
	//// 通过反射直接修改类型直接修改
	//pTyp := reflect.TypeOf(&n)
	//pVal := reflect.ValueOf(&n)
	//
	//t.Log(pTyp)
	//
	//pVal.Elem().SetInt(999)
	//
	//t.Log(n)

}

func TestSliceReflect(t *testing.T) {
	//sliceA := []string{"a", "b", "c"}
	//check(sliceA)
	z := -1123

	check(z)
}

func check(anyData interface{}) {

	t := reflect.TypeOf(anyData)
	fmt.Println(t)

	v := reflect.ValueOf(anyData)

	switch v.Kind() {
	case reflect.Slice:
		fmt.Println("是一个切片")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("整数类型")
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Println("无符号类型")
	case reflect.Float32, reflect.Float64:
		fmt.Println("浮点类型")
	default:
		fmt.Println("有问题")
	}

}

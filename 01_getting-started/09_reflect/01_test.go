/*
* @Time    : 2021-01-17 15:40
* @Author  : CoderCharm
* @File    : 01_test.go
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

type userInfo struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	User     string `json:"user"foo:"test_tag"`
	Avatar   string `json:"avatar"bar:""`
}

// 常见的类型检测场景
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestReflect(t *testing.T) {

	// 实例化结构体
	user := userInfo{Page: 1, PageSize: 60, User: "Nike"}

	typ := reflect.TypeOf(user) // 获取reflect的类型
	t.Log(typ)

	val := reflect.ValueOf(user) // 获取reflect的值
	t.Log(val)

	kd := val.Kind() // 获取到st对应的类别

	t.Log(kd)

	num := val.NumField() // 获取值字段的数量
	t.Log(num)

	tagVal := typ.Field(2) // 获取index为2的类型信息
	val = val.Field(2)     // 获取index为2实例化后的值

	t.Log(tagVal)
	t.Log(val)
	t.Log(tagVal.Type) // 输出结构体字段的类型
	t.Log(tagVal.Name) // 输出结构体字段名称

	// 获取结构体的tag 没有则为空 Get 实际就是调用的Lookup
	t.Log(tagVal.Tag.Get("foo"))
	// 返回两个值 第一个为tag值 第二个为bool值 true表示设置了此tag 无论是否为空字符串
	t.Log(tagVal.Tag.Lookup("foo"))
	t.Log(typ.Field(3).Tag.Lookup("bar"))     // 设置了tag为bar 但是为空字符串 依旧为true
	t.Log(typ.Field(3).Tag.Lookup("ant_tag")) // 没有设置此tag 就为false

}

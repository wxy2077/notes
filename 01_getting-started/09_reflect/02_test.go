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
	"reflect"
	"testing"
)

func TestIntReflect(t *testing.T) {
	n := 3

	typ := reflect.TypeOf(n)  // a reflect.Type
	val := reflect.ValueOf(n) // a reflect.Type

	t.Log(typ)          // int
	t.Log(typ.String()) // int
	t.Log(typ.Kind())

	t.Log(val)
	t.Log(val.String())
	t.Log(val.Kind())

	//t.Log(n.NumField())
	//t.Log(z.Field(0))

	user := userInfo{Page: 1, PageSize: 60, User: "Nike"}

	typ = reflect.TypeOf(user) // 获取reflect的类型
	t.Log(typ)

}

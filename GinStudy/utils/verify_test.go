/*
* @Time    : 2021-01-15 19:25
* @Author  : CoderCharm
* @File    : verify_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
测试参数验证
**/
package utils

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func verify(st interface{}, roleMap Rules) (err error) {
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}

	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st) // 获取reflect.Type类型

	kd := val.Kind() // 获取到st对应的类别

	if kd != reflect.Struct {
		return errors.New("expect struct")
	}
	num := val.NumField()

	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i)
		val := val.Field(i)

		// 如果有此规则设置 就验证
		if len(roleMap[tagVal.Name]) > 0 {
			for _, v := range roleMap[tagVal.Name] {
				switch {
				case v == "notEmpty":
					if isBlank(val) {
						return errors.New(tagVal.Name + "值不能为空")
					}
				case compareMap[strings.Split(v, "=")[0]]:
					if !compareVerify(val, v) {
						return errors.New(tagVal.Name + "长度或值不在合法范围," + v)
					}
				}
			}
		}
	}
	return nil
}

type userInfo struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	User     string `json:"user"a:"123"`
	Avatar   string `json:"avatar"`
}

var (
	demoInfoVerify = Rules{"Page": {Ge("1")}, "PageSize": {Le("50")}, "User": {NotEmpty()}}
)

func TestVerify(t *testing.T) {

	u := userInfo{Page: 1, PageSize: 60, User: "Nike"}

	typ := reflect.TypeOf(u)  // 获取reflect的类型
	val := reflect.ValueOf(u) // 获取reflect的值

	t.Log(typ)
	t.Log(val)

	kd := val.Kind() // 获取到st对应的类别

	t.Log(kd)

	num := val.NumField() // 获取值字段的数量
	t.Log(num)

	tagVal := typ.Field(2) // 获取index为2的类型信息
	val = val.Field(2)     // 获取index为2实例化后的值

	t.Log(tagVal)
	t.Log(val)
	t.Log(tagVal.Type)
	t.Log(demoInfoVerify[tagVal.Name])

	z := len(demoInfoVerify["User"])
	t.Log(z)

	//if err := verify(u, demoInfoVerify); err != nil {
	//	t.Log(err)
	//}
	t.Log("over")

}

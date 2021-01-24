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
	User   string `json:"user"foo:"test_tag"`
	Avatar string `json:"avatar"bar:""`
	Age    int    `json:"age"`
}

func (u userInfo) SayName(name string) {
	fmt.Printf("用户名是 %s\n", name)
}

func (u userInfo) SayAge() {
	fmt.Printf("年龄是 %d\n", u.Age)
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

// 通过断言判断类型
func AssertionType(v interface{}) {
	switch v.(type) {
	case int, int16, int32, int64:
		fmt.Printf("整数类型 %d \n", v)
	case userInfo:
		fmt.Printf("是 userInfo 类型 %v  ", v)
		fmt.Printf("年龄 %d \n", v.(userInfo).Age)
		// 断言调用方法
		v.(userInfo).SayAge()
	default:
		fmt.Println("没有匹配到")
	}
}

func TestReflect(t *testing.T) {

	// 实例化结构体
	user := userInfo{User: "Nike", Age: 18}

	// 断言判断类型 功能不如反射强大
	AssertionType(user)
	AssertionType(1111)

	typ := reflect.TypeOf(user) // 获取reflect的类型
	t.Log(typ)

	val := reflect.ValueOf(user) // 获取reflect的值
	t.Log(val)

	kd := val.Kind() // 获取对应的类别
	t.Log(kd)

	num := val.NumField() // 获取值字段的数量
	t.Log(num)

	// 通过反射调用方法
	m1 := val.MethodByName("SayName")
	//m1 := val.Method(0)
	m1.Call([]reflect.Value{reflect.ValueOf("测试")}) // 传参数
	// 私有方法不可反射调用  Java反射可以暴力调用私有方法
	m2 := val.MethodByName("SayAge")
	m2.Call([]reflect.Value{}) // 不传参数

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
	t.Log(typ.Field(1).Tag.Lookup("bar"))     // 设置了tag为bar 但是为空字符串 依旧为true
	t.Log(typ.Field(1).Tag.Lookup("any_tag")) // 没有设置此tag 就为false

	// 必须使用地址 才可以修改原来的值 否则会panic (反射第三定律，值可以被修改)
	modifyVal(&user)

	t.Log(user)
}

// 通过反射修改值
func modifyVal(user interface{}) {

	// 获取变量的指针
	pVal := reflect.ValueOf(user) // 获取reflect的值

	// 获取指针指向的变量
	v := pVal.Elem()
	// 找到并更新变量的值
	v.FieldByName("User").SetString("Jack")

}

/*
* @Time    : 2021-01-20 21:44
* @Author  : CoderCharm
* @File    : demo_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _3_casbin_gorm

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"strings"
	"testing"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ParamsMatch
//@description: 自定义规则函数
//@param: fullNameKey1 string, key2 string
//@return: bool

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ParamsMatchFunc
//@description: 自定义规则函数
//@param: args ...interface{}
//@return: interface{}, error

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}

// 动态添加角色权限
func TestCasbin(t *testing.T) {

	a, _ := gormadapter.NewAdapter("mysql", "root:Admin12345-@tcp(172.16.137.129:3306)/go_casbin", true)

	e, err := casbin.NewEnforcer("./model.conf", a)

	if err != nil {
		panic(err)
	}

	e.AddFunction("ParamsMatch", ParamsMatchFunc)

	_ = e.LoadPolicy()

	ok, err := e.AddPolicy("99", "/api/getJob", "GET")

	//sub := "10"  // the user that wants to access a resource.
	//obj := "/api/getList" // the resource that is going to be accessed.
	//act := "GET"  // the operation that the user performs on the resource.
	//
	////验证
	//ok, err := e.Enforce(sub, obj, act)
	//
	if err != nil {
		// handle err
		t.Log(fmt.Sprintf("错误 %v", err))
	}

	if ok == true {
		// permit alice to read data1
		t.Log("ok")
	} else {
		// deny the request, show an error
		t.Log("fail")
	}

}

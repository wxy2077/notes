/*
* @Time    : 2021-01-18 23:48
* @Author  : CoderCharm
* @File    : casbin_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

前置条件先搞懂 各个模型 策略的关系
https://casbin.org/docs/zh-CN/how-it-works
https://casbin.org/zh-CN/editor


本文件代码来源
https://casbin.org/docs/zh-CN/get-started

go get github.com/casbin/casbin/v2

**/
package _1_official_demo

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"testing"
)

// 最基础的demo
func TestCasbin01(t *testing.T) {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")

	sub := "nick"  // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	ok, err := e.Enforce(sub, obj, act)

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

// 动态添加角色权限
func TestCasbin02(t *testing.T) {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")

	sub := "Jack"  // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	// 动态添加策略
	added, err := e.AddPolicy("Jack", "data1", "read")

	t.Log(added)

	ok, err := e.Enforce(sub, obj, act)

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

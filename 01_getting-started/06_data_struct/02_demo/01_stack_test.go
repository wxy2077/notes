/*
* @Time    : 2020-08-28 00:03
* @Author  : CoderCharm
* @File    : 01_stack_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _2_demo

import (
	"testing"
)

type Foo struct {
	Allen string
	Levy  string
}

// Foo 指定类型  Spoke 与 Allen Levy 变量 不在同一层
type Bar struct {
	Spoke int
	Foo   Foo
}

// 相当于继承 Foo   Zek 和 Allen Levy 变量在同一层
type Titan struct {
	Zek string
	Foo
}

func TestA(t *testing.T) {
	b := Bar{1, Foo{"阿尔敏", "利维"}}
	t.Log(b)
	t.Log(b.Spoke)
	//t.Log(b.Allen) // 不能取到
	t.Log(b.Foo.Allen) // 只能这样取到

	t.Log("------\n")

	titan := Titan{"吉克", Foo{"aaaaa", "llll"}}
	t.Log(titan)
	t.Log(titan.Zek)
	t.Log(titan.Allen) // 可以直接取到 Foo的值
	t.Log(titan.Levy)
	t.Log(titan.Foo)
}

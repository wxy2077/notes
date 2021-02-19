/*
* @Time    : 2021-02-17 14:43
* @Author  : CoderCharm
* @File    : factory_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

go语言里面结构体没有构造函数这个说法，所以一般会采用`Newxxx`这样的来初始化
NewXXX 函数返回接口时就是简单工厂模式。

**/
package _1_simple_factory

import (
	"testing"
)

// 声明一个接口
type Animal interface {
	action() string
}

// 声明一个结构体 实现action方法
type Dog struct {
	Bark string
}

func (d *Dog) action() string {
	return d.Bark
}

// 同样声明另一个
type Cat struct {
	Sound string
}

func (c *Cat) action() string {
	return c.Sound
}

// 实例化
func NewAnimal(category bool) Animal {
	if category {
		return &Cat{"小猫喵喵喵"}
	} else {
		return &Dog{"小狗汪汪叫"}
	}
}

func TestFactory(t *testing.T) {
	a := NewAnimal(true)
	t.Log(a.action())
}

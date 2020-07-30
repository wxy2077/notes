package main

import (
	"fmt"
)

// 声明一个结构体
type Cat struct {
	name   string
	color  string
	sex    bool
	couple Mouse // 这里赋值为 另一个结构体, 如果和结构体名一样 可以省略一个
}

// 结构体挂载方法
func (c *Cat) CaptureMouse(name string) (res bool) {
	if c.name == name {
		fmt.Println("名字为Tom，能捉到老鼠")
		// 由于使用了指针， 可以改变结构体的值
		c.name = "old tom "
		return true
	} else {
		fmt.Println("名字不为Tom， 捉不到老鼠")
		return false
	}
}

// 在定义一个结构体
type Mouse struct {
	name  string
	color string
	sex   bool
}

// 结构体挂载方法 不使用指针
func (m Mouse) Eat(stuff string) (res bool) {
	if stuff == "奶酪" {
		fmt.Println(m.name, "喜欢吃", stuff)
		// 不能改变外部结构体的值
		m.name = "old jerry"
		return true
	} else {
		fmt.Println(m.name, "不喜欢吃", stuff)
		return false
	}
}

// 模拟构造方法
func NewMouse(name string, color string, sex bool) (m Mouse) {
	return Mouse{
		name:  name,
		color: color,
		sex:   sex,
	}
}

func main() {
	// 实例化结构体
	Tom := Cat{
		name:  "Tom",
		color: "white",
		sex:   true,
		couple: Mouse{
			name: "Jerry",
		},
	}

	// 调用方法
	res := Tom.CaptureMouse("Tom")
	fmt.Println(res)
	fmt.Println(Tom)
	// 通过子类属性调用子类的方法(我自己取的名字)
	Tom.couple.Eat("糖")
	Tom.couple.Eat("奶酪")

	// 构造方法实例化
	Tuffy := Mouse{"Tuffy", "grey", true}
	fmt.Println(Tuffy)
}

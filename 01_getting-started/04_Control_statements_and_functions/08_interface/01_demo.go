/*
接口练习

https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.6.md
https://www.bilibili.com/video/BV1WA411b7ZF
*/
package main

import "fmt"

// 定义一个动物接口 有两个方法
type Animal interface {
	Eat(food string)
	Run()
}

// 定一个Cat结构体
type Cat struct {
	name string
}

// 实现接口的方法
func (c *Cat) Eat(food string) {
	fmt.Println(c.name + " 吃 " + food)
}
func (c Cat) Run() {
	fmt.Println(c.name + "跑223")
}

// 定义一个Dog结构体
type Dog struct {
	name string
}

func (d Dog) Eat(food string) {
	fmt.Println(d.name + " 吃 " + food)
}

func (d Dog) Run() {
	fmt.Println(d.name + "跑666")
}

func main() {
	// 声明一个接口
	var A Animal
	// 实例结构体
	Tom := Cat{name: "Tom"}

	// 只要实现一个接口的方法为指针 就得传地址
	A = &Tom
	A.Eat("鱼")
	A.Run()

	// 声明接口B
	var B Animal
	// 直接 把实例化对象 赋值给接口B
	B = Dog{name: "Park"}
	// 接口B调用方法
	B.Eat("骨头")
	B.Run()

}

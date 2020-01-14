package main

import "fmt"

//a := 1  // error
var abc string

//示例代码
var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明

func test() {
	var available bool // 一般声明
	valid := false     // 简短声明
	available = true   // 赋值操作
	abc = "abcAbc"
	fmt.Println("-----")
	fmt.Println(abc[1:4])
	fmt.Println("-----")
	fmt.Println(available)
	fmt.Println(valid)
	fmt.Println(isActive)

}

func main() {
	//a := 1
	//fmt.Print("hello")
	//fmt.Println(a)
	//
	//fmt.Printf("%d \n", 22)
	//fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	//fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	//fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	//fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	//a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `
	Do you like



	my hat?
	`

	g := 'M'
	var h = "aaa"

	var k, l int
	l = 7
	var m, n string = "mmm", "abc"
	//zzz := 123   //  变量声明之后必须得使用

	o := `this is golang program`

	fmt.Printf("%v \n", abc)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Println(h)
	fmt.Println(k, l)
	fmt.Println(m, n)
	fmt.Println("o - ", o)
	test()
}

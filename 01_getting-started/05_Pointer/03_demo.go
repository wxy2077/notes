package main

import "fmt"

// 定义一个结构体
type Cat struct {
	name  string
	color string
	sex   bool
}

func main() {
	// 实例化
	tom := Cat{
		name:  "Tom",
		color: "white",
		sex:   true,
	}
	fmt.Println(tom)

	// 定义一个结构体指针
	var tom2 *Cat
	// 地址赋值给结构体指针
	tom2 = &tom

	fmt.Println(tom)
	// 结构体指针改变 值
	(*tom2).name = "old tom"

	fmt.Println(tom)
}

/*
练习执政
*/
package main

import "fmt"

func main() {
	a := 123
	fmt.Println(a)
	var b *int // 定义b为指针类型
	b = &a     // & 取地址符号
	*b = 666   // 表示的就是值
	fmt.Println(a, &a, b, *b, &*b)
	// 函数传入地址， 函数内部可以修改值
	changeA1(&a)
	fmt.Println(a) // 777
	changeA2(a)
	fmt.Println(a) // 还是777 不会变成888
	// 同理其他 数据类型 数组 map struct等 可以定义指针

}

func changeA1(z *int) {
	*z = 777
}

func changeA2(n int) {
	n = 888
}

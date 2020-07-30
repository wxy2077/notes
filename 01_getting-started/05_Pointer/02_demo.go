/*

 */
package main

import (
	"fmt"
)

func main() {
	a := [3]string{"a", "b", "c"}
	// 定义一个数组指针 外部是一个指针 内部是元素
	var b *[3]string

	b = &a
	// 先通过*b取出array值 然后修改 array里面的值
	(*b)[0] = "zzz"

	fmt.Println(a)

	c := [3]string{"q", "w", "e"}

	// 定义一个指针数组 外部是array 里面元素是指针
	d := [3]*string{&c[0], &c[1], &c[2]}
	fmt.Println(c)
	*d[0] = "qqqqq"
	fmt.Println(d, *d[0])
	fmt.Println(c)

	// 定义一个map key为值 value为 字符串指针
	e := map[string]*string{}
	f := "r"
	e["a"] = &f
	fmt.Println("e是", e)

	// 定义一个map key value都为 字符串指针
	g := map[*string]*string{}
	h := "t"
	i := "y"
	// 字符串指针 key赋予地址 value也赋予地址
	g[&h] = &i
	fmt.Println("g是", g, "取value值", *g[&h])
}

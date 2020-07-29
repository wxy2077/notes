/*
函数传入和返回不同类型的参数
*/
package main

import "fmt"

func main() {
	a := [3]int{4, 5, 6}
	fmt.Println(test1(a))
}

func test1(t1 [3]int) (v int) {
	z := 0
	for _, v := range t1 {
		z += v
	}
	return v
}

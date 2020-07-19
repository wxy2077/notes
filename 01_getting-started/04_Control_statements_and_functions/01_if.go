package main

import "fmt"

func main() {

	a := 1
	//b := getB()
	fmt.Printf("a: %d - b: %d\n", a, getB())
	if a > getB() {
		fmt.Println("a > b")
	} else {
		fmt.Println("a < b")
	}

	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := getB(); x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
	//fmt.Println(x)

}

func getB() int {
	return 5
}

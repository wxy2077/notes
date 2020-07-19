/*

goto

Go有goto语句——请明智地使用它。用goto跳转到必须在当前函数内定义的标签。




例如假设这样一个循环：

在我的印象里 就C 有goto，而且感觉这个语法很不友好，用不好很混乱，
Python 直接没有goto
Java goto作为保留字，也没有用

*/
package main

func main() {
	print("123")
	myFunc()
}

func myFunc() {
	i := 0
Here: //这行的第一个词，以冒号结束作为标签  标签名是大小写敏感的。
	println(i)
	i++
	if i >= 10 {
		return
	}
	goto Here //跳转到Here去
}

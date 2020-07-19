/*

参数

变参
Go函数支持变参。接受变参的函数是有着不定数量的参数的。为了做到这点，首先需要定义函数使其接受变参：

传值与传指针
当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一份copy，当在被调用函数中修改参数值的时候，调用函数中相应实参不会发生任何变化，因为数值变化只作用在copy上。

为了验证我们上面的说法，我们来看一个例子

*/
package main

import "fmt"

//变参
func haHa(arg ...int) {
	// arg ...int告诉Go这个函数接受不定数量的参数。
	// 注意，这些参数的类型全部是int。在函数体中，变量arg是一个int的slice：
	for i, v := range arg {
		fmt.Println(i, v)
	}

}

//简单的一个函数，实现了参数+1的操作
func add1(a *int) (z int) { // 请注意，
	*a = *a + 1 // 修改了a的值
	return *a
}

/*

defer v.推迟；延缓；展期；
Go语言中有种不错的设计，即延迟（defer）语句，你可以在函数中添加多个defer语句。当函数执行到最后时，
这些defer语句会按照逆序执行，最后该函数返回。特别是当你在进行一些打开资源的操作时，遇到错误需要提前返回，
在返回前你需要关闭相应的资源，不然很容易造成资源泄露等问题。如下代码所示，我们一般写打开一个资源是这样操作的：
*/

func myDefer() {
	a := "aaaa"
	b := "bbbb"

	defer fmt.Println(a)
	defer fmt.Println(b)

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func main() {
	haHa(4, 5, 8, 9)

	x := 1

	fmt.Println("x = ", x) // 应该输出 "x = 3"

	x1 := add1(&x) // 调用 add1(&x) 传x的地址
	//add1(&x)

	fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出 "x = 4"

	myDefer()

}

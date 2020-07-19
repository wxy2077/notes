/*
函数是Go里面的核心设计，它通过关键字func来声明，它的格式如下：

func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	return value1, value2
}

上面的代码我们看出

关键字func用来声明一个函数funcName
函数可以有一个或者多个参数，每个参数后面带有类型，通过,分隔
函数可以返回多个值
上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型
如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
如果没有返回值，那么就直接省略最后的返回信息
如果有返回值， 那么必须在函数的外层添加return语句
*/

package main

import "fmt"

func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func main() {
	var maxValue int

	a, b := 1, 3

	var (
		x = 1
		y = 2
		z = 3
	)

	fmt.Println(myMax(x, y))
	fmt.Println(myMax(x, z))
	fmt.Println(myMax(y, z))

	maxValue = myMax(a, b)

	fmt.Println(maxValue)

	fmt.Println("-----")

	e, f := SumAndProduct(x, z)

	fmt.Println(e, f)

}

/*
上面的例子我们可以看到直接返回了两个参数，当然我们也可以命名返回参数的变量，这个例子里面只是用了两个类型，我们也可以改成如下这样的定义，
然后返回的时候不用带上变量名，因为直接在函数里面初始化了。但如果你的函数是导出的(首字母大写)，
官方建议：最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差。
*/

func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

/*

for 循环 没啥特殊的, 基础用法和js java里面差不多

Go里面最强大的一个控制逻辑就是for，它既可以用来循环读取数据，又可以当作while来控制逻辑，还能迭代操作。它的语法如下：

for expression1; expression2; expression3 {
	//...
}

expression1、expression2和expression3都是表达式，
其中expression1和expression3是变量声明或者函数调用返回值之类的，
expression2是用来条件判断，expression1在循环开始之前调用，expression3在每轮循环结束之时调用。

一个例子比上面讲那么多更有用，那么我们看看下面的例子吧：

*/
package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}

	fmt.Println("sum is equal to ", sum)

	myBBB()
	myCCC()
	myEEE()
	myFFF()
}

// 输出：sum is equal to 45

/*
有些时候需要进行多个赋值操作，由于Go里面没有,操作符，那么可以使用平行赋值i, j = i+1, j-1

有些时候如果我们忽略expression1和expression3：
*/

func myBBB() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println("myBBB>>", sum)
}

/*
其中;也可以省略，那么就变成如下的代码了，是不是似曾相识？对，这就是while的功能。

*/
func myCCC() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("myCCC>>", sum)

}

/*
在循环里面有两个关键操作break和continue ,break操作是跳出当前循环，continue是跳过本次循环。
当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置，详细参考如下例子：
*/

func myDDD() {
	for index := 10; index > 0; index-- {
		if index == 5 {
			break // 或者continue
		}
		fmt.Println(index)
	}
	// break打印出来10、9、8、7、6
	// continue打印出来10、9、8、7、6、4、3、2、1
}

/*
break和continue还可以跟着标号，用来跳到多重循环中的外层循环

for配合range可以用于读取slice和map的数据：
*/
func myEEE() {

	a := make(map[string]int)
	a["a"] = 1
	a["b"] = 2
	a["c"] = 999

	for k, v := range a {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}
}

/*
循环 slice 类似
*/

func myFFF() {
	a := []int32{1, 2, 3}

	for i, v := range a {
		fmt.Println("slice's val:", i, v)
	}
}

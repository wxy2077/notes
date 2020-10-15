/*

array就是数组，它的定义方式如下：

var arr [n]type

在[n]type中，n表示数组的长度，type表示存储元素的类型。对数组的操作和其它语言类似，都是通过[]来进行读取或赋值：


*/
package main

import (
	"fmt"
	"reflect"
)

func test1() {
	var arr [10]int                                  // 声明了一个int类型的数组
	arr[0] = 42                                      // 数组下标是从0开始的
	arr[1] = 13                                      // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
	fmt.Printf("The second element is %d\n", arr[1]) // 获取数据，返回13
	fmt.Printf("The last element is %d\n", arr[9])   //返回未赋值的最后一个元素，默认返回0
	fmt.Printf("The all elements is %d\n", arr)      //输出全部数据
	var myAll [2]string
	//myAll = append(myAll, 1)
	//myAll = append(myAll, 1)
	fmt.Printf("myAll type: %T \n", myAll)
	fmt.Println("myAll type:", reflect.TypeOf(arr))
	fmt.Println("-----------------")
}
func test2() {
	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

	b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Println(doubleArray)
	fmt.Println(easyArray)
}

func test3() {
	// interface{} 就能放入各种类型的值
	xy := []interface{}{1, "2", "3", true}
	fmt.Println(xy)
}

func test4() {
	// 数组的增删改查
	tempXy := []string{"a", "b", "c", "d"}
	tempXy = append(tempXy, "e")
	z := tempXy[:1]
	z[0] = "zzz"
	fmt.Println(z)
	fmt.Println(tempXy)
	// 再次添加 一遍自己
	tempXy = append(tempXy, tempXy...)
	fmt.Println(tempXy)
}

func main() {

	//test1()
	//test2()
	//test3()
	test4()
}

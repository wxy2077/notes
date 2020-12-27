/*

array就是数组，它的定义方式如下：

var arr [n]type

在[n]type中，n表示数组的长度，type表示存储元素的类型。对数组的操作和其它语言类似，都是通过[]来进行读取或赋值：


*/
package _3_array

import (
	"reflect"
	"testing"
)

func Test01(t *testing.T) {
	var arr [10]int                              // 声明了一个int类型的数组
	arr[0] = 42                                  // 数组下标是从0开始的
	arr[1] = 13                                  // 赋值操作
	t.Log("The first element is %d \n", arr[0])  // 获取数据，返回42
	t.Log("The second element is %d \n", arr[1]) // 获取数据，返回13
	t.Log("The last element is %d \n", arr[9])   //返回未赋值的最后一个元素，默认返回0
	t.Log("The all elements is %d \n", arr)      //输出全部数据
	var myAll [2]string
	//myAll = append(myAll, 1)
	//myAll = append(myAll, 1)
	t.Log("myAll type: %T \n", myAll)
	t.Log("myAll type:", reflect.TypeOf(arr))
	t.Log("-----------------")
}
func Test02(t *testing.T) {
	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组

	b := [10]int{1, 2, 3} // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0

	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

	t.Log(a)
	t.Log(b)
	t.Log(c)

	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// 上面的声明可以简化，直接忽略内部的类型
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	t.Log(doubleArray)
	t.Log(easyArray)
}

func Test03(t *testing.T) {
	// interface{} 就能放入各种类型的值
	xy := []interface{}{1, "2", "3", true}
	t.Log(xy)
}

func Test04(t *testing.T) {
	// 数组的增删改查
	tempXy := []string{"a", "b", "c", "d"}
	tempXy = append(tempXy, "e")
	z := tempXy[:1]
	z[0] = "zzz"
	t.Log(z)
	t.Log(tempXy)
	// 再次添加 一遍自己
	tempXy = append(tempXy, tempXy...)
	t.Log(tempXy)
}

/*
* @Time    : 2020-10-15 09:55
* @Author  : CoderCharm
* @File    : test02_array.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

快速判断字符串是否在一个数组中 类似Python in 函数
https://mp.weixin.qq.com/s/ffcGphLoU-V-TO31wbJjsg
**/
package _3_array

import (
	"fmt"
	"sort"
	"testing"
)

// 首先用循环实现 但是这种方式有一个弊端，就是要遍历整个字符串数组 这是一个非常费时间的操作
func loopIn(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element {
			return true
		}
	}
	return false
}

// 如果是有序的整型数组，那么我们可以使用二分查找，把时间复杂度O(n)降到对数时间复杂度。字符串能不能也这样操作呢？实际上是可以的。
func in(target string, strArray []string) bool {
	// sort.Strings()函数，可以对字符串数组进行排序
	sort.Strings(strArray)

	// sort.SearchStrings()函数，会用二分法在一个有序字符串数组中寻找特定字符串的索引
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func TestArray(t *testing.T) {
	nameList := []string{"赵云", "王小右", "张飞"}
	myName := "王小右"
	fmt.Println(loopIn(myName, nameList))
	fmt.Println(in(myName, nameList))
}

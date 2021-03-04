/*
* @Time    : 2021-03-04 20:55
* @Author  : CoderCharm
* @File    : first_show_string_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

简单

在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

示例:

s = "abaccdeff"
返回 "b"

s = ""
返回 " "


限制：

0 <= s 的长度 <= 50000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


**/
package _0_first_show_string

import (
	"testing"
)

// 时间复杂度 O(2N) 两次for循环 随着N增长而增长
// 空间复杂度 O(1)  list 大小是固定的，并不会随着N增长变化
func firstUniqChar(s string) byte {
	// 一共26个字母 构建长度26的数组
	list := make([]int, 26)
	for _, v := range s {
		// 循环字符串为bite - 97即a 为数组中的位置  +1 表示有值
		list[v-97]++
	}
	for _, v := range s {

		// 当前位置有值
		if list[v-97] == 1 {
			return byte(v)
		}
	}
	return ' '
}

func TestFirstChar(t *testing.T) {
	z := firstUniqChar("abaccdeff")
	t.Log(z)
}

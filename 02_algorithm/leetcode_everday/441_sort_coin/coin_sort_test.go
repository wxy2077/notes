/*
* @Time    : 2021/3/18 22:21
* @Author  : CoderCharm
* @File    : coin_sort_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :


441 硬币排序 简单

你总共有n枚硬币，你需要将它们摆成一个阶梯形状，第k行就必须正好有 k 枚硬币。

给定一个数字n，找出可形成完整阶梯行的总行数。

n是一个非负整数，并且在32位有符号整型的范围内。

示例 1:

n = 5

硬币可排列成以下几行:
¤
¤ ¤
¤ ¤

因为第三行不完整，所以返回2.
示例 2:

n = 8

硬币可排列成以下几行:
¤
¤ ¤
¤ ¤ ¤
¤ ¤

因为第四行不完整，所以返回3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/arranging-coins
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


**/
package _41_sort_coin

import "testing"

// 解法 1
// 直接循环 减去每一行所需要的硬币 剩下的硬币 >= 当前行 则继续分， 否则直接返回行数
// 时间复杂度 O(N) 就一个for循环
// 空间复杂度 O(1) 也没涉及到数据的占用空间大的变量
// 执行用时: 8 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗: 2.2 MB, 在所有 Go 提交中击败了100.00%的用户
func arrangeCoins(n int) int {
	// 默认第一行为1
	i := 1

	// 判断最后 一行的条件是否满足
	for n >= i {
		// n 依次减去每行的硬币
		n -= i
		// 最后 + 1 代表换行
		i++
	}
	// 减去 最后 + 1 多的值
	return i - 1
}

// 解法 2 二分查找
// 看的答案 明显这种 执行用时最短
// 执行用时: 0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗: 2.2 MB, 在所有 Go 提交中击败了100.00%的用户
//
func sortCoins(n int) int {
	l, r := 0, n
	for l <= r {
		mid := (l + r) >> 1
		v := n - (mid+1)*mid/2
		switch {
		case v == 0:
			return mid
		case v > 0:
			l = mid + 1
		case v < 0:
			r = mid - 1
		}
	}
	return r
}

func TestSortCoin(t *testing.T) {
	t.Log(arrangeCoins(10))
	t.Log(sortCoins(10))
}

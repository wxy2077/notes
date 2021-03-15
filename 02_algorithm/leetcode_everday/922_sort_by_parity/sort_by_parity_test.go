/*
* @Time    : 2021/3/15 23:02
* @Author  : CoderCharm
* @File    : sort_by_parity_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

922. 按奇偶排序数组 II  难度简单

给定一个非负整数数组A， A 中一半整数是奇数，一半整数是偶数。

对数组进行排序，以便当A[i] 为奇数时，i也是奇数；当A[i]为偶数时， i 也是偶数。

你可以返回任何满足上述条件的数组作为答案。

示例：

输入：[4,2,5,7]
输出：[4,5,2,7]
解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。


提示：

2 <= A.length <= 20000
A.length % 2 == 0
0 <= A[i] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-array-by-parity-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

思路: 用两个指针 一个为奇数 一个为偶数，不断移动其中一个，每次移动两个单位，根据条件跳转位置。

时间复杂度: O(N)
空间复杂度：O(1)
**/
package _22_sort_by_parity

import "testing"

// 假设输入的值是正确的
func sortArrayByParityII(nums []int) []int {
	// 使用双指针 i为偶数下角标 j为奇数
	for i, j := 0, 1; i < len(nums); i += 2 {

		// 判断是否为奇数
		if nums[i]%2 == 1 {
			// 循环找到奇数的值
			for nums[j]%2 == 1 {
				j += 2
			}
			// 两树互换
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}

func TestSort(t *testing.T) {
	t.Log(sortArrayByParityII([]int{4, 2, 5, 7}))
}

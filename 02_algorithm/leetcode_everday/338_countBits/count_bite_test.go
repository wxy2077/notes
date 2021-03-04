/*
* @Time    : 2021-03-03 21:32
* @Author  : CoderCharm
* @File    : count_bite_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :


338. 比特位计数 中等

给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:

输入: 2
输出: [0,1,1]
示例 2:

输入: 5
输出: [0,1,1,2,1,2]
进阶:

给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
要求算法的空间复杂度为O(n)。
你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/counting-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


**/
package _38_countBits

import "testing"

func onesCount(x int) (ones int) {
	// 利用位运算的技巧，可以在一定程度上提升计算速度。按位与运算（&）的一个性质是：
	// 对于任意整数 x，令 x=x&(x−1)，该运算将 x 的二进制表示的最后一个 1 变成 0。
	// 因此，对 x 重复该操作，直到 x 变成 0，则操作次数即为 x 的「一比特数」。

	for ; x > 0; x &= x - 1 {
		ones++
	}
	return
}

func countBits(num int) []int {
	//
	bits := make([]int, num+1)

	for i := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

// 解法2 评论上的
func count2Bites(num int) []int {
	//	1: 0001     3:  0011      0: 0000
	//	2: 0010     6:  0110      1: 0001
	//	4: 0100     12: 1100      2: 0010
	//	8: 1000     24: 11000     3: 0011
	//	16:10000    48: 110000    4: 0100
	//	32:100000   96: 1100000   5: 0101
	//
	//	由上可见：
	//	1、如果 i 为偶数，那么f(i) = f(i/2) ,因为 i/2 本质上是i的二进制左移一位，低位补零，所以1的数量不变。
	//	2、如果 i 为奇数，那么f(i) = f(i - 1) + 1， 因为如果i为奇数，那么 i - 1必定为偶数，而偶数的二进制最低位一定是0，
	//	那么该偶数 +1 后最低位变为1且不会进位，所以奇数比它上一个偶数bit上多一个1，即 f(i) = f(i - 1) + 1。
	// 时间复杂度 O(N) num越大循环次数越多
	// 空间复杂度 O(N) num越大，res 切片不断的增大
	res := []int{0}
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			res = append(res, res[i/2])
		} else {
			res = append(res, res[i-1]+1)
		}
	}
	return res

}

func TestCount(t *testing.T) {

	t.Log(countBits(2))
	t.Log(countBits(5))

	t.Log(count2Bites(2))
	t.Log(count2Bites(5))

}

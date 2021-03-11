/*
* @Time    : 2021/3/11 23:52
* @Author  : CoderCharm
* @File    : calc_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

中级

给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

整数除法仅保留整数部分。

示例 1：

输入：s = "3+2*2"
输出：7
示例 2：

输入：s = " 3/2 "
输出：1
示例 3：

输入：s = " 3+5 / 2 "
输出：5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/basic-calculator-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

1 <= s.length <= 3 * 105
s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
s 表示一个 有效表达式
表达式中的所有整数都是非负整数，且在范围 [0, 231 - 1] 内
题目数据保证答案是一个 32-bit 整数

思路:
	1 减法可以用 负数表示

**/
package basic_calc_2

import "testing"

func calculate(s string) (res int) {
	// 用于存储 计算好乘法和除法的数据 最后各项相加即可
	var stack []int
	// 默认计算符号为 + 号
	preSign := '+'
	// 由于不使用eval函数  用这种方式把字符串还原成数字
	num := 0

	for k, v := range s {
		// 判断是否是数字
		isDigit := '0' <= v && v <= '9'
		if isDigit {
			// 把字符串还原成数字
			num = num*10 + int(v-'0')
		}

		// 根据运算符判断操作
		if !isDigit && v != ' ' || k == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				// 乘除法 可以直接与栈顶元素计算
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = v
			num = 0
		}
	}
	// 栈内 数据累加
	for _, v := range stack {
		res += v
	}
	return

}

func TestCalc(t *testing.T) {

	t.Log(calculate(" 3/2 "))
	num := 0

	num = num*10 + int('3'-'0')
	t.Log(num)

}

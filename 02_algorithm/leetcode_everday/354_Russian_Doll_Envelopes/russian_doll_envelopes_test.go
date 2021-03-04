/*
* @Time    : 2021-03-04 20:46
* @Author  : CoderCharm
* @File    : russian_doll_envelopes_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

俄罗斯套娃 困难

给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。

当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

注意：不允许旋转信封。


示例 1：

输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
示例 2：

输入：envelopes = [[1,1],[1,1],[1,1]]
输出：1


提示：

1 <= envelopes.length <= 5000
envelopes[i].length == 2
1 <= wi, hi <= 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/russian-doll-envelopes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

动态规划 不熟 后面再回来解
**/
package _54_Russian_Doll_Envelopes

import "testing"

//func maxEnvelopes(envelopes [][]int) int {
//
//}

func TestLopes(t *testing.T) {
}

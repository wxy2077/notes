/*
* @Time    : 2020-11-10 14:58
* @Author  : CoderCharm
* @File    : bubblesort.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package sorts

import (
	"testing"
)

func BubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] > arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}

	return arr
}

func TestBubbleSort(t *testing.T) {

	v := []int{5, 1, 3, 2, 4}

	t.Log(BubbleSort(v))
}

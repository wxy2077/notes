/*
* @Time    : 2021-03-08 20:47
* @Author  : CoderCharm
* @File    : quick_sort_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

快排
**/
package sorts

import (
	"math/rand"
	"testing"
)

func QuickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = QuickSort(lowPart)
	highPart = QuickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

func TestQuickSort(t *testing.T) {
	l := []int{4, 2, 77, 23, 67, 23, 56, 9, 1, 0, 99}
	t.Log(QuickSort(l))
}

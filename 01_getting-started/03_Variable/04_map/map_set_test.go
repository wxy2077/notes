/*
* @Time    : 2020-12-26 21:31
* @Author  : CoderCharm
* @File    : ch8_map_set.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _04_map

import (
	"testing"
)

func TestMapSet(t *testing.T) {
	mySet := map[int]bool{}

	mySet[1] = true
	n := 1

	if mySet[n] {
		t.Log("n is existing")
	} else {
		t.Log("n not existing")
	}

	mySet[3] = false

	t.Log(len(mySet))

	delete(mySet, 1)

	if mySet[n] {
		t.Log("n is existing")
	} else {
		t.Log("n not existing")
	}

}

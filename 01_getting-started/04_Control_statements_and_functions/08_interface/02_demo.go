/*
* @Time    : 2020-08-26 21:16
* @Author  : CoderCharm
* @File    : 02_demo.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("错误1")
var err2 = errors.New("错误2")

func fibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, err1
	}
	if n > 100 {
		return nil, err2
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func main() {
	if v, err := fibonacci(1); err != nil {
		if err == err1 {
			fmt.Println("是错误1")
		}
		fmt.Println("出错了", err)
	} else {
		fmt.Println(v)
	}
}

/*
* @Time    : 2021-02-21 19:44
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

cd 当前文件目录下
go run main.go > cpu.pprof


// web界面查看
go tool pprof -http=:9999 cpu.pprof


// 交互界面查看
go tool pprof cpu.pprof

> help // 查看所有命令


**/
package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

// 随机生成 整型数组
func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// 冒泡 排序
func bubbleSort(arr []int) []int {
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

func main() {
	_ = pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
}

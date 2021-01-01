/*
* @Time    : 2020-12-30 14:19
* @Author  : CoderCharm
* @File    : once_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

只运行一次, 类似单例模式

var once sync.Once
var obj *SingletonObj

func GetSingletonObj() *SingletonObj {

	once.Do(func(){
		fmt.Println("Create Singleton obj.")
		obj = &SingletonObj{}
	})

	return obj
}

**/
package _3_channel_context

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 新建一个结构体用于测试 查看是否地址 是否
type Singleton struct {
}

var singleInstance *Singleton

type ObjItem struct {
}

var genObjItem *ObjItem

var once sync.Once

func GetSingletonObj(i int) *Singleton {
	fmt.Printf("单例程序执行 %d----\n", i)

	once.Do(func() {
		// 这里面只执行一次
		fmt.Println("Create Singleton obj.")
		singleInstance = new(Singleton)
	})

	return singleInstance
}

func GenObj(i int) *ObjItem {
	fmt.Printf("生成程序执行 %d----\n", i)
	genObjItem = new(ObjItem)
	return genObjItem
}

func NormalFunc(i int) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf(" %d 测试函数 %s \n", i, timeStr)
}

func SingleFunc(i int) {
	fmt.Printf("单例测试函数执行++ %d \n", i)

	once.Do(func() {
		// 这里面只执行一次
		timeStr := time.Now().Format("2006-01-02 15:04:05")

		fmt.Printf("%d------单例子测试函数 只执行一次 %s \n", i, timeStr)
		//singleInstance = new(Singleton)
	})
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			//obj := GetSingletonObj(i)
			//fmt.Printf("单例:  %d++++%x\n", i, unsafe.Pointer(obj))

			NormalFunc(i)
			SingleFunc(i)

			wg.Done()
		}(i)
	}
	wg.Wait()
}

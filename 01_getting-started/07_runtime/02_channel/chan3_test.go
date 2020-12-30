/*
* @Time    : 2020-12-29 14:38
* @Author  : CoderCharm
* @File    : chan3_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

channel select 多channel操作

与case的顺序无关

**/
package _2_channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func AsyncCh1(n int, c chan string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		c <- fmt.Sprintf("++ %d ++", i)
	}
	wg.Done()
}

func AsyncCh2(n int, c chan string, wg *sync.WaitGroup) {
	for i := n; i > 0; i-- {
		c <- fmt.Sprintf("-- %d --", i)
	}
	wg.Done()
}

func TestSelect(t *testing.T) {
	var wg sync.WaitGroup

	ch1 := make(chan string, 5)
	ch2 := make(chan string, 5)

	wg.Add(2)
	go AsyncCh1(cap(ch1), ch1, &wg)
	//wg.Add(1)
	go AsyncCh2(cap(ch2), ch2, &wg)

	// 等待协程全部执行完
	wg.Wait()

	for i := 0; i < 10; i++ {
		select {
		case ret1 := <-ch1:
			t.Log(ret1)
		case ret2 := <-ch2:
			t.Log(ret2)
		case <-time.After(time.Millisecond * 100):
			t.Error("time out")
		}
	}
}

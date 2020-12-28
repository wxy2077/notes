/*
* @Time    : 2020-12-28 23:10
* @Author  : CoderCharm
* @File    : counter_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _1_goroutine

import (
	"sync"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	counter := 0

	for i := 0; i < 5000; i++ {
		go func() {

			counter++
		}()
	}

	time.Sleep(time.Second * 1)
	t.Log(counter)
}

func TestLock(t *testing.T) {
	// 声明一个锁
	var m sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {

			// 最后都要解锁
			defer func() {
				m.Unlock()
			}()

			// 先锁住
			m.Lock()
			counter++
		}()
	}

	// sleep 1 秒时间随机的  让协程执行完
	time.Sleep(time.Second * 1)
	t.Log(counter)
}

func TestWait(t *testing.T) {
	// 声明一个锁
	var m sync.Mutex
	var wg sync.WaitGroup

	counter := 0
	for i := 0; i < 5000; i++ {

		// 添加一个等待
		wg.Add(1)

		go func() {

			// 最后都要解锁
			defer func() {
				m.Unlock()
			}()

			// 先锁住
			m.Lock()
			counter++
			// 一个结束
			wg.Done()
		}()
	}
	// 等待协程全部执行完
	wg.Wait()

	t.Log(counter)
}

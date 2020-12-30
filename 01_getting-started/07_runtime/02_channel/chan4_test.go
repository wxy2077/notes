/*
* @Time    : 2020-12-29 23:20
* @Author  : CoderCharm
* @File    : chan4_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

生产者 消费者
**/
package _2_channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 不断的生产
func Producer1(c chan int, wg *sync.WaitGroup) {

	for i := 0; i <= 10; i++ {
		fmt.Println(fmt.Sprintf("生产了++++++ %d", i))
		time.Sleep(time.Millisecond * 100)
		c <- i
	}

	// 关闭channel
	close(c)
	//c <- 22  // 关闭后就不能发了  panic: send on closed channel

	wg.Done()
}

// 不断的从chanel里面拿
func Consumer(c chan int, wg *sync.WaitGroup) {

	for {
		//time.Sleep(time.Millisecond*800)
		// 判断生产者是否已经停止了
		v, ok := <-c
		if ok {
			fmt.Println(fmt.Sprintf("-------消费了 %d", v))
		} else {
			fmt.Println("结束")
			break
		}
	}
	wg.Done()

}

func TestChan4(t *testing.T) {
	//
	var wg sync.WaitGroup

	c := make(chan int, 20)

	wg.Add(2)

	go Producer1(c, &wg)
	go Consumer(c, &wg)

	wg.Wait()

}

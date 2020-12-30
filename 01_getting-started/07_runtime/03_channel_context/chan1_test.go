/*
* @Time    : 2020-12-30 10:25
* @Author  : CoderCharm
* @File    : chan1_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _3_channel_context

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

//func cancel_1(cancelChan chan struct{}) {
//	cancelChan <- struct {
//
//	}{}
//}

func cancel2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {

	cancelChan := make(chan struct{}, 0)

	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}

	//cancel_1(cancelChan)  //只会有一个任务被取消
	cancel2(cancelChan)

	time.Sleep(time.Second * 1)
}

//func TestChan1(t *testing.T){
//
//}

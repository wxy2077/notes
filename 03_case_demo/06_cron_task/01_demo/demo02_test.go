/*
* @Time    : 2021-01-04 21:13
* @Author  : CoderCharm
* @File    : demo02_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _1_demo

import (
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func Test02Cron(t *testing.T) {
	c := cron.New()

	_, _ = c.AddFunc("* * * * *", jobTask)

	c.Start()

	time.Sleep(3 * time.Minute)

	//t1 := time.NewTimer(time.Second * 10)
	//
	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Second * 10)
	//	}
	//}
}

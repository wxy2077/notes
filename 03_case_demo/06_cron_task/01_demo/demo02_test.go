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
	// 设置时区
	nyc, _ := time.LoadLocation("Asia/Shanghai")

	c := cron.New(cron.WithLocation(nyc))

	_, _ = c.AddFunc("* * * * *", jobTask)

	// 定时运行一次 https://github.com/robfig/cron/pull/317  目前不支持
	//_, _ = c.AddFunc("@once 2021-01-09 10:45:00", func() {
	//	fmt.Println("Hello! Now is 2021-01-09 10:45:00")
	//})

	c.Start()

	//time.Sleep(3 * time.Minute)

	// 阻塞
	select {}

	//t1 := time.NewTimer(time.Second * 10)
	//
	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Second * 10)
	//	}
	//}
}

/*
* @Time    : 2021-01-04 20:31
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _1_demo

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func jobTask() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func TestCron(t *testing.T) {
	c := cron.New()

	_, _ = c.AddFunc("@every 3s", jobTask)

	c.Start()
	time.Sleep(time.Second * 9)
}

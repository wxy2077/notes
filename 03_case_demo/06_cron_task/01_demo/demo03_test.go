/*
* @Time    : 2021-01-05 21:41
* @Author  : CoderCharm
* @File    : demo03_test.go
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

type testJob struct {
	Name string
}

//实现了 type Job interface {Run()} func名字必须为Run
func (t *testJob) Run() {
	fmt.Println("test job 参数为", t.Name)
}

func Test03Cron(t *testing.T) {

	// 设置时区
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	// 设置时区，  支持秒级调度
	GlobalCron := cron.New(cron.WithLocation(nyc), cron.WithSeconds())

	// 调用第一个
	entryId01, _ := GlobalCron.AddFunc("@every 1s", jobTask)
	fmt.Printf("第一个id %v \n", entryId01)

	// 执行第二个
	entryId02, _ := GlobalCron.AddJob("*/2 * * * * *", &testJob{"JobName"})
	fmt.Printf("第二个id %v \n", entryId02)

	GlobalCron.Start()

	select {}
}

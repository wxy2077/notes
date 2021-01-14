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

// 定时任务
func jobTask() {
	fmt.Printf("任务启动: %s \n", time.Now().Format("2006-01-02 15:04:05"))
}

func TestCron(t *testing.T) {
	// 创建一个cron对象
	c := cron.New()

	// 任务调度
	enterId, err := c.AddFunc("@every 3s", jobTask)
	if err != nil {
		panic(err)
	}
	fmt.Printf("任务id是 %d \n", enterId)

	// 同步执行任务会阻塞当前执行顺序  一般使用Start()
	//c.Run()
	//fmt.Println("当前执行顺序.......")

	// goroutine 协程启动定时任务(看到后面Start函数和run()函数，就会明白启动这一步也可以写在任务调度之前执行)
	c.Start()
	// Start()内部有一个running 布尔值 限制只有一个Cron对象启动 所以这一步多个 c.Start() 也只会有一个运行
	c.Start()
	c.Start()

	// 用于阻塞 后面可以使用 select {} 阻塞
	time.Sleep(time.Second * 9)

	// 关闭定时任务(其实不关闭也可以，主进程直接结束了, 内部的goroutine协程也会自动结束)
	c.Stop()

}

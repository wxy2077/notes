/*
* @Time    : 2021-01-06 20:02
* @Author  : CoderCharm
* @File    : sys_cron.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

初始化 cron 定时任务

不支持持久化
不支持调度后立即运行 得手动 go func()  https://github.com/robfig/cron/issues/297

**/
package initialize

import (
	"github.com/robfig/cron/v3"
	"time"
)

// 初始化定时任务
func SysCron() *cron.Cron {

	// 时区北京时间
	nyc, _ := time.LoadLocation("Asia/Shanghai")

	// 支持秒级别任务
	return cron.New(
		cron.WithLocation(nyc),
		cron.WithSeconds(),
	)
}

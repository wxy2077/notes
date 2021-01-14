/*
* @Time    : 2021-01-07 19:23
* @Author  : CoderCharm
* @File    : sys_system.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package model

import "gin_study/config"

// 配置文件结构体
type System struct {
	Config config.Server
}

/*
* @Time    : 2020-12-31 10:27
* @Author  : CoderCharm
* @File    : system.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 读取系统配置
**/
package config

type System struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}

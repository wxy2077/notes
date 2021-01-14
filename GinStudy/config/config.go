/*
* @Time    : 2020-11-19 10:41
* @Author  : CoderCharm
* @File    : config.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 读取配置文件
**/
package config

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

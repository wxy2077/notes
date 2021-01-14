/*
* @Time    : 2021-01-08 20:27
* @Author  : CoderCharm
* @File    : jwt.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"`
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`
}

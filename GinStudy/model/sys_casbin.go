/*
* @Time    : 2021-01-07 19:20
* @Author  : CoderCharm
* @File    : sys_casbin.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package model

type CasbinModel struct {
	Ptype       string `json:"ptype" gorm:"column:p_type"`
	AuthorityId string `json:"rolename" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}

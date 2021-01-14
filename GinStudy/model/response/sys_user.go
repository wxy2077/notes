/*
* @Time    : 2021-01-08 19:15
* @Author  : CoderCharm
* @File    : sys_user.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package response

import "gin_study/model"

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

type LoginResponse struct {
	User      model.SysUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}

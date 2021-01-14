/*
* @Time    : 2021-01-07 19:36
* @Author  : CoderCharm
* @File    : sys_user.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package model

import (
	"gin_study/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.COMMON_MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`
	Username    string       `json:"userName" gorm:"comment:用户登录名"`
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`
	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg   string       `json:"headerImg" gorm:"default:https://image.3001.net/images/20200504/1588558613_5eaf7b159c8e9.jpeg;comment:用户头像"`
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}

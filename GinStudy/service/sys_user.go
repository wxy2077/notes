/*
* @Time    : 2021-01-07 20:25
* @Author  : CoderCharm
* @File    : sys_user.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package service

import (
	"errors"
	"gin_study/global"
	"gin_study/model"
	"gin_study/utils"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

func Register(u model.SysUser) (userInter model.SysUser, err error) {
	var user model.SysUser
	if !errors.Is(global.GIN_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GIN_DB.Create(&u).Error
	return u, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Login(u *model.SysUser) (userInter *model.SysUser, err error) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GIN_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return &user, err
}

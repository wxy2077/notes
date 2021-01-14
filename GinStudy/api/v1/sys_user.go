/*
* @Time    : 2021-01-07 17:00
* @Author  : CoderCharm
* @File    : sys_user.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package v1

import (
	"gin_study/global"
	"gin_study/middleware"
	"gin_study/model"
	"gin_study/model/request"
	"gin_study/model/response"
	"gin_study/service"
	"gin_study/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	registerInfo := request.Register{}
	_ = c.ShouldBindJSON(&registerInfo)
	if err := utils.Verify(registerInfo, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &model.SysUser{Username: registerInfo.Username, NickName: registerInfo.NickName,
		Password:    registerInfo.Password,
		HeaderImg:   registerInfo.HeaderImg,
		AuthorityId: registerInfo.AuthorityId}
	userReturn, err := service.Register(*user)
	if err != nil {
		global.GIN_LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	U := &model.SysUser{Username: L.Username, Password: L.Password}
	if user, err := service.Login(U); err != nil {
		global.GIN_LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		tokenNext(*user, c)
	}
}

// 登录以后签发jwt
func tokenNext(user model.SysUser, c *gin.Context) {
	j := &middleware.JWT{SigningKey: []byte(global.GIN_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
		BufferTime:  global.GIN_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GIN_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "wxy",                                                 // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GIN_LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GIN_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
}

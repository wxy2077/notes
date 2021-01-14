/*
* @Time    : 2021-01-08 20:58
* @Author  : CoderCharm
* @File    : sys_user.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package request

// User register structure
type Register struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'王小右'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'https://image.3001.net/images/20200504/1588558613_5eaf7b159c8e9.jpeg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

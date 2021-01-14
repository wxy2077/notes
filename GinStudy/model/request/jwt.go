/*
* @Time    : 2021-01-08 20:30
* @Author  : CoderCharm
* @File    : jwt.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package request

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}

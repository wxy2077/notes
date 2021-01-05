/*
* @Time    : 2021-01-05 21:24
* @Author  : CoderCharm
* @File    : token_summary_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

汇总生成 和 解析
**/
package _7_jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

type CustomClaims struct {
	ID       uint
	Username string
	jwt.StandardClaims
}

// 生成token
func generateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Key))
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

// 解析token
func parseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(Key), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

func TestToken(t *testing.T) {
	//fmt.Println(123)
	//claims := &CustomClaims{ID:123, Username:"222"}
	//claims.ExpiresAt = time.Now().Unix() + 3600
	//token, err := generateToken(*claims)
	//fmt.Println(token, err)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MTIzLCJVc2VybmFtZSI6IjIyMiIsImV4cCI6MTYwOTg1ODE4OH0.VCLoPQjW9DjAyi_NwQXmC2CVnevZ0DB-wq-wHyxIl9c"
	claims, err := parseToken(token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*claims)
	}

}

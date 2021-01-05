/*
* @Time    : 2021-01-05 20:47
* @Author  : CoderCharm
* @File    : generate_token_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

https://godoc.org/github.com/dgrijalva/jwt-go#example-New--Hmac

生成Token
**/
package _7_jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

const Key = "kkkkk"

func TestGenerateToken(t *testing.T) {

	//n := time.Now().Unix()
	//
	//fmt.Println(n)

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		//"any_key": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"any_key": time.Now().Unix(),
	})
	//Key := "kkkkk"

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(Key))
	if err != nil {
		fmt.Printf("错误: %s \n", err)
	}

	fmt.Println(tokenString)
}

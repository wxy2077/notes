/*
* @Time    : 2021-01-05 20:58
* @Author  : CoderCharm
* @File    : parse_token_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac

**/
package _7_jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func TestParseToken(t *testing.T) {
	// sample token string taken from the New example
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhbnlfa2V5IjoxNjA5ODUxOTAzLCJmb28iOiJiYXIifQ.Bl5hme75RSGTnOhsdYBPjrICMnepYMJoc__AWhRK9bU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(Key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}
}

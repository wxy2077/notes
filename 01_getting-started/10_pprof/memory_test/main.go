/*
* @Time    : 2021-02-20 22:05
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

go run main.go

会生成一个文件 mem.pprof  路径会有所变化

go tool pprof -http=:9999 /var/folders/x8/w3zpmkks46s1v0shvtkn_6zc0000gn/T/profile718597885/mem.pprof


**/
package main

import (
	"github.com/pkg/profile"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	concat(100)
}

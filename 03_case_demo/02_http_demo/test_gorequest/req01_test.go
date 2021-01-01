/*
* @Time    : 2020-12-31 11:24
* @Author  : CoderCharm
* @File    : req01_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package test_gorequest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"testing"
)

func StringToMap(content string) (map[string]interface{}, error) {
	var resMap map[string]interface{}
	err := json.Unmarshal([]byte(content), &resMap)
	if err != nil {
		return nil, err
	}
	return resMap, nil
}

var UserAgent = "this is custom user agent"

func GetRequest(targetUrl string, queryMap map[string]string) (content string, err error) {
	resp, body, errs := gorequest.New().Get(targetUrl).Query(queryMap).
		Set("User-Agent", UserAgent).
		End()
	if errs != nil {
		return "", errors.New(fmt.Sprintf("%s 请求错误 %s", targetUrl, errs))
	}
	// 可以使用 _ 接收 resp
	fmt.Printf("url:%s 请求状态 %s \n", targetUrl, resp.Status)

	return body, nil
}

type UserInfo struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func PostRequest(targetUrl string) (body string, err error) {

	// 设置post json数据  也开始使用下面的 Send(`{"user": "aaa"}`)
	jsonData := UserInfo{Name: "Nick", Age: "25"}

	resp, body, errs := gorequest.New().Post(targetUrl).Send(jsonData).Query(map[string]string{"go": "go_on_learn"}).
		Set("User-Agent", UserAgent).
		Send(`{"user": "aaa"}`).Type("multipart").
		End()

	fmt.Printf("url:%s 请求状态 %s \n", targetUrl, resp.Status)

	if errs != nil {
		return "", errors.New(fmt.Sprintf("%s 请求错误 %s", targetUrl, errs))
	}
	return body, nil

}

func Test01(t *testing.T) {
	getUrl := "https://httpbin.org/get"

	// 查询参数
	queryMap := map[string]string{"Python": "learn"}
	resStr, err := GetRequest(getUrl, queryMap)
	if err != nil {
		panic(fmt.Sprintf("请求错误 %s", err))
	}

	fmt.Printf(resStr)

	fmt.Printf("\n\n\n------------------------\n\n\n")

	postUrl := "https://httpbin.org/post"
	resPostStr, err := PostRequest(postUrl)

	if err != nil {
		panic(fmt.Sprintf("请求错误 %s", err))
	}

	fmt.Printf(resPostStr)

}

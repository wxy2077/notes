/*
* @Time    : 2020-12-27 15:11
* @Author  : CoderCharm
* @File    : req3_post_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

post form 请求
**/
package http_demo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func PostReq(targetUrl string) (content string, err error) {
	// 最简单的发送 post请求 携带 form参数
	resp, err := http.PostForm(targetUrl, url.Values{"username": {"wxy"}, "password": {"wxy666"}})

	// 相当于这种
	//req.Header.Set("Content-Type", "multipart/form-data")

	if err != nil {
		return "", errors.New(fmt.Sprintf("请求错误 %s", err))
	}
	fmt.Println(fmt.Sprintf("响应状态 %s", resp.Status))

	// 记得关闭客户端
	defer resp.Body.Close()
	// 读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("读取内容错误 %s", err))
	}
	content = string(body)
	return content, nil

}

func TestReq3(t *testing.T) {
	targetUrl := "https://httpbin.org/post"

	content, err := PostReq(targetUrl)
	if err != nil {
		fmt.Println("请求错误", err)
		return
	}
	fmt.Println(content)
}

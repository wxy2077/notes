/*
* @Time    : 2020-12-27 19:03
* @Author  : CoderCharm
* @File    : req4_post_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package http_demo

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

// post form请求 务必设置  Content-Type
func ReqPostForm(targetUrl string) (content string, err error) {
	// 创建 http 客户端
	client := &http.Client{}

	form := url.Values{}
	form.Add("ln", "ln222")
	form.Add("ip", "1.1.1.1")
	form.Add("ua", "ua123")

	// 创建请求
	req, _ := http.NewRequest("POST", targetUrl, strings.NewReader(form.Encode()))

	req.Header.Set("User-Agent", "test")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	resp, err := client.Do(req)

	if err != nil {
		// 格式化返回错误
		return "", errors.New(fmt.Sprintf("请求出错 %s", err))
	}

	// 最后关闭连接
	defer resp.Body.Close()

	// 读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("解析内容出错 %s", err))
	}
	return string(body), nil
}

// post json方式请求
func ReqPostJson(targetUrl string) (content string, err error) {

	var jsonStr = []byte(`{"title":"this is a title", "cate": 1}`)

	req, _ := http.NewRequest("POST", targetUrl, bytes.NewBuffer(jsonStr))

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return "", errors.New(fmt.Sprintf("请求出错 %s", err))
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil

}

func TestReqPost(t *testing.T) {
	targetUrl := "https://httpbin.org/post"

	c1, _ := ReqPostForm(targetUrl)
	t.Log(fmt.Sprintf("\nfrom表单返回: \n %s", c1))

	c2, _ := ReqPostJson(targetUrl)
	t.Log(fmt.Sprintf("\nJson请求返回: \n %s", c2))

}

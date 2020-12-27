/*
* @Time    : 2020-12-26 21:52
* @Author  : CoderCharm
* @File    : req1_get_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

最基本的发送请求解析内容

req2 比较完善

**/
package http_demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func GetHtml(url string) {

	// 构建http请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	fmt.Println(fmt.Sprintf("响应状态 %s", resp.Status))

	// 记得关闭客户端
	defer resp.Body.Close()

	// 读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	fmt.Println(string(body))
}

func TestReq(t *testing.T) {
	url := "https://www.charmcode.cn/"
	GetHtml(url)
}

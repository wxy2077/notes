/*
* @Time    : 2020-12-26 22:49
* @Author  : CoderCharm
* @File    : demo2.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

添加请求 发送请求 然后获取 解析json数据

**/
package http_demo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Requests(url string) (content string, err error) {
	// 创建 http 客户端
	client := &http.Client{}

	// 创建请求
	req, _ := http.NewRequest("GET", url, nil)

	// GET 请求携带查询参数
	q := req.URL.Query()
	q.Add("auth_key", "sky_kk")
	q.Add("another_thing", "foo & bar")
	req.URL.RawQuery = q.Encode()

	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	// 设置请求头
	req.Header.Set("User-Agent", "test")

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
	//fmt.Println(string(body))
	return string(body), nil
}

func StringToMap(content string) map[string]interface{} {
	var resMap map[string]interface{}
	err := json.Unmarshal([]byte(content), &resMap)
	if err != nil {
		fmt.Println("string转map失败", err)
	}
	return resMap
}

func TestReqUrl(t *testing.T) {
	url := "https://httpbin.org/get"
	content, err := Requests(url)
	if err != nil {
		fmt.Println("请求错误", err)
		return
	}
	t.Log(fmt.Sprintf("结果 \n %s", content))
	jsonContent := StringToMap(content)

	t.Log(jsonContent)
}

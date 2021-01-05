/*
* @Time    : 2020-12-31 11:46
* @Author  : CoderCharm
* @File    : tools.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 常见工具类
**/
package utils

import "encoding/json"

// string 转 map
func StringToMap(content string) (map[string]interface{}, error) {
	var resMap map[string]interface{}
	err := json.Unmarshal([]byte(content), &resMap)
	if err != nil {
		return nil, err
	}
	return resMap, nil
}

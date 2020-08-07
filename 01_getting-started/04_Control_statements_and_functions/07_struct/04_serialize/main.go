/*
结构体的序列化 和反序列化

*/
package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	// 大写公有原则 小写私有
	Name  string `json:"Name"` // 序列化时取别名
	Color string `json:"Color"`
}

func main() {
	// 实力化一个结构体
	Tom := Cat{Name: "Tom", Color: "grey"}
	// 序列化
	v, err := json.Marshal(&Tom)
	if err != nil {
		fmt.Println("序列化出错了")
	}
	fmt.Println(string(v))

	var t Cat
	// json格式到字符串
	jsonRes := `{"Name": "HaHa","Color": "qqq"}`
	// 反序列化
	err = json.Unmarshal([]byte(jsonRes), &t)
	if err != nil {
		fmt.Println("反序列化出错了")
	}
	fmt.Println(t)
}

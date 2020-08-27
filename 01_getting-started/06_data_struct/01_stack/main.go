/*
* @Time    : 2020-08-27 22:42
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 用slice实现的最基本的栈结构
**/
package main

import (
	"encoding/json"
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) push(n int) {
	// 压入栈
	s.items = append(s.items, n)
}

func (s *Stack) pop() int {
	// 弹出栈 没有判断范围
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value
}

func (s *Stack) peek() int {
	// 查看栈顶元素
	return s.items[len(s.items)-1]
}

func (s *Stack) isEmpty() bool {
	// 判断是否为空
	return len(s.items) == 0
}

func (s *Stack) size() int {
	// 查看栈长度
	return len(s.items)
}

func (s *Stack) toString() string {
	// 返回栈字符串
	v, _ := json.Marshal(s.items)
	return string(v)
}

func main() {
	var s Stack
	s.push(1)
	fmt.Println("压入栈", s.items)
	s.push(2)
	fmt.Println("压入栈", s.items)
	s.push(3)
	fmt.Println("压入栈", s.items)
	fmt.Println("toString返回字符形式", s.toString())
	fmt.Println("peek查看栈顶元素", s.peek())
	fmt.Println("isEmpty查看栈是否为空", s.isEmpty())
	fmt.Println("size查看栈个数", s.size())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.items)
	fmt.Println("isEmpty查看栈是否为空", s.isEmpty())
}

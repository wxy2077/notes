/*
* @Time    : 2021-03-05 21:15
* @Author  : CoderCharm
* @File    : double_stack_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

仅使用两个栈实现先入先出队列

将一个栈当作输入栈，用于push数据 一个当作输出栈，用于pop和peek


**/
package _32_queue_use_stack

import "testing"

type MyQueue struct {
	inStack, outStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (m *MyQueue) Push(x int) {
	m.inStack = append(m.inStack, x)
}

/**
一次把 inStack 的数据压入到 outStack里面
*/
func (m *MyQueue) in2OutStack() {
	for len(m.inStack) > 0 {
		// 最后一个元素依次压入到输出栈
		m.outStack = append(m.outStack, m.inStack[len(m.inStack)-1])
		// 输入栈 去掉 压入到输出栈里的数据
		m.inStack = m.inStack[:len(m.inStack)-1]
	}
}

/** Removes the element from in front of queue and returns that element. */
func (m *MyQueue) Pop() int {
	if len(m.outStack) == 0 {
		m.in2OutStack()
	}
	n := m.outStack[len(m.outStack)-1]
	m.outStack = m.outStack[:len(m.outStack)-1]
	return n
}

/** Get the front element. */
func (m *MyQueue) Peek() int {
	if len(m.outStack) == 0 {
		m.in2OutStack()
	}
	// 输出栈  最后的一个 元素即为 最前面的元素
	return m.outStack[len(m.outStack)-1]
}

///** Returns whether the queue is empty. */
func (m *MyQueue) Empty() bool {
	return len(m.inStack) == 0 && len(m.outStack) == 0
}

func TestQueueStack(t *testing.T) {
	obj := Constructor()
	obj.Push(3)
	obj.Push(6)
	obj.Push(8)

	t.Log(obj.inStack)
	t.Log(obj.outStack)

	t.Log(obj.Peek())

	t.Log(obj.Pop())
	t.Log(obj.Pop())
	t.Log(obj.Pop())

	t.Log(obj.inStack)
	t.Log(obj.outStack)

	t.Log(obj.Empty())
}

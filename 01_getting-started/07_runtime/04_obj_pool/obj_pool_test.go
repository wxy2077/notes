/*
* @Time    : 2020-12-30 21:59
* @Author  : CoderCharm
* @File    : obj_pool_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package _4_obj_pool

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// 新建一个空结构体 相当于对象
type Tool struct {
	name string
}

// 对象池 用于存储 Tool对象
type ToolsBox struct {
	// 属性是一个 channel 内容是 Tool 结构体指针
	bufChan chan *Tool
}

// 获取工具 给结构体绑定方法
func (p *ToolsBox) GetTool(timeout time.Duration) (*Tool, error) {
	select {
	case tool := <-p.bufChan:
		return tool, nil
	case <-time.After(timeout): //超时控制
		return nil, errors.New("time out")
	}
}

// 用完归还(释放)
func (p *ToolsBox) ReleaseTool(tool *Tool) error {
	select {
	case p.bufChan <- tool:
		return nil
	default:
		return errors.New("overflow")
	}
}

// new一个 ToolBox对象
func NewToolsBox(poolNum int) *ToolsBox {
	objPool := ToolsBox{}
	objPool.bufChan = make(chan *Tool, poolNum)

	for i := 0; i < poolNum; i++ {

		// 生成一个 工具结构体
		tool := &Tool{fmt.Sprintf("🔧--%d", i)}
		// 存入对象池
		objPool.bufChan <- tool
	}

	return &objPool

}

func TestObjPool(t *testing.T) {

	pool := NewToolsBox(5)

	//tool,_ := pool.GetTool(time.Second * 1)
	//t.Log(fmt.Sprintf("取出一个当前容量%d", len(pool.bufChan)))
	//t.Log(tool)
	//
	//_ = pool.ReleaseTool(tool)
	//t.Log(fmt.Sprintf("归还后当前容量%d", len(pool.bufChan)))

	for i := 0; i < 8; i++ {
		tool, err := pool.GetTool(time.Second * 1)

		if err != nil {
			t.Log(fmt.Sprintf("---取出有问题 %s 当前容量%d", tool, len(pool.bufChan)))
		} else {
			// 取出没问题
			t.Log(fmt.Sprintf("----取出一个 %s 当前容量%d", tool, len(pool.bufChan)))

			// 接着就释放 和判断写在一起
			if err := pool.ReleaseTool(tool); err != nil {
				t.Log("释放有问题")
			} else {
				t.Log(fmt.Sprintf("释放一个 +++ %s 当前容量%d", tool, len(pool.bufChan)))
			}
		}

	}

	t.Log("结束")
}

/*
* @Time    : 2021-01-19 21:13
* @Author  : CoderCharm
* @File    : 01_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

go get github.com/casbin/gorm-adapter/v3

**/
package _2_casbin_gorm

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"testing"
)

// 动态添加角色权限
func TestCasbin(t *testing.T) {

	a, _ := gormadapter.NewAdapter("mysql", "root:Admin12345-@tcp(172.16.137.129:3306)/go_casbin", true)

	e, err := casbin.NewEnforcer("./model.conf", a)

	if err != nil {
		panic(err)
	}

	// 查询
	//res := e.GetAllSubjects()
	//res := e.GetAllObjects()
	//res := e.GetAllActions()
	//res := e.GetPolicy()
	res := e.GetFilteredPolicy(0, "10")
	t.Log(res)

	// 动态添加策略 自己定义比如 权限id  接口  请求
	//res, err := e.AddPolicy("20", "/api/delUser", "DELETE")
	// AddNamedPolicies 添加多个

	// 删除策略
	//res, err := e.RemovePolicy("10", "/api/getUserList", "GET")
	// RemovePolicies 删除多个

	// 更新策略 UpdatePolicy 此方法暂时未实现
	//res, err := e.UpdatePolicy([]string{"10", "/api/getUserList", "GET"}, []string{"100", "/api/getUserList", "GET"})

	//if err != nil{
	//	panic(err)
	//}else{
	//	t.Log(res)
	//}

	//sub := "Jack"  // the user that wants to access a resource.
	//obj := "data1" // the resource that is going to be accessed.
	//act := "read"  // the operation that the user performs on the resource.

	// 验证
	//ok, err := e.Enforce(sub, obj, act)
	//
	//if err != nil {
	//	// handle err
	//	t.Log(fmt.Sprintf("错误 %v", err))
	//}
	//
	//if ok == true {
	//	// permit alice to read data1
	//	t.Log("ok")
	//} else {
	//	// deny the request, show an error
	//	t.Log("fail")
	//}

}

# go 语言学习笔记

## 笔记说明
> 练习代码，都以编写测试程序的方式书写，注意`go`以文件夹为包单位，同一文件夹下包名一样。

- 1 源码文件以`_test.go`结尾， 如`str_test.go`
- 2 测试方法名以大写`Test`开头，如`func TestXXX(t *testing.T){t.Log(111)}`

如以下代码结构

```go
// 在一个空文件夹下新建一个`hello_test.go`文件 
package hello

import "testing"

func TestString(t *testing.T){
	t.Log("hello world")
}
```

## 目录 

### 1 入门学习
- [1 hello world](01_getting-started/01_HelloWorld/README.md)

### 2 基础算法学习


### 3 一些案例框架练习 
- [2 go http请求库的使用](03_case_demo/02_http_demo/README.md)
- [3 go mysql 和 gorm的基本使用(使用Viper读取yaml配置文件)](03_case_demo/)
- [4 go 七牛云上传在线视频](03_case_demo/04_qiniu_upload)
- [5 gRPC学习使用](03_case_demo/05_go_grpc/README.md)
- [6 cron定时任务基本使用](03_case_demo/06_cron_task/README.md)
- [7 JWT token的基本使用](03_case_demo/07_jwt/README.md)
- [8 casbin权限管理demo](03_case_demo/08_casbin/README.md)

### 参考

- [他人学习笔记1](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md)
- [github.com/inancgumus/learngo](https://github.com/inancgumus/learngo)
- [极客go学习笔记](https://github.com/CoderCharm/gostudy)
- [go语言程序设计](https://docs.hacknode.org/gopl-zh/index.html)
- [go web编程在线书籍](https://learnku.com/docs/build-web-application-with-golang)
- [go在线练习](https://tour.golang.org/list)
- [https://golang.google.cn/](https://golang.google.cn/)
- [JSON解析](https://stackoverflow.com/questions/35583735/unmarshaling-into-an-interface-and-then-performing-type-assertion)
- [Go语言设计与实现](https://draveness.me/golang/)
- [Go语言高级编程](https://hezhiqiang8909.gitbook.io/go/)
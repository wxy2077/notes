# GoLang 入门学习

## `go`环境变量

`go env GOPATH ` 这个命令可以有效输出当前使用go的路径。
如果没有设置环境变量，就回打印默认安装位置

一般为了更方便，可以把工作区的`/bin`子目录添加到路径中,如下:
`export PATH=$PATH:$(go env GOPATH)/bin`

-----
## 第一个程序
> https://golang.google.cn/doc/code.html#Overview

1 第一步，创建一个 `hello.go` 文件

```go
package main  // 定义包名
// You must statement  the package name in the code first line (not notes) of the source files, e.g. package main,
// package main indicates a program that independent and enforceable, each go application should be included a package called main.
                                           
            
import "fmt"  // The full name of fmt package is format, which implements the function of formatting IO (input / output).

func main() {
	fmt.Println("Hello, world.") // Line feed printing
                                    
}
```


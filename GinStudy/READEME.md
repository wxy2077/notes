
# Golang Web框架 Gin学习笔记

参考 
- https://goproxy.io/zh/
- Go Modules 官网 https://blog.golang.org/using-go-modules

## 环境配置 Go 版本是 1.13 及以上 （推荐）
```
go env -w GO111MODULE=auto
go env -w GOPROXY=https://goproxy.io,direct
```

### ide goland 配置
```
setting->Go->Go Modules->Proxy:https://goproxy.io
```
Enable Go Mod 打勾


### 初始化
```
cd 项目文件夹            // 进入项目文件夹下
go mod init gin-stydy  // gin-stydy为项目名(随意取)

// 就会创建go.mod文件
```


## GO Modules常见命令

```
go list // 列出主模块(当前模块)
```

```
go list -m all // 列出了当前模块及其所有依赖项
```


删除未使用的依赖项
```
go mod tidy
```

github 拉取 他人含有go.mod的项目时,下载所有第三方包
```
go mod download
```

验证依赖是否正确

```
go mod verify
```


升级库依赖, 比如Gin框架 升级
```
go get -u github.com/gin-gonic/gin
```

安装依赖指定版本
```
go get github.com/gin-gonic/gin@version
```

安装依赖指指定分支
```
go get github.com/gin-gonic/gin@master
```
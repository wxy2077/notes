
# Golang Web框架 Gin学习笔记

参考 
- https://goproxy.io/zh/
- Go Modules 官网 https://blog.golang.org/using-go-modules

> 由于是练习，所有Go笔记代码都写在一个仓库里面，由于使用goland IDE，所以把`GinStudy`标记为根目录

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


## 项目目录说明

<details>
<summary>目录</summary>

```
.
|____api                接口目录
| |____v1                   v1
| | |____article.go             文章接口

|____config             配置文件
| |____config.go            汇聚配置文件结构体
| |____zap.go               定义读取配置文件yaml格式
| |____gorm.go              定义读取数据库配置文件yaml格式

|____core               核心的基础服务配置
| |____server.go            通过全局路由 启动http服务
| |____zap.go               日志文件配置
| |____viper.go             读取配置文件

|____global             定义全局变量
| |____global.go            

|____initialize         各种初始化文件
| |____router.go            汇聚各模块路由
| |____gorm.go              定义`gorm`数据库配置

|____middleware         定义中间价
| |____cors.go              跨域中间件
| |____zaplogger.go         记录日志中间价

|____model              定义各种model
| |____request              请求参数model
| | |____common.go              通用请求参数
| |____response             响应参数model
| | |____common.go              通用响应参数
| | |____response.go    
| |____article.go           文章表model

|____router             路由分组(配置中间价)
| |____article.go           文章model路由

|____service            表操作逻辑
| |____article.go           文章表的curd操作

|____utils              通用的工具方法
| |____rotatelogs_unix.go
| |____directory.go
| |____constant.go

|____config.yaml        配置文件
|____go.mod
|____go.sum
|____main.go            启动文件
|____README.md

```

</details>

### 初始化
```
cd 项目文件夹            // 进入项目文件夹下
go mod init gin_stydy  // gin_stydy为项目名(随意取)

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
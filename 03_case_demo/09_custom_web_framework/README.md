# 自定义web框架

## 目标
构建一个微型的`gin`框架，拥有基本的功能，如路由分组，中间件等，异常恢复。

最终使用目录结构。
```
gee/            // 自定义gee框架
  |-- gee.go    // 核心文件 封装net/http
  |-- xxx.go    // 其他扩展文件
  |-- ...
  
main.go         // 用户web服务代码
go.mod
```

## 目的

学习web框架的构建原理，方便使用Go web框架。


## 参考:
- https://geektutu.com/post/gee-day1.html
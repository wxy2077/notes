# pprof 性能分析


安装 graphviz

```
内置的pprof
https://github.com/google/pprof
```

```

go run main.go > cpu.pprof


go tool pprof -http=:9999 cpu.pprof

// 交互模式查看
go tool pprof cpu.pprof
```


https://geektutu.com/post/hpg-pprof.html
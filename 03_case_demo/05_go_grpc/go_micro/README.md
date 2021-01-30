
# go-micro demo


## 声明

此案例运行失败

安装v3版本
```shell
go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3
```


## 官方示例

https://github.com/asim/go-micro/tree/master/cmd/protoc-gen-micro


生成其他文件
```shell
// 先安装
$ go get github.com/golang/protobuf/protoc-gen-go

// 自己手动生成
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto

```

慕课网教程  xxxx为慕课网购买课程后的校验码
```shell
sudo docker run --rm -v $(pwd):$(pwd) -w $(pwd) -e ICODE=xxxxx cap1573/cap-protoc -I ./ --go_out=./ --micro_out=./ ./cap/imooc.proto
```
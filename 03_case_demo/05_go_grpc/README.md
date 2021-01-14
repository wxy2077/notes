# gRPC 使用案例

> "Definition - gRPC is a modern, open source remote procedure call (RPC) framework that can run anywhere"
> gRPC 使用了google的`protobuf`消息格式，而不是传统的`JSON`， 相对`JSON`，`protobuf`具有速度快，序列化效率高，天然的加密(得有.proto文件)

目前个人只是看看demo，还没有场景需要用到这个。目前还是REST API使用更多一些, 也可以构建 RESTful 服务发送`protobuf`消息。

## 依赖安装

```shell
go get -u google.golang.org/grpc
```


## gRPC 和 REST的区别

- gRPC 使用 HTTP/2 而 REST 使用 HTTP 1.1.(参考[StackOverflow](https://stackoverflow.com/questions/43682366/how-is-grpc-different-from-rest)）
- gRPC 使用 protocol buffer 数据格式，而不是REST中常常使用的JSON数据格式。
- 可以使用 gRPC HTTP/2 的功能，例如 服务器到客户端流, 客户端到服务器流 或者 双向流媒体.


## gRPC REST GraphQL Web

## 使用场景


# 参考

- [官方文档](https://grpc.io/docs/languages/go/quickstart/)
- [官方示例](https://github.com/grpc/grpc-go/tree/master/examples/helloworld)
- [https://tutorialedge.net/golang/go-grpc-beginners-tutorial/](https://tutorialedge.net/golang/go-grpc-beginners-tutorial/)
- [https://nordicapis.com/when-to-use-what-rest-graphql-webhooks-grpc/](https://nordicapis.com/when-to-use-what-rest-graphql-webhooks-grpc/)
- [https://docs.microsoft.com/en-us/aspnet/core/grpc/comparison?view=aspnetcore-5.0](https://docs.microsoft.com/en-us/aspnet/core/grpc/comparison?view=aspnetcore-5.0)
- [https://stackoverflow.com/questions/43682366/how-is-grpc-different-from-rest](https://stackoverflow.com/questions/43682366/how-is-grpc-different-from-rest)
- [http2 对比 websocket](https://stackoverflow.com/questions/28582935/does-http-2-make-websockets-obsolete)

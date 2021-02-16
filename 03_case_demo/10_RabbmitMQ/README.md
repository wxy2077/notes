## Golang 操作 RabbitMQ


## docker 安装 RabbitMQ

https://hub.docker.com/_/rabbitmq

安装tag 带有 `management`也就是带有web管理界面的

拉取镜像
```
docker pull rabbitmq:3.8-management
```
启动
```
docker run --name rabbitmq -d -p 15672:15672 -p 5672:5672 rabbitmq:3.8-management
```
```
--name指定了容器名称
-d 指定容器以后台守护进程方式运行
-p指定容器内部端口号与宿主机之间的映射，rabbitMq默认要使用15672为其web端界面访问时端口，5672为数据通信端口
```

然后会有默认的账号访问
```
http://127.0.0.1:15672
账号: guest
密码: guest
```

## 设置新的用户

查看容器运行 找到运行容器id
```
docker ps
```
进入容器shell交互环境 `49bced9dea83`为运行rabbitmq容器id
```
docker exec -i -t 49bced9dea83 bin/bash
```

新增超级管理员角色
```
// 设置用户名为 root 密码为admin12345
rabbitmqctl add_user root admin12345

// 设置所有权限
rabbitmqctl set_permissions -p / root ".*" ".*" ".*"

// 设置root用户administrator角色
rabbitmqctl set_user_tags root administrator

```

删除默认角色
```
rabbitmqctl delete_user guest
```

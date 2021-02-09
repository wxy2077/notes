/*
* @Time    : 2021-02-08 09:33
* @Author  : CoderCharm
* @File    : main.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :



使用consul的用处

1 服务发现 (不必关心服务端的IP和端口)
2 健康检查 (需要服务注册，能检测到服务端运行状况，不健康的会被主动剔除)
3 key/value 配置 (统一配置 相当于全局变量)

// 另一个注册配置中心
https://etcd.io/docs/v3.3.12/learning/why/#consul


两个重要的协议
gossip  八卦协议
raft	选举协议


docker pull consul

// 启动(不要改动端口)
docker run -d -p 8500:8500 consul

// 内部端口不能改，对外端口可以更改 如下,  访问地址就为 http://127.0.0.1:8600
// docker run -d -p 8600:8500 consul
// 注意事项 部署需要数据落盘

从 consul 中读取key/value配置

// 设置 key/value
http://127.0.0.1:8500/ui/dc1/kv/create

// 读取路径
micro/config/mysql

// 设置值
{
	"host": "172.16.137.129",
	"user": "root",
	"pwd": "Admin12345-",
	"port": 3306,
	"database": "micro"
}

**/
package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

//设置配置中心
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		//设置配置中心的地址
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		//设置前缀，不设置默认前缀 /micro/config
		consul.WithPrefix(prefix),
		//是否移除前缀，这里是设置为true，表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)
	//配置初始化
	config2, err := config.NewConfig()

	if err != nil {
		return config2, err
	}
	//加载配置
	err = config2.Load(consulSource)
	return config2, err
}

// MySQL配置结构体
type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

//获取 mysql 的配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	_ = config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}

func main() {
	//配置中心
	consulConfig, err := GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		fmt.Println(err)
	}

	//获取mysql配置,路径中不带前缀
	mysqlInfo := GetMysqlFromConsul(consulConfig, "mysql")
	fmt.Println(mysqlInfo)

}

/*
* @Time    : 2020-12-18 16:00
* @Author  : CoderCharm
* @File    : redis.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : Redis 初始化
**/
package initialize

import (
	"gin_study/global"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GIN_CONFIG.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GIN_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.GIN_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GIN_REDIS = client
	}
}

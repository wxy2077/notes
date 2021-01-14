/*
* @Time    : 2020-11-19 10:31
* @Author  : CoderCharm
* @File    : global.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    : 全局变量
**/
package global

import (
	"gin_study/config"
	"github.com/go-redis/redis"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GIN_DB     *gorm.DB
	GIN_VP     *viper.Viper
	GIN_CONFIG config.Server
	GIN_LOG    *zap.Logger
	GIN_REDIS  *redis.Client
	GIN_CRON   *cron.Cron
)

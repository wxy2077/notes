/*
* @Time    : 2020-11-22 15:11
* @Author  : CoderCharm
* @File    : gorm.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package initialize

import (
	"gin_study/global"
	"gin_study/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {

	return GormMysql()

}

// MysqlTables
//@author: SliverHorn
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
		//model.SysAuthority{},
		//model.System{},
		//model.CasbinModel{},
		//model.SysApi{},
		//model.SysBaseMenu{},
		//model.SysBaseMenuParameter{},
		//model.JwtBlacklist{},
		//model.SysDictionary{},
		//model.SysDictionaryDetail{},
		//model.ExaFileUploadAndDownload{},
		//model.ExaFile{},
		//model.ExaFileChunk{},
		//model.ExaSimpleUploader{},
		//model.ExaCustomer{},
		//model.SysOperationRecord{},
		//model.WorkflowProcess{},
		//model.WorkflowNode{},
		//model.WorkflowEdge{},
		//model.WorkflowStartPoint{},
		//model.WorkflowEndPoint{},
		//model.WorkflowMove{},
		//model.ExaWfLeave{},
	)
	if err != nil {
		global.GIN_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GIN_LOG.Info("register table success")
}

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	m := global.GIN_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.GIN_LOG.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(mod bool) *gorm.Config {
	if mod {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
}

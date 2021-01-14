/*
* @Time    : 2021-01-07 19:09
* @Author  : CoderCharm
* @File    : common_model.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

表通用的一些字段

**/
package global

import (
	"gorm.io/gorm"
	"time"
)

// ORM 通用字段
type COMMON_MODEL struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

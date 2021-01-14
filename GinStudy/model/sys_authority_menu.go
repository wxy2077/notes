/*
* @Time    : 2021-01-07 19:17
* @Author  : CoderCharm
* @File    : sys_authority_menu.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package model

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string                 `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}

package model

import (
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time      // 创建时间
	UpdatedAt       time.Time      // 更新时间
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityId     string         `json:"authorityId" gorm:"unique;primaryKey;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	//DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	DefaultRouter   string        `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

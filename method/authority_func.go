package method

import (
	"errors"
	"gorm.io/gorm"
	"skyplatform-auth/model"
	"skyplatform-auth/global"
)

type AuthorityService struct {
}
var AuthorityServiceApp = new(AuthorityService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthority
//@description: 创建一个角色(!!!!!!!!!!数据库表设计有问题，user表的外键导致无法独立创建新的角色!!!!!!!!!!!!)
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority

func (authorityService *AuthorityService) CreateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	if !errors.Is(global.DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = global.DB.Create(&auth).Error
	return err, auth
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority

func (authorityService *AuthorityService) UpdateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	err = global.DB.Where("authority_id = ?", auth.AuthorityId).First(&model.SysAuthority{}).Updates(&auth).Error
	return err, auth
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *model.SysAuthority
//@return: err error

func (authorityService *AuthorityService) DeleteAuthority(auth *model.SysAuthority) (err error) {
	if !errors.Is(global.DB.Where("authority_id = ?", auth.AuthorityId).First(&model.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	db := global.DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	if len(auth.SysBaseMenus) > 0 {
		err = global.DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		if err != nil {
			return
		}
		//err = db.Association("SysBaseMenus").Delete(&auth)
	} else {
		err = db.Error
		if err != nil {
			return
		}
	}
	//err = global.DB.Delete(&[]model.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error
	CasbinServiceApp.ClearCasbin(0, auth.AuthorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

// input: info request.PageInfo
// original output: (err error, list interface{}, total int64)
func (authorityService *AuthorityService) GetAuthorityInfoList() (error, []model.SysAuthority, int64){
	var authorities []model.SysAuthority
	result := global.DB.Find(&authorities)
	return result.Error, authorities, result.RowsAffected
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
	//db := global.DB.Model(&model.SysAuthority{})
	//err = db.Where("parent_id = 0").Count(&total).Error
	//var authority []model.SysAuthority
	//err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	//if len(authority) > 0 {
	//	for k := range authority {
	//		err = authorityService.findChildrenAuthority(&authority[k])
	//	}
	//}
	//return err, authority, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.SysAuthority
//@return: err error, sa model.SysAuthority

func (authorityService *AuthorityService) GetAuthorityInfo(auth model.SysAuthority) (err error, sa model.SysAuthority) {
	err = global.DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

////@author: [piexlmax](https://github.com/piexlmax)
////@function: SetDataAuthority
////@description: 设置角色资源权限
////@param: auth model.SysAuthority
////@return: error
//
//func (authorityService *AuthorityService) SetDataAuthority(auth model.SysAuthority) error {
//	var s model.SysAuthority
//	global.DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
//	err := global.DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
//	return err
//}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error

func (authorityService *AuthorityService) SetMenuAuthority(auth *model.SysAuthority) error {
	var s model.SysAuthority
	global.DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

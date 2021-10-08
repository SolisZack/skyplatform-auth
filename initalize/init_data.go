package initalize

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
	"skyplatform-auth/global"
	"skyplatform-auth/model"
	"time"
)


var authorities = []model.SysAuthority{
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "888", AuthorityName: "admin", DefaultRouter: "dashboard"},
	{CreatedAt: time.Now(), UpdatedAt: time.Now(), AuthorityId: "1000", AuthorityName: "user", DefaultRouter: "dashboard"},
}

var admins = []model.SysUser{
	{BaseModel: model.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "123456", NickName: "超级管理员", HeaderImg: "https:///qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
	{BaseModel: model.BaseModel{CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "wwd", Password: "123456", NickName: "caonimabi", HeaderImg: "https:///qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "1000"},
}

var carbines = []gormadapter.CasbinRule{
	{Ptype: "p", V0: "888", V1: "/base/login", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/user/register", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/createApi", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/getApiList", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/getApiById", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/deleteApi", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/updateApi", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/api/getAllApis", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/authority/createAuthority", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/authority/deleteAuthority", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/authority/getAuthorityList", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/authority/setDataAuthority", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/authority/updateAuthority", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/authority/copyAuthority", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/menu/getMenu", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/menu/getMenuList", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/menu/addBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/menu/getBaseMenuTree", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/menu/addMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/menu/getMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/menu/updateBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/menu/getBaseMenuById", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/user/getUserList", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/user/setUserAuthority", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/user/deleteUser", V2: "DELETE"},
	{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/upload", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/casbin/updateCasbin", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/casbin/casbinTest/:pathParam", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/system/getSystemConfig", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/system/setSystemConfig", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/system/getServerInfo", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "DELETE"},
	{Ptype: "p", V0: "888", V1: "/customer/customer", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/customer/customerList", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/autoCode/createTemp", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/autoCode/preview", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/autoCode/getTables", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/autoCode/getDB", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/autoCode/getColumn", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/createSysDictionaryDetail", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/deleteSysDictionaryDetail", V2: "DELETE"},
	{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/updateSysDictionaryDetail", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/findSysDictionaryDetail", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/sysDictionaryDetail/getSysDictionaryDetailList", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/sysDictionary/createSysDictionary", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/sysDictionary/deleteSysDictionary", V2: "DELETE"},
	{Ptype: "p", V0: "888", V1: "/sysDictionary/updateSysDictionary", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/sysDictionary/findSysDictionary", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/sysDictionary/getSysDictionaryList", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
	{Ptype: "p", V0: "888", V1: "/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},
	{Ptype: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
	{Ptype: "p", V0: "888", V1: "/email/emailTest", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/simpleUploader/upload", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/simpleUploader/checkFileMd5", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/simpleUploader/mergeFileMd5", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/excel/importExcel", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/excel/loadExcel", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/excel/exportExcel", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/excel/downloadTemplate", V2: "GET"},
	{Ptype: "p", V0: "888", V1: "/api/deleteApisByIds", V2: "DELETE"},
	{Ptype: "p", V0: "888", V1: "/autoCode/getSysHistory", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/autoCode/rollback", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/autoCode/getMeta", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/autoCode/delSysHistory", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/user/setUserAuthorities", V2: "POST"},
	{Ptype: "p", V0: "888", V1: "/user/getUserInfo", V2: "GET"},
	{Ptype: "p", V0: "8881", V1: "/base/login", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/user/register", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/api/createApi", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/api/getApiList", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/api/getApiById", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/api/deleteApi", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/api/updateApi", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/api/getAllApis", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/authority/createAuthority", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/authority/deleteAuthority", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/authority/getAuthorityList", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/authority/setDataAuthority", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/menu/getMenu", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/menu/getMenuList", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/menu/addBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/menu/getBaseMenuTree", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/menu/addMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/menu/getMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/menu/updateBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/menu/getBaseMenuById", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/user/changePassword", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/user/getUserList", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/user/setUserAuthority", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/upload", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/casbin/updateCasbin", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/jwt/jsonInBlacklist", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/system/getSystemConfig", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/system/setSystemConfig", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "POST"},
	{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "PUT"},
	{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "DELETE"},
	{Ptype: "p", V0: "8881", V1: "/customer/customer", V2: "GET"},
	{Ptype: "p", V0: "8881", V1: "/customer/customerList", V2: "GET"},
	{Ptype: "p", V0: "8881", V1: "/user/getUserInfo", V2: "GET"},
	{Ptype: "p", V0: "9528", V1: "/base/login", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/user/register", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/api/createApi", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/api/getApiList", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/api/getApiById", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/api/deleteApi", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/api/updateApi", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/api/getAllApis", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/authority/createAuthority", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/authority/deleteAuthority", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/authority/getAuthorityList", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/authority/setDataAuthority", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/menu/getMenu", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/menu/getMenuList", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/menu/addBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/menu/getBaseMenuTree", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/menu/addMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/menu/getMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/menu/deleteBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/menu/updateBaseMenu", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/menu/getBaseMenuById", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/user/changePassword", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/user/getUserList", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/user/setUserAuthority", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/upload", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/casbin/updateCasbin", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/jwt/jsonInBlacklist", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/system/getSystemConfig", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/system/setSystemConfig", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "PUT"},
	{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "DELETE"},
	{Ptype: "p", V0: "9528", V1: "/customer/customer", V2: "GET"},
	{Ptype: "p", V0: "9528", V1: "/customer/customerList", V2: "GET"},
	{Ptype: "p", V0: "9528", V1: "/autoCode/createTemp", V2: "POST"},
	{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_authorities 表数据初始化
func InitAuthority() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("authority_id IN ? ", []string{"888", "9528"}).Find(&[]model.SysAuthority{}).RowsAffected == 2 {
			log.Println("ERROR: sys_authorities 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorities).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		log.Println("sys_authorities 表初始数据成功!")
		return nil
	})
}



//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_users 表数据初始化
func InitUser() error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysUser{}).RowsAffected == 2 {
			log.Println("ERROR: sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		log.Println("sys_users 表初始数据成功!")
		return nil
	})
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: casbin_rule 表数据初始化
func InitCasbin() error {
	global.DB.AutoMigrate(gormadapter.CasbinRule{})
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&[]gormadapter.CasbinRule{}).RowsAffected == 154 {
			log.Println("ERROR: casbin_rule 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&carbines).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		log.Println("casbin_rule 表初始数据成功!")
		return nil
	})
}
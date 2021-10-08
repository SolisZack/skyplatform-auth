package api

import (
	"skyplatform-auth/method"
)

//type ApiGroup struct {
//	AuthorityApi
//	CasbinApi
//}

var userService = method.UserService{}
var authorityService = method.AuthorityService{}
var casbinService = method.CasbinService{}
//var autoCodeService = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
//var autoCodeHistoryService = service.ServiceGroupApp.SystemServiceGroup.AutoCodeHistoryService
//var dictionaryService = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
//var dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
//var initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
var jwtService = method.JWT{}
//var baseMenuService = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
//var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService


package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"skyplatform-auth/model"
	"skyplatform-auth/request"
	"skyplatform-auth/response"
	"skyplatform-auth/utils"
)

type AuthorityApi struct {
}

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body authority.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority model.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	log.Println(authority)
	if err := utils.Verify(authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := authorityService.CreateAuthority(authority); err != nil {
		log.Println("ERROR: create authority failed ", err)
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		//_ = menuService.AddMenuAuthority(request.DefaultMenu(), authority.AuthorityId)
		_ = casbinService.UpdateCasbin(authority.AuthorityId, request.DefaultCasbin())
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}

// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body authority.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authority/deleteAuthority [post]
func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var authority model.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorityService.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		log.Println("ERROR: delete authority failed ", err)
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body authority.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authority/updateAuthority [post]
func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	//if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err, authority := authorityService.UpdateAuthority(auth); err != nil {
		log.Println("ERROR: update authority failed ", err)
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authority}, "更新成功", c)
	}
}

// @Tags Authority
// @Summary 获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	//if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err, list, total := authorityService.GetAuthorityInfoList(); err != nil {
		log.Println("ERROR: get authority info failed ", err)
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		log.Println(list)
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
		}, "获取成功", c)
	}
}

//// @Tags Authority
//// @Summary 设置角色资源权限
//// @Security ApiKeyAuth
//// @accept application/json
//// @Produce application/json
//// @Param data body authority.SysAuthority true "设置角色资源权限"
//// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
//// @Router /authority/setDataAuthority [post]
//func (a *AuthorityApi) SetDataAuthority(c *gin.Context) {
//	var auth model.SysAuthority
//	_ = c.ShouldBindJSON(&auth)
//	//if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
//	//	response.FailWithMessage(err.Error(), c)
//	//	return
//	//}
//	if err := authorityService.SetDataAuthority(auth); err != nil {
//		log.Println("ERROR: set authority failed ", err)
//		response.FailWithMessage("设置失败"+err.Error(), c)
//	} else {
//		response.OkWithMessage("设置成功", c)
//	}
//}

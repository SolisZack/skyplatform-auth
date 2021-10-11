package service

import "github.com/gin-gonic/gin"

func StartService() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	user := router.Group("/auth/user").Use(JWTAuth())
	{
		user.POST("/changePassword", userApi.ChangePassword)
		user.POST("/setUserAuthority", userApi.SetUserAuthority)
		user.POST("/deleteUser", userApi.DeleteUser)
		user.POST("/setUserInfo", userApi.SetUserInfo)
	}
	userWithoutToken := router.Group("/auth/user")
	{
		userWithoutToken.POST("/login", userApi.Login)
		userWithoutToken.POST("/register", userApi.Register)
		userWithoutToken.GET("/getUserList", userApi.GetUserList)
		userWithoutToken.GET("/getUserInfo", userApi.GetUserInfo)
	}

	casbin := router.Group("/auth/casbin").Use(JWTAuth())
	{
		casbin.POST("/updateCasbin", casbinApi.UpdateCasbin)
		casbin.POST("/getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}

	authority := router.Group("/auth/authority").Use(JWTAuth())
	{
		authority.POST("/createAuthority", authApi.CreateAuthority)
		authority.POST("/deleteAuthority", authApi.DeleteAuthority)
		authority.POST("/updateAuthority", authApi.UpdateAuthority)
		authority.GET("/getAuthorityList", authApi.GetAuthorityList)
	}

	router.Run("127.0.0.1:8080")
}

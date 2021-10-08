package service

import "github.com/gin-gonic/gin"

func StartService() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//router.Use(JWTAuth())

	user := router.Group("/auth/user")
	{
		user.POST("/login", userApi.Login)
		user.POST("/register", userApi.Register)
		user.POST("/changePassword ", userApi.ChangePassword)
		user.POST("/setUserAuthority", userApi.SetUserAuthority)
		user.POST("/deleteUser", userApi.DeleteUser)
		user.POST("/setUserInfo", userApi.SetUserInfo)
		user.GET("/getUserList", userApi.GetUserList)
		user.GET("/getUserInfo", userApi.GetUserInfo)
	}

	casbin := router.Group("/auth/casbin")
	{
		casbin.POST("/updateCasbin", casbinApi.UpdateCasbin)
		casbin.POST("/getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}

	authority := router.Group("/auth/authority")
	{
		authority.POST("/createAuthority", authApi.CreateAuthority)
		authority.POST("/deleteAuthority", authApi.DeleteAuthority)
		authority.POST("/updateAuthority", authApi.UpdateAuthority)
		authority.GET("/getAuthorityList", authApi.GetAuthorityList)
	}



	router.Run("127.0.0.1:8080")
}

package api

import (
	"github.com/gin-gonic/gin"
	"skyplatform-auth/global"
	"skyplatform-auth/method"
	"testing"
)

//func TestAuthAPI(t *testing.T) {
//	api := AuthorityApi{}
//	global.DB = method.InitDB()
//	r := gin.Default()
//	r.GET("/test", api.GetAuthorityList)
//	r.Run("127.0.0.1:8080")
//}

//func TestAuthAPI(t *testing.T) {
//	api := CasbinApi{}
//	global.DB = method.InitDB()
//	r := gin.Default()
//	r.POST("/test", api.UpdateCasbin)
//	r.Run("127.0.0.1:8080")
//}

func TestUserAPI(t *testing.T) {
	api := UserApi{}
	global.DB = method.InitDB()
	global.ETCD = method.InitETCD()
	r := gin.Default()
	r.GET("/test", api.GetUserList)
	r.Run("127.0.0.1:8080")
}
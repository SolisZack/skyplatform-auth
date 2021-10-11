package method

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"skyplatform-auth/global"
	"skyplatform-auth/request"
	"testing"
	"time"
)

//func TestAuthFunc(t *testing.T){
//	global.DB = InitDB()
//	var testAuthority = model.SysAuthority{
//		CreatedAt: time.Now(),
//		UpdatedAt: time.Now(),
//		AuthorityId: "888",
//		AuthorityName: "ADMIN",
//		DefaultRouter: "dashboard",
//	}
//
//	_, authList, _ := AuthorityServiceApp.GetAuthorityInfoList()
//	t.Log(authList)
//}

//func TestCasbinFunc(t *testing.T){
//	global.DB = InitDB()
//	testAuthID := "888"
//	testCasbinInfo := []request.CasbinInfo{
//		{Path: "/user/getUserInfo", Method: "GET"},
//		{Path: "/autoCode/createTemp ", Method: "GET"},
//	}
//	CasbinServiceApp.Casbin()
//
//	err := CasbinServiceApp.UpdateCasbin(testAuthID, testCasbinInfo)
//	if err != nil{
//		t.Log("ERROR: casbin test failed:", err)
//	}
//
//	t.Log(CasbinServiceApp.GetPolicyPathByAuthorityId("888"))
//
//}

//func TestUserFunc(t *testing.T) {
//	global.DB = InitDB()
//	//var testUser = model.SysUser{
//	//	BaseModel: auth.BaseModel{
//	//		CreatedAt: time.Now(),
//	//		UpdatedAt: time.Now(),
//	//	},
//	//	UUID:        uuid.NewV4(),
//	//	Username:    "wsx",
//	//	Password:    "980811",
//	//	NickName:    "超级管理员",
//	//	HeaderImg:   "https:///qmplusimg.henrongyi.top/gva_header.jpg",
//	//	AuthorityId: "888",
//	//}
//
//	userUUID, _ := uuid.FromString("681c22e6-a6f7-4d6d-b0b7-58699c746b56")
//	err, user := UserServiceApp.GetUserInfo(userUUID)
//	if err != nil{
//		t.Log("ERROR: user test failed:", err)
//	}
//	t.Log(user)
//
//}

func TestJWTFunc(t *testing.T){
	claims := request.CustomClaims{
		UUID:        uuid.NewV4(),
		ID:          123,
		Username:    "wsx",
		AuthorityId: "888",
		BufferTime:  111111111, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + 11111111111111, 	  // 过期时间 7天  配置文件
			Issuer:    "wwd",                                              // 签名的发行者
		},
	}
	global.DB = InitDB()
	global.ETCD = InitETCD()
	JWTApp := NewJWT()
	token, err := JWTApp.CreateToken(claims)
	if err != nil {
		t.Fatal(err)
	}
	err = JWTApp.SetETCDJWT("wsx", token)
	if err != nil {
		t.Fatal(err)
	}

	result, err := JWTApp.GetETCDJWT("wsx")
	if err != nil {
		t.Fatal(err)
	}
	parseResult, err := JWTApp.ParseToken(result)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(parseResult.Username)

}
package initalize

import (
	"skyplatform-auth/model"
	"skyplatform-auth/global"
	"skyplatform-auth/method"
)

func InitDB()  {
	global.DB = method.InitDB()
	//method.MysqlTables(global.DB)
}

func InitETCD() {
	global.ETCD = method.InitETCD()
}

func InitConfig() {
	global.CONFIG.Casbin = model.Casbin{ModelPath: "/home/wwd/go/src/skyplatform/common/auth/method/rbac_model.conf"}

	global.CONFIG.JWT = model.JWT{
		SigningKey: "123456",
		ExpiresTime: 1111111,
		BufferTime: 1111111,
	}

	global.CONFIG.Mysql = model.Mysql{}
	global.CONFIG.Mysql.Default()
}


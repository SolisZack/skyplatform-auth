package method

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"skyplatform-auth/model"
)

func InitDB() *gorm.DB {
	m := model.Mysql{}
	m.Default()
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Dsn()
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		log.Println("Error: Open mysql failed")
		log.Fatalln(err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}


// MysqlTables
//@author: SliverHorn
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysAuthority{},
		model.SysUser{},

	)
	if err != nil {
		log.Println("ERROR: register table failed")
		log.Fatalln(err)
		return
	}
	log.Println("register table success")
}


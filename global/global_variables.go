package global

import (
	client "go.etcd.io/etcd/client/v3"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"skyplatform-auth/model"
)

const (
	UserPrefix = "/user/"
)

var (
	DB                 *gorm.DB
	ETCD               *client.Client
	CONFIG             model.Server
	ConcurrencyControl = 	&singleflight.Group{}
)


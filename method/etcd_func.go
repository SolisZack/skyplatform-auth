package method

import (
	"errors"
	client "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

var (
	ErrorKeyNotFound = errors.New("key not found")
)


func InitETCD() *client.Client{
	// ETCD
	etcdCfg := client.Config{
		Endpoints:   []string{"0.0.0.0:2379"},
		DialTimeout: 5 * time.Second,
	}
	etcdClient, err := client.New(etcdCfg)
	if err != nil {
		log.Fatal("Error: cannot connect to etcd: ", err)
	}
	
	return etcdClient
}

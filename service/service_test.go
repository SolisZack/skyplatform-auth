package service

import (
	"skyplatform-auth/initalize"
	"testing"
)

func TestService(t *testing.T) {
	initalize.InitDB()
	initalize.InitConfig()
	initalize.InitETCD()
	StartService()
}

package initalize

import (
	"testing"
)

func TestInit(t *testing.T){
	InitDB()
	InitCasbin()
	InitUser()
	InitAuthority()

}

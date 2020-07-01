package config

import (
	"fmt"
	"testing"
)

func TestIfConfig_SetInterfaceCfg(t *testing.T) {
	//获取redis链接
	err := CmlDbReconnect(CFG_DB)
	if err != nil {
		fmt.Println("db state error!")
		return
	}
	dbclient := cmldbinfo.cmldbclient[CmlDbNo[CFG_DB]]
	err = IfNameSetDb("test", FIELD_INTF_INDEX, "123", dbclient)
}

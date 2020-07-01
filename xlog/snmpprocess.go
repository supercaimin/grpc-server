/*告警、日志的snmp 处理文件,用于snmp trap发送以及snmp get
  陈维  20200227*/
package xlog

import (
	"fmt"
	"net"
	"strconv"

	"github.com/k-sone/snmpgo"
)

func snmptrap(module, code, level, desc string, status bool, sver net.UDPAddr) {

	var sverstr string

	sverstr = sver.IP.String()
	sverstr += ":" + strconv.Itoa(sver.Port)
	snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.V2c,
		Address:   sverstr, /*"192.168.50.1:162"*/
		Retries:   1,
		Community: "public",
	})
	if err != nil {
		// Failed to create snmpgo.SNMP object
		fmt.Println(err)
		return
	}

	// Build VarBind list
	var varBinds snmpgo.VarBinds
	varBinds = append(varBinds, snmpgo.NewVarBind(snmpgo.OidSysUpTime, snmpgo.NewTimeTicks(1000)))

	oid, _ := snmpgo.NewOid("1.3.6.1.6.3.1.1.5.3")
	varBinds = append(varBinds, snmpgo.NewVarBind(snmpgo.OidSnmpTrap, oid))

	oid, _ = snmpgo.NewOid("1.3.6.1.2.1.2.2.1.1.3")
	varBinds = append(varBinds, snmpgo.NewVarBind(oid, snmpgo.NewInteger(2)))

	oid, _ = snmpgo.NewOid("1.3.6.1.2.1.31.1.1.1.1.2")
	varBinds = append(varBinds, snmpgo.NewVarBind(oid, snmpgo.NewOctetString([]byte(module))))

	oid, _ = snmpgo.NewOid("1.3.6.1.2.1.31.1.1.1.1.3")
	varBinds = append(varBinds, snmpgo.NewVarBind(oid, snmpgo.NewOctetString([]byte(code))))

	oid, _ = snmpgo.NewOid("1.3.6.1.2.1.31.1.1.1.1.4")
	varBinds = append(varBinds, snmpgo.NewVarBind(oid, snmpgo.NewOctetString([]byte(level))))

	oid, _ = snmpgo.NewOid("1.3.6.1.2.1.31.1.1.1.1.5")
	varBinds = append(varBinds, snmpgo.NewVarBind(oid, snmpgo.NewOctetString([]byte(desc))))

	oid, _ = snmpgo.NewOid("1.3.6.1.2.1.31.1.1.1.1.6")
	if status {
		varBinds = append(varBinds, snmpgo.NewVarBind(oid, snmpgo.NewOctetString([]byte("occur"))))
	} else {
		varBinds = append(varBinds, snmpgo.NewVarBind(oid, snmpgo.NewOctetString([]byte("clear"))))
	}

	if err = snmp.Open(); err != nil {
		// Failed to open connection
		fmt.Println(err)
		return
	}
	defer snmp.Close()

	if err = snmp.V2Trap(varBinds); err != nil {
		// Failed to request
		fmt.Println(err)
		return
	}
}

// Cml_subsys_module_define.go
/* Cml_subsys_module_define.go
	作者：
	日期：2019-10-12

   CML各子系统内的模块ID编号定义,用于命令ID。各子系统的各业务模块的具体命令ID定义在独立文件中

*/

package config

/*子系统ID,从1开始编号,0保留不用*/
const (
	CMD_SUBSYSTEMID_SYSMGR = 1
	CMD_SUBSYSTEMID_DBS    = 2
	CMD_SUBSYSTEMID_OAM    = 3
	CMD_SUBSYSTEMID_FRR    = 4
	CMD_SUBSYSTEMID_APP    = 5
	CMD_SUBSYSTEMID_DIAG   = 6
	CMD_SUSSYSTEMID_MAX    = 6 /*当前确定的最大子系统ID*/
)

/*SYSMGR子系统内的业务模块定义*/
const (
	CML_SYSMGR_MODULEID_STACK  = 1 /*多虚一模块*/
	CML_SYSMGR_MODULEID_SYSLOG = 2 /*syslog模块*/
)
const CMD_SYSMGR_BASE_STACK = (CMD_SUBSYSTEMID_SYSMGR<<24 + CML_SYSMGR_MODULEID_STACK<<9)
const CMD_SYSMGR_BASE_SYSLOG = (CMD_SUBSYSTEMID_SYSMGR<<24 + CML_SYSMGR_MODULEID_STACK<<9)

/*OAM子系统内的业务模块定义*/
const (
	CML_OAM_MODULEID_USERMNG = 1 /*用户管理模块*/
	CML_OAM_MOUDLEID_SSH     = 2 /*SSH管理模块*/
)

/*路由协议平台*/
const (
	CML_FRR_MODULEID_IP  = 1 /*ip模块,如IP的ping等*/
	CML_FRR_MODULEID_TCP = 2
	CML_FRR_MODULEID_BGP = 3
)

/*应用子系统*/
const (
	CML_APP_MODULEID_ACL      = 1 /*ACL模块*/
	CML_APP_MODULEID_ROUTEMAP = 2 /*ROUTEMAP模块，用于策略路由及路由过滤*/
	CML_APP_MODULEID_QOS      = 3 /*QOS模块*/
)

const (
	CMD_APP_BASE_ACL      = (CMD_SUBSYSTEMID_APP<<24 + CML_APP_MODULEID_ACL<<9)
	CMD_APP_BASE_ROUTEMAP = (CMD_SUBSYSTEMID_APP<<24 + CML_APP_MODULEID_ROUTEMAP<<9)
	CMD_APP_BASE_QOS      = (CMD_SUBSYSTEMID_APP<<24 + CML_APP_MODULEID_QOS<<9)
)

/*诊断子系统*/

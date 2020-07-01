package config

import (
	"context"
	pb "hstcmler/cml"
	"strconv"
)

// cmdid 参考cmdid-const-define.h文件
//#define INTERFACE_CMD_SET_TYPE_NUM       0x01020006
//#define INTERFACE_CMD_LOOKBACK        	 0x02020006
//#define INTERFACE_CMD_NULL        	  	 0x03020006
type IfConfig struct {
}

func NewIfConfig() *IfConfig {
	return new(IfConfig)
}

func (*IfConfig) SetInterfaceCfg(ctx context.Context, req *pb.Ifbaseinfocfg) (*pb.Cfgrtninfo, error) {
	var cmdRtnInfo pb.Cfgrtninfo
	cmdRtnInfo.Rtncode = 1
	var dbchange bool = false
	var err error = nil
	var ifdb = NewIfDb()

	//链接数据库
	err = ifdb.InitDbClient()
	if err != nil {
		cmdRtnInfo.Rtnstr = "db state error!\n"
		return &cmdRtnInfo, err
	}

	//2.1.1首先判断是否存在ifname,如果不存在，则结束
	ifname := req.GetIfname()
	if "" == ifname {
		cmdRtnInfo.Rtnstr = "Error, no ifname!\n"
		return &cmdRtnInfo, err
	}

	//根据命令id判断所属命令
	switch req.Cmdid {
	case IfCmdSetTypeNum:
		err = ifdb.IfTypeNumSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "interface set db failed!\n"
		} else {
			dbchange = true
		}
	case IfCmdLoopBack:
		err = ifdb.IfLookBackSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "interface loopback set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdNull:
		err = ifdb.IfNullSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "interface null set db failed!\n"
		} else {
			dbchange = true
		}
	case IfCmdDescrip:
		err = ifdb.IfDescripSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "description set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdClearCfgThis:
		err = ifdb.IfClearConfInfoSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "clear this set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdClearCfgIntf:
		err = ifdb.IfClearConfInfoSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "clear configuration set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdProtoUpDelay:
		err = ifdb.IfProtoUpDelaySetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "proto up delay set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdFlowStatInterval:
		err = ifdb.IfFlowStatIntervalSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "flow stat interval set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdShutdown:
		err = ifdb.IfShutdownSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "shutdown set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdShutdownIntf:
		err = ifdb.IfShutdownInfoSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "shutdown Info set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdShutdownNet:
		err = ifdb.IfShutdownNetSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "shutdown network set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdShutdownTransmit:
		err = ifdb.IfShutdownTransmitSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "shutdown transmit set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdShutdownReceive:
		err = ifdb.IfShutdownReceiveSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "shutdown receive  set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdFecModeBase:
		err = ifdb.IfFecModeBaseSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "fec mode set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdFecModeNone:
		err = ifdb.IfFecModeNoneSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "fec mode set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdFecModeRS:
		err = ifdb.IfFecModeRSSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "fec mode set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdControlFlap:
		err = ifdb.IfControlFlapSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "control flap set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdIfRange:
		//field暂未定义
		err = ifdb.IfRangeSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "interface range set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdJumboframe:
		err = ifdb.IfJumboframeSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "jumboframe set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdLoopback:
		err = ifdb.IfLoopbackSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "loopback set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdPortModeLan:
		err = ifdb.IfPortModeLanSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "port mode set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdPortModeWan:
		err = ifdb.IfPortModeWanSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "port mode set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdPortMode10g:
		err = ifdb.IfPortMode10gSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "port mode set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdPortModeGe:
		err = ifdb.IfPortModeGeSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "port mode set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdPortMode25ge:
		err = ifdb.IfPortMode25geSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "port mode set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdPortMode100ge:
		err = ifdb.IfPortMode100geSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "port mode set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdPortSwitch:
		err = ifdb.IfPortSwitchSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "portswitch set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdPortSwitchBatch:
		err = ifdb.IfPortSwitchBatchSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "portswitch set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSetUpDelay:
		err = ifdb.IfUpDelaySetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "up delay set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSpeed10:
		err = ifdb.IfSpeed10SetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "speed set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSpeed100:
		err = ifdb.IfSpeed100SetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "speed set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSpeed1000:
		err = ifdb.IfSpeed1000SetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "speed set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSpeed10000:
		err = ifdb.IfSpeed10000SetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "speed set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSpeed40000:
		err = ifdb.IfSpeed40000SetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "speed set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSpeedAuto10:
		err = ifdb.IfSpeedAuto10SetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "speed set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSpeedAuto100:
		err = ifdb.IfSpeedAuto100SetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "speed set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSpeedAuto1000:
		err = ifdb.IfSpeedAuto1000SetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "speed set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSpeedAuto10000:
		err = ifdb.IfSpeedAuto10000SetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "speed set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdSubIf:
		err = ifdb.IfSubIfSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "subinterface trap updown set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdStatistics:
		err = ifdb.IfStatisticsSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "statistics set db failed!\n"
		} else {
			dbchange = true
		}

	case IfCmdMtu:
		err = ifdb.IfMtuSetDb(req)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "flow stat interval set db failed!\n"
		} else {
			dbchange = true
		}
	}

	/*作为数据库操作,需要发布更新,带上时间等*/
	flushtime := strconv.FormatInt(req.GetUpdatetime(), 10)
	if dbchange {
		CmlDbCfgPubstring(CFG_DB, IfCfgChan, flushtime)
	}

	/*不能把err的信息透传到客户端,只把错误信息的字符串转换为提示信息带回*/
	return &cmdRtnInfo, nil
}

func (*IfConfig) DelInterfaceCfg(ctx context.Context, req *pb.Ifbaseinfocfg) (*pb.Cfgrtninfo, error) {

	return nil, nil
}

func (*IfConfig) ShowInterfaceCfg(ctx context.Context, req *pb.Showcfginfo) (*pb.Showrtninfo, error) {

	return nil, nil
}

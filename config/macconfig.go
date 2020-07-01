package config

import (
	"context"
	"fmt"
	pb "hstcmler/cml"
	"strconv"
)

type MacConfig struct {
}

func NewMacConfig() *MacConfig {
	return new(MacConfig)
}

const (
	MacMaxNum = 256
)

//Macglobalcfg的数据结构对应 mac-config.proto中的message macglobalcfg
//message macglobalcfg {
//	int32 agingtime = 1; //mac老化时间, 0xffffffff表示未配置
//	macblackhole machole = 2; //
//	macflapcfg  macflap = 3; //mac漂移相关配置
//	repeated macstatic static = 4;//静态mac配置
//	int64 updatetime = 6;
//}
func (self *MacConfig) SetMacGlobalCfg(ctx context.Context, req *pb.Macglobalcfg) (*pb.Cfgrtninfo, error) {
	var cmdRtnInfo pb.Cfgrtninfo
	var staticinfo []*pb.Macstatic
	var err error
	var dbchange bool
	var macdb = NewMacDb()
	//chenwei:链接数据库
	err = macdb.InitDbClient()
	if err != nil {
		cmdRtnInfo.Rtnstr = "db state error!\n"
		return &cmdRtnInfo, err
	}
	flushtime := strconv.FormatInt(req.GetUpdatetime(), 10)
	switch req.GetCmdid() {
	case MacCmdAgingTime:
		//该变量后配置的覆盖前配置,无其它依赖检查和冲突检查
		//调用接口写db
		err = macdb.MacAgeTimeSetDb(req.GetAgingtime())
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "Agingtime set db failed!\n"
		} else {
			dbchange = true
		}
	case MacCmdBlackHole:
		//通过vlan接口key是否存在来判断vlan id是否已配置
		//如果是判断key下field是否有值,用HExists(KEY, FIELD).Result()
		namestr := KEY_INTF_VLANIF_PREFIX + strconv.FormatInt(int64(req.GetMachole().GetVlanid()), 10)
		exist, _ := macdb.MacIsExistData(namestr)
		if 1 == exist {
			//调用db接口写入db中
			err = macdb.MacBlackholeSetDb(req.GetMachole())
			if nil != err {
				cmdRtnInfo.Rtncode = 2
				cmdRtnInfo.Rtnstr = "Blackhole mac set db failed!\n"
			} else {
				dbchange = true
			}
		} else {
			//vlan id未配置,返回提示信息
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += fmt.Sprintf("Error, vlan %d not exist!\n", req.GetMachole().GetVlanid())
		}
	case MacCmdFlapAgingTime:
		//该变量后配置的覆盖前配置,无其它依赖检查和冲突检查
		//调用接口写db
		err = macdb.MacFlapingagetimeSetDb(req.GetMacflap().GetFlapingagetime())
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "Flapingagetime set db failed!\n"
		} else {
			dbchange = true
		}
	case MacCmdFlapDetect:
		fallthrough
	case MacCmdFlapDetectLevelLow:
		fallthrough
	case MacCmdFlapDetectLevelMiddle:
		fallthrough
	case MacCmdFlapDetectLevelHigh:
		err = macdb.MacFlapdetectSetDb(req.GetMacflap().GetFlapdetectlevel())
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "Macflapdetect set db failed!\n"
		} else {
			dbchange = true
		}
	case MacCmdStatic:
		for _, datainfo := range staticinfo {
			//检查mac对应的vlan是否已配置,如果没配置则不允许配置
			namestr := KEY_INTF_VLANIF_PREFIX + strconv.FormatInt(int64(datainfo.GetVlanid()), 10)
			exist, _ := macdb.MacIsExistData(namestr)
			if 1 == exist {
				//调用db接口写入db中
				err = macdb.MacStaticSetDb(datainfo)
				if nil != err {
					cmdRtnInfo.Rtncode = 2
					cmdRtnInfo.Rtnstr = "MacStatic mac set db failed!\n"
				} else {
					dbchange = true
				}
			} else {
				//vlan id未配置,返回提示信息
				cmdRtnInfo.Rtncode = 2
				cmdRtnInfo.Rtnstr += fmt.Sprintf("Error, vlan %d not exist!\n", datainfo.GetVlanid())
			}
		}
	default:
		return &cmdRtnInfo, nil
	}
	/*作为数据库操作,需要发布更新,带上时间等*/
	if dbchange {
		CmlDbCfgPubstring(CFG_DB, MacCfgChan, flushtime)
	}
	return &cmdRtnInfo, nil
}

func (self *MacConfig) DelMacGlobalCfg(ctx context.Context, req *pb.Macglobalcfg) (*pb.Cfgrtninfo, error) {
	var cmdRtnInfo pb.Cfgrtninfo
	var dbchange bool = false
	var flushtime int64
	var rtnstring string = ""
	var err error
	var macdb = NewMacDb()
	err = macdb.InitDbClient()
	if err != nil {
		cmdRtnInfo.Rtnstr = "db state error!\n"
		return &cmdRtnInfo, err
	}
	flushtime = req.GetUpdatetime()
	//根据key执行删除
	if nil != req {
		var reserr error = nil
		switch req.GetCmdid() {
		case MacCmdAgingTime:
			reserr = macdb.MacAgeTimeDelDb()
		case MacCmdBlackHole:
			reserr = macdb.MacBlackholeDelDb(req.GetMachole().GetMac())
		case MacCmdFlapAgingTime:
			reserr = macdb.MacFlapingagetimeDelDb()
		case MacCmdFlapDetect:
			fallthrough
		case MacCmdFlapDetectLevelLow:
			fallthrough
		case MacCmdFlapDetectLevelMiddle:
			fallthrough
		case MacCmdFlapDetectLevelHigh:
			reserr = macdb.MacFlapdetectDelDb(req.GetMacflap().GetFlapdetectlevel())
		case MacCmdStatic:
			for _, datainfo := range req.GetStatic() {
				namestr := KEY_INTF_VLANIF_PREFIX + strconv.FormatInt(int64(datainfo.GetVlanid()), 10)
				exist, _ := macdb.MacIsExistData(namestr)
				if 1 == exist {
					reserr = macdb.MacStaticDelDb(datainfo.GetIfname(), datainfo.GetMac())
				}
			}
		default:
			return &cmdRtnInfo, nil
		}
		if reserr != nil {
			rtnstring += "mac del db error! \n"
		}
		dbchange = true
	}

	if dbchange {
		err = macdb.MacDbChange(flushtime)
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(CFG_DB, MacCfgChan, strconv.FormatInt(flushtime, 10))
	}
	cmdRtnInfo.Rtnstr, _ = rtnstring, err
	cmdRtnInfo.Rtncode = 1
	return &cmdRtnInfo, nil
}

func (self *MacConfig) ShowMacGlobalCfg(ctx context.Context, req *pb.Showcfginfo) (*pb.Showrtninfo, error) {
	var showRtninfo pb.Showrtninfo
	showRtninfo.Cmdcode = req.GetCmdcode()
	showRtninfo.Regstr = "qos Showcfginfo command return!!!"

	//switch req.GetCmdid() {
	//}
	//switch(req.GetShowoption()){
	/*过滤等选项*/
	//}

	return &showRtninfo, nil
}

func (self *MacConfig) SetMacIfCfg(ctx context.Context, req *pb.Macifcfg) (*pb.Cfgrtninfo, error) {
	var cmdRtnInfo pb.Cfgrtninfo
	var err error
	var dbchange bool

	//获取数据库连接
	var macdb = NewMacDb()
	err = macdb.InitDbClient()
	if err != nil {
		cmdRtnInfo.Rtnstr = "db state error!\n"
		return &cmdRtnInfo, err
	}

	flushtime := strconv.FormatInt(req.GetUpdatetime(), 10)
	switch req.GetCmdid() {
	case MacCmdLearn:
		err = macdb.MaclearndisSetDb(req.GetMaclearn().GetAction())
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "Learndisable set db failed!\n"
		} else {
			dbchange = true
		}
	case MacCmdLimit:
		afterstr := strconv.FormatInt(int64(req.GetMaclimit().GetMaxnum()), 10) + req.GetMaclimit().GetAction() + strconv.FormatInt(int64(req.GetMaclimit().GetAlarm()), 10)
		err = macdb.MaclimitdisSetDb(afterstr)
		if nil != err {
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += "Maclimit set db failed!\n"
		} else {
			dbchange = true
		}
	case MacCmdPortBridge:
		//检查mac对应的bridge是否已配置,如果没配置则不允许配置
		namestr := KEY_INTF_BRIDGE_IF_PREFIX + strconv.FormatInt(int64(req.GetPortbridge()), 10)
		exist, _ := macdb.MacIsExistData(namestr)
		if 1 == exist {
			err = macdb.MacPortbridgeSetDb(req.GetPortbridge())
			if nil != err {
				cmdRtnInfo.Rtncode = 2
				cmdRtnInfo.Rtnstr += "Portbridge set db failed!\n"
			} else {
				dbchange = true
			}
		} else {
			//bridge未配置,返回提示信息
			cmdRtnInfo.Rtncode = 2
			cmdRtnInfo.Rtnstr += fmt.Sprintf("Error, bridge %d not exist!\n", req.GetPortbridge())
		}
	default:
		return &cmdRtnInfo, nil
	}
	/*作为数据库操作,需要发布更新,带上时间等*/
	if dbchange {
		CmlDbCfgPubstring(CFG_DB, MacCfgChan, flushtime)
	}
	return &cmdRtnInfo, nil
}

func (self *MacConfig) DelMacIfCfg(ctx context.Context, req *pb.Macifcfg) (*pb.Cfgrtninfo, error) {
	var cmdRtnInfo pb.Cfgrtninfo
	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var rtnstring string = ""

	//获取数据库连接
	var macdb = NewMacDb()
	err = macdb.InitDbClient()
	if err != nil {
		cmdRtnInfo.Rtnstr = "db state error!\n"
		return &cmdRtnInfo, err
	}

	flushtime = req.GetUpdatetime()
	//根据key执行删除
	if nil != req {
		var reserr error = nil
		switch req.GetCmdid() {
		case MacCmdLearn:
			reserr = macdb.MaclearndisDelDb()
		case MacCmdLimit:
			reserr = macdb.MaclimitdisDelDb()
		case MacCmdPortBridge:
			//检查mac对应的bridge是否已配置,如果没配置则不允许配置
			namestr := KEY_INTF_BRIDGE_IF_PREFIX + strconv.FormatInt(int64(req.GetPortbridge()), 10)
			exist, _ := macdb.MacIsExistData(namestr)
			if 1 == exist {
				reserr = macdb.MacPortbridgeDelDb()
			} else {
				//bridge未配置,返回提示信息
				cmdRtnInfo.Rtncode = 2
				cmdRtnInfo.Rtnstr += fmt.Sprintf("Error, bridge %d not exist!\n", req.GetPortbridge())
			}
		default:
			return &cmdRtnInfo, nil
		}
		if reserr != nil {
			rtnstring += "macif del db error! \n"
		}
		dbchange = true
	}

	if dbchange {
		err = macdb.MacDbChange(flushtime)
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(CFG_DB, MacCfgChan, strconv.FormatInt(flushtime, 10))
	}

	cmdRtnInfo.Rtnstr, _ = rtnstring, err
	cmdRtnInfo.Rtncode = 1
	return &cmdRtnInfo, nil
}

func (self *MacConfig) ShowMacIfCfg(ctx context.Context, req *pb.Showcfginfo) (*pb.Showrtninfo, error) {
	var showRtninfo pb.Showrtninfo
	//var outstr string

	/*show config,statistic use rtn buffer*/
	/*show buffer use file return data*/
	showRtninfo.Cmdcode = req.GetCmdcode()
	showRtninfo.Rtnstr = "qos ShowQosCarGloCfg command return!!!"

	//switch req.GetCmdid() {
	//}
	//switch(req.GetShowoption()){
	/*过滤等选项*/
	//}

	return &showRtninfo, nil
}

func (self *MacConfig) SetMacClearAct(ctx context.Context, req *pb.Macclearact) (*pb.Cfgrtninfo, error) {
	panic("implement me")
}

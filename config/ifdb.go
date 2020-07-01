package config

import (
	"github.com/go-redis/redis"
	pb "hstcmler/cml"
	"strconv"
)

type IfDb struct {
	dbclient *redis.Client
}

func NewIfDb() *IfDb {
	return new(IfDb)
}

//获取数据库连接
func (self *IfDb) InitDbClient() error {
	var err error
	//chenwei:链接数据库
	err = CmlDbReconnect(CFG_DB)
	if err != nil {
		return err
	}
	self.dbclient = cmldbinfo.cmldbclient[CmlDbNo[CFG_DB]]
	return err
}

func (self *IfDb) IfTypeNumSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	ifnum := req.GetIfnum()
	value := ifnum.GetType1() + "," + ifnum.GetNumber1() + "," + ifnum.GetSubnumber1()
	if "" != ifnum.GetType2() {
		value += ifnum.GetType2() + "," + ifnum.GetNumber2() + "," + ifnum.GetSubnumber2()
	}
	err := self.dbclient.HSet(key, FIELD_INTF_INDEX, value).Err()
	return err
}

func (self *IfDb) IfLookBackSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetIfnum().Number1
	err := self.dbclient.HSet(key, FIELD_INTF_INDEX, value).Err()
	return err
}

func (self *IfDb) IfNullSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetIfnum().Number1
	err := self.dbclient.HSet(key, FIELD_INTF_INDEX, value).Err()
	return err
}

func (self *IfDb) IfDescripSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetDescrip()
	err := self.dbclient.HSet(key, FIELD_INTF_DESCRIP, value).Err()
	return err
}

func (self *IfDb) IfClearConfInfoSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	//首先根据key获取hash的所有内容
	vals, err := self.dbclient.HGetAll(key).Result()
	//如果没有问题，则循环map
	if err == nil {
		for filed := range vals {
			//循环删除
			_, err := self.dbclient.HDel(key, filed).Result()
			//如果有错误，则直接退出
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (self *IfDb) IfProtoUpDelaySetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetProtoupdelay()
	err := self.dbclient.HSet(key, FIELD_INTF_PROTO_UPDELAY, value).Err()
	return err
}

func (self *IfDb) IfFlowStatIntervalSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetFlowstatinterval()
	err := self.dbclient.HSet(key, FIELD_INTF_STATIS_INTERVAL, value).Err()
	return err
}

func (self *IfDb) IfShutdownSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetAdminStatus()
	err := self.dbclient.HSet(key, FIELD_INTF_ADMIN_STATUS, value).Err()
	return err
}

func (self *IfDb) IfShutdownInfoSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetAdminStatus()
	err := self.dbclient.HSet(key, FIELD_INTF_ADMIN_STATUS, value).Err()
	return err
}

func (self *IfDb) IfShutdownNetSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetAdminStatus()
	err := self.dbclient.HSet(key, FIELD_INTF_ADMIN_STATUS, value).Err()
	return err
}

func (self *IfDb) IfShutdownTransmitSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetAdminStatus()
	err := self.dbclient.HSet(key, FIELD_INTF_ADMIN_STATUS, value).Err()
	return err
}

func (self *IfDb) IfShutdownReceiveSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetAdminStatus()
	err := self.dbclient.HSet(key, FIELD_INTF_ADMIN_STATUS, value).Err()
	return err
}

func (self *IfDb) IfFecModeBaseSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetAdminStatus()
	err := self.dbclient.HSet(key, FIELD_PORT_FEC_MODE, value).Err()
	return err
}

func (self *IfDb) IfFecModeNoneSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetAdminStatus()
	err := self.dbclient.HSet(key, FIELD_PORT_FEC_MODE, value).Err()
	return err
}

func (self *IfDb) IfFecModeRSSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetAdminStatus()
	err := self.dbclient.HSet(key, FIELD_PORT_FEC_MODE, value).Err()
	return err
}

func (self *IfDb) IfControlFlapSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	controlflap := req.GetControlflap()
	suppress := controlflap.GetSuppress()
	reuse := controlflap.GetReuse()
	ceiling := controlflap.GetCeiling()
	decayok := controlflap.GetDecayok()
	decayng := controlflap.GetDecayng()
	value := strconv.Itoa(int(suppress)) + "," + strconv.Itoa(int(reuse)) + "," + strconv.Itoa(int(ceiling)) + "," + strconv.Itoa(int(decayok)) + "," + strconv.Itoa(int(decayng))
	err := self.dbclient.HSet(key, FIELD_PORT_FLAP_INFO, value).Err()
	return err
}

func (self *IfDb) IfRangeSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	ifnum := req.GetIfnum()
	value := ifnum.GetType1() + "," + ifnum.GetNumber1()
	if "" != ifnum.GetType2() {
		value += ifnum.GetType2() + "," + ifnum.GetNumber2()
	}
	err := self.dbclient.HSet(key, "", value).Err()
	return err
}

func (self *IfDb) IfJumboframeSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	jumboinfo := req.GetJumboinfo()
	jumbersize := jumboinfo.GetJumbersize()
	jumberstatsize := jumboinfo.GetJumberstatsize()
	value := strconv.Itoa(int(jumbersize))
	if 0 != jumberstatsize {
		value += "," + strconv.Itoa(int(jumberstatsize))
	}
	err := self.dbclient.HSet(key, FIELD_PORT_JUMBER, value).Err()
	return err
}

func (self *IfDb) IfLoopbackSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetLoopinternal()
	err := self.dbclient.HSet(key, FIELD_PORT_LOOP_MODE, value).Err()
	return err
}

func (self *IfDb) IfPortModeLanSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetIfmode()
	err := self.dbclient.HSet(key, FIELD_PORT_LAN_WAN, value).Err()
	return err
}

func (self *IfDb) IfPortModeWanSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetIfmode()
	err := self.dbclient.HSet(key, FIELD_PORT_LAN_WAN, value).Err()
	return err
}

func (self *IfDb) IfPortMode10gSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetPortmode()
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_MODE, value).Err()
	return err
}

func (self *IfDb) IfPortModeGeSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetPortmode()
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_MODE, value).Err()
	return err
}

func (self *IfDb) IfPortMode25geSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	ifnum := req.GetIfnum()
	value := ifnum.GetType1() + "," + ifnum.GetNumber1()
	if "" != ifnum.GetType2() {
		value += ifnum.GetType2() + "," + ifnum.GetNumber2()
	}
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_MODE, value).Err()
	return err
}

func (self *IfDb) IfPortMode100geSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	ifnum := req.GetIfnum()
	value := ifnum.GetType1() + "," + ifnum.GetNumber1()
	if "" != ifnum.GetType2() {
		value += ifnum.GetType2() + "," + ifnum.GetNumber2()
	}
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_MODE, value).Err()
	return err
}

func (self *IfDb) IfPortSwitchSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSwitchmode()
	err := self.dbclient.HSet(key, FIELD_INTF_SWITCH_MODE, value).Err()
	return err
}

func (self *IfDb) IfPortSwitchBatchSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	ifnum := req.GetIfnum()
	value := ifnum.GetType1() + "," + ifnum.GetNumber1()
	if "" != ifnum.GetNumber2() {
		value += ifnum.GetNumber2()
	}
	err := self.dbclient.HSet(key, FIELD_INTF_SWITCH_MODE, value).Err()
	return err
}

func (self *IfDb) IfUpDelaySetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSetupdelay()
	err := self.dbclient.HSet(key, FIELD_PORT_UP_DELAY, value).Err()
	return err
}

func (self *IfDb) IfSpeed10SetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSpeedinfo().Speedvalue
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_NEGO, value).Err()
	return err
}

func (self *IfDb) IfSpeed100SetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSpeedinfo().Speedvalue
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_NEGO, value).Err()
	return err
}

func (self *IfDb) IfSpeed1000SetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSpeedinfo().Speedvalue
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_NEGO, value).Err()
	return err
}

func (self *IfDb) IfSpeed10000SetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSpeedinfo().Speedvalue
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_NEGO, value).Err()
	return err
}

func (self *IfDb) IfSpeed40000SetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSpeedinfo().Speedvalue
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_NEGO, value).Err()
	return err
}

func (self *IfDb) IfSpeedAuto10SetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSpeedinfo().Speedvalue
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_NEGO, value).Err()
	return err
}

func (self *IfDb) IfSpeedAuto100SetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSpeedinfo().Speedvalue
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_NEGO, value).Err()
	return err
}

func (self *IfDb) IfSpeedAuto1000SetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSpeedinfo().Speedvalue
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_NEGO, value).Err()
	return err
}

func (self *IfDb) IfSpeedAuto10000SetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetSpeedinfo().Speedvalue
	err := self.dbclient.HSet(key, FIELD_PORT_SPEED_NEGO, value).Err()
	return err
}

func (self *IfDb) IfSubIfSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetTrapdowndisable()
	err := self.dbclient.HSet(key, FIELD_INTF_DOWN_TRAP_DIS, value).Err()
	return err
}

func (self *IfDb) IfStatisticsSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetStatisticenable()
	err := self.dbclient.HSet(key, FIELD_INTF_STATISTIC_ENABLE, value).Err()
	return err
}

func (self *IfDb) IfMtuSetDb(req *pb.Ifbaseinfocfg) error {
	key := req.GetIfname()
	value := req.GetMtu()
	err := self.dbclient.HSet(key, FIELD_INTF_MTU, value).Err()
	return err
}

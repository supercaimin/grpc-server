package config

import (
	"github.com/go-redis/redis"
	pb "hstcmler/cml"
	"strconv"
)

var MAC_UNIT_value = map[int32]int32{
	1: 1000,
	2: 1000000,
	3: 1000000000,
}

var MAC_DETECT_ACTION = map[string]string{
	"low ":    "low ",
	"middle ": "middle ",
	"high  ":  "high  ",
}

type MacDb struct {
	dbclient *redis.Client
}

func NewMacDb() *MacDb {
	return new(MacDb)
}

//获取数据库连接
func (self *MacDb) InitDbClient() error {
	var err error
	//chenwei:链接数据库
	err = CmlDbReconnect(CFG_DB)
	if err != nil {
		return err
	}
	self.dbclient = cmldbinfo.cmldbclient[CmlDbNo[CFG_DB]]
	return err
}

//写db的接口,单独放一个对应文件中
func (self *MacDb) MacAgeTimeSetDb(agetime int32) error {
	data := strconv.FormatInt(int64(agetime), 10)
	err := self.dbclient.HSet(KEY_MAC_GLOBAL, FIELD_MAC_AGTIME, data).Err()
	return err
}

//黑洞mac写DB
func (self *MacDb) MacBlackholeSetDb(blackmac *pb.Macblackhole) error {
	var field, data string
	//fieldname BLACK| +macstr
	field = FIELD_MAC_BLACK_PREFIX + blackmac.GetMac()
	data = strconv.FormatInt(int64(blackmac.GetVlanid()), 10)
	err := self.dbclient.HSet(KEY_MAC_GLOBAL, field, data).Err()
	return err
}

//Flapingagetime
func (self *MacDb) MacFlapingagetimeSetDb(flapingagetime int32) error {
	data := strconv.FormatInt(int64(flapingagetime), 10)
	err := self.dbclient.HSet(KEY_MAC_GLOBAL, FIELD_MAC_FLAP_INFO, data).Err()
	return err
}

//写db的接口,单独放一个对应文件中
func (self *MacDb) MacFlapdetectSetDb(flapdetect string) error {
	data := flapdetect
	err := self.dbclient.HSet(KEY_MAC_GLOBAL, FIELD_MAC_FLAP_INFO, data).Err()
	return err
}

//静态MAC写db
func (self *MacDb) MacStaticSetDb(staticmac *pb.Macstatic) error {
	var dkey, field, data string
	//fieldname BLACK| +macstr
	field = FIELD_MAC_ITEM_PREFIX + staticmac.GetMac()
	data = strconv.FormatInt(int64(staticmac.GetVlanid()), 10)
	dkey = KEY_MAC_STATIC_PREFIX + staticmac.GetIfname()
	err := self.dbclient.HSet(dkey, field, data).Err()
	return err
}

func (self *MacDb) MaclearndisSetDb(actionStr string) error {
	data := actionStr
	err := self.dbclient.HSet(FIELD_IF_MAC_LEARN, FIELD_IF_MAC_LEARN, data).Err()
	return err
}

func (self *MacDb) MaclimitdisSetDb(actionStr string) error {
	data := actionStr
	err := self.dbclient.HSet(FIELD_IF_MAC_LIMIT, FIELD_IF_MAC_LIMIT, data).Err()
	return err
}

func (self *MacDb) MacPortbridgeSetDb(portbridge int32) error {
	data := strconv.FormatInt(int64(portbridge), 10)
	err := self.dbclient.HSet(FIELD_IF_MAC_BRIDGE_ENABLE, FIELD_IF_MAC_BRIDGE_ENABLE, data).Err()
	return err
}

func (self *MacDb) MacAgeTimeDelDb() error {
	return self.dbclient.HDel(KEY_MAC_GLOBAL, FIELD_MAC_AGTIME).Err()
}

func (self *MacDb) MacBlackholeDelDb(str string) error {
	return self.dbclient.HDel(KEY_MAC_GLOBAL, FIELD_MAC_BLACK_PREFIX+str).Err()
}

//Flapingagetime
func (self *MacDb) MacFlapingagetimeDelDb() error {
	return self.dbclient.HDel(KEY_MAC_GLOBAL, FIELD_MAC_FLAP_INFO).Err()
}

//写db的接口,单独放一个对应文件中
func (self *MacDb) MacFlapdetectDelDb(str string) error {
	return self.dbclient.HDel(KEY_MAC_GLOBAL, FIELD_MAC_FLAP_INFO+str).Err()
}

//静态MAC写db
func (self *MacDb) MacStaticDelDb(str1, str2 string) error {
	return self.dbclient.HDel(KEY_MAC_STATIC_PREFIX+str1, FIELD_MAC_ITEM_PREFIX+str2).Err()
}

func (self *MacDb) MaclearndisDelDb() error {
	return self.dbclient.HDel(FIELD_IF_MAC_LEARN, FIELD_IF_MAC_LEARN).Err()
}

func (self *MacDb) MaclimitdisDelDb() error {
	return self.dbclient.HDel(FIELD_IF_MAC_LIMIT, FIELD_IF_MAC_LIMIT).Err()
}

func (self *MacDb) MacPortbridgeDelDb() error {
	return self.dbclient.HDel(FIELD_IF_MAC_BRIDGE_ENABLE, FIELD_IF_MAC_BRIDGE_ENABLE).Err()
}

func (self *MacDb) MacIsExistData(str string) (int64, error) {
	return self.dbclient.Exists(str).Result()
}

func (self *MacDb) MacDbChange(flushtime int64) error {
	return self.dbclient.HSet(KEY_MAC_GLOBAL, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
}

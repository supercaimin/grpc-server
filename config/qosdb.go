// qosdb.go
// qos module cml set db process

package config

import (
	"fmt"
	"strconv"
	"strings"

	pb "hstcmler/cml"

	"github.com/go-redis/redis"
)

/*qos 队列等信息定义*/
const (
	qos_queue_index_min int = -1 /*0-7*/
	qos_queue_index_max     = 8  /*0-7*/
)

const (
	PUB_QOSCHAN    string = "qos.config"
	PUB_MQCCHANCFG        = "mqc.config"
)

var QOS_COLOR_value = map[string]int32{
	"green":  0,
	"yellow": 1,
	"red":    2,
}

var QOS_COLOR_ACTION = map[string]string{
	"discard": "DROP",
	"pass":    "FORWARD",
}

/*cir,pir,cbs,pbs的单位值定义*/
/*1:k, 2:m,3:g ;写入DB时需要统一转为bytes*/
var QOS_UNIT_value = map[int32]int32{
	1: 1000,
	2: 1000000,
	3: 1000000000,
}

/*qos diffserv 8021p 映射关系*/
/*8021p in bound,默认映射*/
/*
8021p  phb-behaiver   color
0      BE    green
1      AF1   green
2      AF2   green
3      AF3   green
4      AF4   green
5      EF    green
6      CS6   green
7      CS7   green
*/

/*
8021P out bound 默认映射
phb  color   8021p
BE   green    0
BE   yellow   0
BE   red      0
AF1  green    1
AF1  yellow   1
AF1  red      1
AF2  green    2
AF2  yellow   2
AF2  red      2
AF3  green    3
AF3  yellow   3
AF3  red      3
AF4  green    4
AF4  yellow   4
AF4  red      4
EF   green    5
EF   yellow   5
EF   red      5
CS6  green    6
CS6  yellow   6
CS6  red      6
CS7  green    7
CS7  yellow   7
CS7  red      7

*/

/*qos 映射表数据库操作,各参数间的冲突在配置处理中已检查过,该部分只处理数据库读写的操作*/
/*配置数据与已有的配置数据（存在db中，不放内存）冲突的检查也在配置处理部分进行*/
/*config-db写入后,orchagent会同步到app-ASICDB,所以数据表的key,fied等得一致*/
func QosMapSetRedis(dbname string, profile *pb.Qosdiffservmap) (string, error) {
	var index int32
	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring, colorstr, classstr, field string = "", "", "", ""
	var qos8021pmap []*pb.Qos8021Pmap
	var dscpmap []*pb.Dscpmap

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	/*取diffserv name,相关配置在diffserv name下*/
	//domainname = profile.GetDiffservName()

	/*qos 8021p in map*/
	/*最多可以配8个,0-7 队列*/
	qos8021pmap = profile.GetIn8021Pmap()
	for _, datainfo := range qos8021pmap {
		/*index(0-7), class(BE,AF1-4,CS6,CS7), color(green,yellow,red)*/
		colorstr = datainfo.GetColor()
		classstr = datainfo.GetServclass()
		index = datainfo.GetQos8021Pvalue()
		field = FIELD_DIFFSERV_IN8021P + fmt.Sprintf("%d", index)
		/*按DB格式写入数据库*/
		/*若原来有配置直接覆盖*/
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, field, (colorstr + " " + classstr)).Err()
		if err != nil {
			rtnstring += "qos in map info set db error! \n"
			break
		}
		dbchange = true
	}

	/*qos 8021p out map*/
	/*最多可以配8个,0-7 队列*/
	qos8021pmap = profile.GetOut8021Pmap()
	for _, datainfo := range qos8021pmap {
		/*index(0-7), class(BE,AF1-4,CS6,CS7), color(green,yellow,red)*/
		colorstr = datainfo.GetColor()
		classstr = datainfo.GetServclass()
		index = datainfo.GetQos8021Pvalue()

		field = FIELD_DIFFSERV_OUT8021P + fmt.Sprintf("%d", index)
		/*按DB格式写入数据库*/
		/*若原来有配置直接覆盖*/
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, field, (colorstr + " " + classstr)).Err()
		if err != nil {
			rtnstring += "server info set db error! \n"
			break
		}
		dbchange = true
	}

	/*qos dscp in map*/
	/*最多可以配8个,0-7 队列*/
	dscpmap = profile.GetIndscpmap()
	for _, datainfo := range dscpmap {
		/*index(0-7), class(BE,AF1-4,CS6,CS7), color(green,yellow,red)*/
		colorstr = datainfo.GetColor()
		classstr = datainfo.GetServclass()
		index = datainfo.GetDscpvalue()

		field = FIELD_DIFFSERV_INDSCP + fmt.Sprintf("%d", index)
		/*按DB格式写入数据库*/
		/*若原来有配置直接覆盖*/
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, field, (colorstr + " " + classstr)).Err()
		if err != nil {
			rtnstring += "server info set db error! \n"
			break
		}
		dbchange = true
	}

	/*qos dscp OUT map*/
	/*最多可以配8个,0-7 队列*/
	dscpmap = profile.GetOutdscpmap()
	for _, datainfo := range dscpmap {
		/*index(0-7), class(BE,AF1-4,CS6,CS7), color(green,yellow,red)*/
		colorstr = datainfo.GetColor()
		classstr = datainfo.GetServclass()
		index = datainfo.GetDscpvalue()

		field = FIELD_DIFFSERV_OUTDSCP + fmt.Sprintf("%d", index)
		/*按DB格式写入数据库*/
		/*若原来有配置直接覆盖*/
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, field, (colorstr + " " + classstr)).Err()
		if err != nil {
			rtnstring += "server info set db error! \n"
			break
		}
		dbchange = true
	}
	/*EXP,MPLS ...*/

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

/*qos diffserv删除,直接根据servname删除该Key下的所有内容*/
func QosMapUndoSetRedis(dbname string, profile *pb.Qosdiffservmap) (string, error) {

	//var index, i int32
	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring, domainname string = "", ""

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	/*取diffserv name,相关配置在diffserv name下*/
	domainname = profile.GetDiffservName()

	/*组织key,直接删除该key以及key下的所有内容*/
	err = dbclient.HDel(KEY_QOS_DIFFSERVER, domainname).Err()
	if err != nil {
		rtnstring += "qos diff del db error! \n"
	}
	dbchange = true

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

/*qos diffserv在接口上的配置应用*/
func QosDiffIfSetRedis(dbname string, profile *pb.Qosdiffservapply) (string, error) {

	var index, prevalue int32
	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring, field string = "", ""
	var localpre []*pb.Localpremap
	//var diffserv, iflist string

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	//diffserv = profile.GetTrustdiffserv()
	//iflist = profile.GetIflist()

	/*local pre in map*/
	/*最多可以配8个,0-7 队列*/
	localpre = profile.GetLocalprecfg()
	for _, datainfo := range localpre {
		prevalue = datainfo.GetLocalPre()
		index = datainfo.GetQueueindex()

		field = FIELD_DIFFSERV_IN8021P + fmt.Sprintf("%d", index)
		/*按DB格式写入数据库*/
		/*若原来有配置直接覆盖*/
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, field, fmt.Sprintf("%d", prevalue)).Err()
		if err != nil {
			rtnstring += "qos in map info set db error! \n"
			break
		}
		dbchange = true
	}

	/*EXP,MPLS ...*/

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

func QosDiffIfUndoSetRedis(dbname string, profile *pb.Qosdiffservapply) (string, error) {

	var index, prevalue int32
	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring, field string = "", ""
	var localpre []*pb.Localpremap
	//var diffserv, iflist string

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	//diffserv = profile.GetTrustdiffserv()
	//iflist = profile.GetIflist()

	/*local pre in map*/
	/*最多可以配8个,0-7 队列*/
	localpre = profile.GetLocalprecfg()
	for _, datainfo := range localpre {
		prevalue = datainfo.GetLocalPre()
		index = datainfo.GetQueueindex()

		field = FIELD_DIFFSERV_IN8021P + fmt.Sprintf("%d", index)
		/*按DB格式写入数据库*/
		/*若原来有配置直接覆盖*/
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, field, fmt.Sprintf("%d", prevalue)).Err()
		if err != nil {
			rtnstring += "qos in map info set db error! \n"
			break
		}
		dbchange = true
	}

	/*EXP,MPLS ...*/

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

/*qos car 全局配置应用*/
func QosCarGloSetRedis(dbname string, profile *pb.Qoscarglobalcfg) (string, error) {

	var datavalue int32
	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring, field, carname string = "", "", ""
	var carinfo []*pb.Qoscar

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	carinfo = profile.GetQoscarcfg()
	if nil != carinfo {
		for _, datainfo := range carinfo {
			carname = KEY_QOS_POLICER + datainfo.GetCarname()

			_ = dbclient.HSet(carname, FIELD_METERTYPE, "BYTES").Err()
			_ = dbclient.HSet(carname, FIELD_MODE, "SR_TCM").Err()
			_ = dbclient.HSet(carname, FIELD_COLOR_SOURCE, "AWARE").Err()
			_ = dbclient.HSet(carname, FIELD_GREEN_PACKET_ACTION, "FORWARD").Err()

			datavalue = datainfo.GetCirvalue()
			if 0 < datavalue {
				/*转换为byte,cbs基本单位是bit*/
				datavalue = datavalue / 8 * QOS_UNIT_value[datainfo.GetCirunit()]
				_ = dbclient.HSet(carname, FIELD_CIR, fmt.Sprintf("%d", datavalue)).Err()

				datavalue = datainfo.GetCirCbs()
				/*转换为byte,cbs基本单位是byte*/
				datavalue = datavalue * QOS_UNIT_value[datainfo.GetCcbsunit()]
				_ = dbclient.HSet(carname, FIELD_CBS, fmt.Sprintf("%d", datavalue)).Err()
			}

			datavalue = datainfo.GetPirvalue()
			if 0 < datavalue {
				/*转换为byte,cbs基本单位是bit*/
				datavalue = datavalue / 8 * QOS_UNIT_value[datainfo.GetPirunit()]
				_ = dbclient.HSet(carname, FIELD_PIR, fmt.Sprintf("%d", datavalue)).Err()

				datavalue = datainfo.GetPirPbs()
				/*转换为byte,cbs基本单位是byte*/
				datavalue = datavalue * QOS_UNIT_value[datainfo.GetPpbsunit()]
				_ = dbclient.HSet(carname, FIELD_PBS, fmt.Sprintf("%d", datavalue)).Err()
			}

			field = profile.GetYellowaction()
			if "" != field {
				_ = dbclient.HSet(carname, FIELD_YELLOW_PACKET_ACTION, QOS_COLOR_ACTION[field]).Err()
			}

			field = profile.GetRedaction()
			if "" != field {
				_ = dbclient.HSet(carname, FIELD_RED_PACKET_ACTION, QOS_COLOR_ACTION[field]).Err()
			}
		}
	}

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

func QosCarGloUndoSetRedis(dbname string, profile *pb.Qoscarglobalcfg) (string, error) {

	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring, carname string = "", ""
	var carinfo []*pb.Qoscar

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	carinfo = profile.GetQoscarcfg()
	if nil != carinfo {
		for _, datainfo := range carinfo {
			carname = KEY_QOS_POLICER + datainfo.GetCarname()
			err = dbclient.Del(carname).Err()
			if nil != err {
				rtnstring += "del qos car failed! \n"
			}
		}
	}

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

/*qos car if配置应用*/
func QosCommIfSetRedis(dbname string, profile *pb.Qoscommifcfg) (string, error) {

	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring string = ""

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

func QosCommIfUndoSetRedis(dbname string, profile *pb.Qoscommifcfg) (string, error) {

	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring string = ""

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

/*qos label mode配置应用*/
func QosLableModeCfgSetRedis(dbname string, profile *pb.Labelqosmode) (string, error) {

	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring string = ""

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

func QosLableModeCfgUndoSetRedis(dbname string, profile *pb.Labelqosmode) (string, error) {

	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring string = ""

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

/*qos phb info 配置应用*/
func QosPhbInfoCfgSetRedis(dbname string, profile *pb.Qosphbcfg) (string, error) {

	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring string = ""

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

func QosPhbInfoCfgUndoSetRedis(dbname string, profile *pb.Qosphbcfg) (string, error) {

	var err error = nil
	var dbchange bool = false
	var flushtime int64
	var dbclient *redis.Client
	var rtnstring string = ""

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect qos db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatetime()

	if dbchange {
		err = dbclient.HSet(KEY_QOS_DIFFSERVER, FIELD_UPDATETIME, strconv.FormatInt(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_QOSCHAN, strconv.FormatInt(flushtime, 10))
	}

	return rtnstring, err
}

/*应该按GRPC-PROTOBUF模式，逐个读出参数并按protobuf格式填写后返回*/
/*对于CLI,在CLI端根据参数新组织为命令字符串显示,取巧的方式是将命令配置字符串也写库*/
/*show显示时也同时将命令字符串通过protobuf带回,省去cli的反向组织*/
/*对于web-gui 或 netconf等终端则根据protobuf的数据反向显示*/
/*grpc的show接口需要重新定义,以满足protobuf的返回模式*/
func QosShowConfig(dbname string, rtnstr *string) {

	var dbclient *redis.Client
	var err error
	var fieldvalue, outstr, matchkey string
	var keys []string

	err = CmlDbReconnect(dbname)
	if err != nil {
		*rtnstr = "Connect qos db failed!\n"
		return
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	outstr = "#qos config, last change:"

	fieldvalue, err = dbclient.HGet(KEY_QOS_POLICER_GLOBAL, FIELD_UPDATETIME).Result()
	if "" != fieldvalue {
		fmt.Println(fieldvalue)
		outstr += fieldvalue + " msec\n"
	} else {
		outstr += "0 msec\n"
	}

	/*读取显示qos car*/
	matchkey = KEY_QOS_POLICER + "*"
	keys, err = dbclient.Keys(matchkey).Result()
	if err != nil {
		/*读取失败,显示获取数据失败*/
		outstr += "Get qos car config failed!\n"
	} else {
		/*逐个读取显示,适用于命令行配置的显示*/
		for i, carname := range keys {
			fieldvalue, err = dbclient.HGet(keys[i], FIELD_COMMANDLINE).Result()
			if nil == err {
				outstr += "  qos car " + strings.TrimLeft(carname, KEY_QOS_POLICER) + fieldvalue + "\n"
			}
		}
	}

	/*读取显示color action等*/
	fieldvalue, err = dbclient.HGet(KEY_QOS_POLICER, FIELD_YELLOW_PACKET_ACTION).Result()
	if "" != fieldvalue {
		outstr += "  qos car yellow " + fieldvalue + "\n"
	}

	fieldvalue, err = dbclient.HGet(KEY_QOS_POLICER, FIELD_RED_PACKET_ACTION).Result()
	if "" != fieldvalue {
		outstr += "  qos car red " + fieldvalue + "\n"
	}

	*rtnstr = outstr

}

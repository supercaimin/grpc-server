// syslogconfig.go
// syslog cml process

package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	pb "hstcmler/cml"

	"github.com/go-redis/redis"
)

const (
	PUB_SYSLOGCHAN string = "syslog.config"
	PUB_ACLCHANCFG        = "acl.config"
)

/*
syslog config db
   "syslog" "console" level-value
   "syslog" "monitor" level-value
    ......
   "syslog" "source" source-struct
*/
const (
	console_defaut_level   int = 6
	monitor_defaut_level       = 6
	trap_defaut_level          = 0 /*invalid max*/
	server_defaut_level        = 0 /*invalid max*/
	persis_defaut_level        = 5
	syslog_defaut_bufsize      = 2048 /*2k record*/
	syslog_defaut_memsize      = 4    /*4M*/
	syslog_bufsize_min         = 256
	syslog_bufsize_max         = 10240
	syslog_memsize_min         = 1
	syslog_memsize_max         = 32
	syslog_default_filedur     = 30 /*默认保存文件周期为30天*/
	syslog_filedur_min         = 1
	syslog_filedur_max         = 60
	syslog_level_max           = 8 /*0-7 is valid, 8 is invalid*/
)

/*对persistent等，若配置的level与默认值相等也不删除数据库中的节点,因为考虑以后版本升级的兼容性
等。如果删除，且后续版本改变了默认级别，则将导致升级前后版本表现不一致。保留则可通过db兼容保持一致*/
/*返回的returncode需要定义,包含数据库未连接(可通过err反馈具体信息),参数错误(无法通过err反馈）等*/
/*传入的参数level等在KLISH命令行就需要完成合法性检查,本函数只做有效性检查*/
func SyslogSetRedis(dbname string, profile *pb.SyslogConfProfile) (string, error) {

	var loglevel pb.LOG_LEVEL
	var intsize int32
	var err error = nil
	var dbchange bool = false
	var flushtime uint64
	var dbclient *redis.Client
	var rtnstring, serverstr, serverfield string = "", "", ""
	var syslogserver []*pb.SyslogServerInfo
	var syslogsource *pb.SyslogSoureInfo

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect syslog db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatedAt()

	loglevel = profile.GetConsoleLevel()
	/*参数合法的就都有效,在此检查用于判断是否进行了配置.无效值表示未配置.下同*/
	if syslog_level_max > loglevel {
		err = SyslogDbLevelSet(dbclient, KEY_SYSLOG, FIELD_CONSOLE, loglevel)
		if err != nil {
			rtnstring += "console set db error! \n"
		}
		dbchange = true
	}

	/*default no config trap, only use no command to set*/
	loglevel = profile.GetMonitorLevel()
	if syslog_level_max > loglevel {
		err = SyslogDbLevelSet(dbclient, KEY_SYSLOG, FIELD_MONITOR, loglevel)
		if err != nil {
			rtnstring += "monitor set db error! \n"
		}
		dbchange = true
	}

	/*default no config trap, only use no command to set*/
	loglevel = profile.GetTrapLevel()
	if syslog_level_max > loglevel {
		err = SyslogDbLevelSet(dbclient, KEY_SYSLOG, FIELD_TRAP, loglevel)
		if err != nil {
			rtnstring += "trap set db error! \n"
		}
		dbchange = true
	}

	/*default no config loghost, only use no command to set*/
	loglevel = profile.GetServerLevel()
	if syslog_level_max > loglevel {
		err = SyslogDbLevelSet(dbclient, KEY_SYSLOG, FIELD_SERVER, loglevel)
		if err != nil {
			rtnstring += "loghost set db error! \n"
		}
		dbchange = true
	}

	loglevel = profile.GetPersisLevel()
	if syslog_level_max > loglevel {
		if persis_defaut_level == loglevel {
			err = dbclient.HDel(KEY_SYSLOG, FIELD_PERSISTENT).Err()
		} else {
			err = SyslogDbLevelSet(dbclient, KEY_SYSLOG, FIELD_PERSISTENT, loglevel)
			if err != nil {
				rtnstring += "persistent set db error! \n"
			}
			dbchange = true
		}
	}

	intsize = profile.GetLogbufsize()
	if (syslog_bufsize_min <= intsize) && (syslog_bufsize_max >= intsize) {
		if syslog_defaut_bufsize == intsize {
			/*如果配置的是默认值,执行恢复默认值的操作*/
			err = dbclient.HDel(KEY_SYSLOG, FIELD_BUFFERSIZE).Err()
		} else {
			err = dbclient.HSet(KEY_SYSLOG, FIELD_BUFFERSIZE, intsize).Err()
		}
		if err != nil {
			rtnstring += "buffersize set db error! \n"
		}
		dbchange = true
	}

	intsize = profile.GetLogmemsize()
	if (syslog_memsize_min <= intsize) && (syslog_memsize_max >= intsize) {
		if syslog_defaut_memsize == intsize {
			/*如果配置的是默认值,执行恢复默认值的操作*/
			err = dbclient.HDel(KEY_SYSLOG, FIELD_STORAGESIZE).Err()
		} else {
			err = dbclient.HSet(KEY_SYSLOG, FIELD_STORAGESIZE, intsize).Err()
		}
		if err != nil {
			rtnstring += "storagesize set db error! \n"
		}
		dbchange = true
	}

	intsize = profile.GetLogfileduration()
	if (syslog_filedur_min <= intsize) && (syslog_filedur_max >= intsize) {
		if syslog_default_filedur == intsize {
			/*如果配置的是默认值,执行恢复默认值的操作*/
			err = dbclient.HDel(KEY_SYSLOG, FIELD_FILEDURATION).Err()
		} else {
			err = dbclient.HSet(KEY_SYSLOG, FIELD_FILEDURATION, intsize).Err()
		}
		if err != nil {
			rtnstring += "file duration reset db failed!\n"
		}
		dbchange = true
	}

	/*field use: serverip1 , serverip2; max 2 server*/
	/*syslog server*/
	var serverdata string
	syslogserver = profile.GetLogserver()
	for _, serverinfo := range syslogserver {
		if serverinfo.GetAddrtype() < 20 {
			serverfield = FIELD_SYSLOG_SERVER1
			serverdata, _ = dbclient.HGet(KEY_SYSLOG_SEVER, FIELD_SYSLOG_SERVER2).Result()
		} else {
			serverfield = FIELD_SYSLOG_SERVER2
			serverdata, _ = dbclient.HGet(KEY_SYSLOG_SEVER, FIELD_SYSLOG_SERVER1).Result()
		}
		/*check if exist*/
		if 0 == serverinfo.GetPort() {
			serverstr = fmt.Sprintf("vrf:%s,type:%d,ip:%s,port:514", serverinfo.GetVrfname(),
				serverinfo.GetAddrtype()%10, serverinfo.GetIpaddr())
		} else {
			serverstr = fmt.Sprintf("vrf:%s,type:%d,ip:%s,port:%d", serverinfo.GetVrfname(),
				serverinfo.GetAddrtype()%10, serverinfo.GetIpaddr(), serverinfo.GetPort())
		}
		/*若原来有配置直接覆盖,但需要检查server1和2是否相同*/
		if "" != serverdata && (serverdata == serverstr) {
			/*server1 and 2 is same,*/
			rtnstring += "Exist same config, conflict!\n"
		} else {
			err = dbclient.HSet(KEY_SYSLOG_SEVER, serverfield, serverstr).Err()
			if err != nil {
				rtnstring += "server info set db error! \n"
			}
			dbchange = true
		}
	}

	/*syslog source*/
	syslogsource = profile.GetLogsourceinfo()
	if nil != syslogsource {
		srcip := syslogsource.GetIpaddr()
		/*0.0.0.0 equal no command, 255.255.255.255 is not config*/
		if srcip != "255.255.255.255" {
			if srcip == "0.0.0.0" {
				err = dbclient.HDel(KEY_SYSLOG, FIELD_SORCEIP).Err()
			} else {
				err = dbclient.HSet(KEY_SYSLOG, FIELD_SORCEIP, srcip).Err()
			}
			if err != nil {
				rtnstring += "source address set db error! \n"
			}
			dbchange = true
		}
		srcport := syslogsource.GetPort()
		/*0 equal no command, 65535 is not config*/
		if srcport != 65535 {
			if srcport == 0 {
				err = dbclient.HDel(KEY_SYSLOG, FIELD_SOURCEPORT).Err()
			} else {
				err = dbclient.HSet(KEY_SYSLOG, FIELD_SOURCEPORT, srcport).Err()
			}
			if err != nil {
				rtnstring += "source port set db error! \n"
			}
			dbchange = true
		}

	}

	if dbchange {
		err = dbclient.HSet(KEY_SYSLOG, FIELD_UPDATETIME, strconv.FormatUint(flushtime, 10)).Err()
		if err != nil {
			rtnstring += "updatetime set db error! \n"
		}
		CmlDbCfgPubstring(dbname, PUB_SYSLOGCHAN, strconv.FormatUint(flushtime, 10))
	}

	return rtnstring, err
}

func SyslogPortValid(port int32) bool {

	/*还应该排除知名端口*/
	if 65535 > port {
		return true
	} else {
		return false
	}
}

func SyslogUndoSetRedis(dbname string, profile *pb.SyslogConfProfile) (string, error) {

	var loglevel pb.LOG_LEVEL
	var intsize int32
	var err error = nil
	var dbchange bool = false
	var flushtime uint64
	var dbclient *redis.Client
	var rtnstring string = ""

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect syslog db failed!\n"
		return rtnstring, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	flushtime = profile.GetUpdatedAt()

	loglevel = profile.GetConsoleLevel()
	/*no command ,valid is ok*/
	if pb.LOG_LEVEL_LOG_LEVEL_DEBUG >= loglevel {
		err = SyslogDbLevelDel(dbclient, KEY_SYSLOG, FIELD_CONSOLE)
		if err != nil {
			rtnstring += "  console del db failed!\n"
		}
		dbchange = true
	}

	loglevel = profile.GetMonitorLevel()
	/*no command ,valid is ok*/
	if pb.LOG_LEVEL_LOG_LEVEL_DEBUG >= loglevel {
		err = dbclient.HDel(KEY_SYSLOG, FIELD_MONITOR).Err()
		if err != nil {
			rtnstring += "  monitor del db failed!\n"
		}
		dbchange = true
	}

	loglevel = profile.GetTrapLevel()
	/*no command ,valid is ok*/
	if pb.LOG_LEVEL_LOG_LEVEL_DEBUG >= loglevel {
		err = dbclient.HDel(KEY_SYSLOG, FIELD_TRAP).Err()
		if err != nil {
			rtnstring += "  trap del db failed!\n"
		}
		dbchange = true
	}

	loglevel = profile.GetServerLevel()
	/*no command ,valid is ok*/
	if pb.LOG_LEVEL_LOG_LEVEL_DEBUG >= loglevel {
		err = dbclient.HDel(KEY_SYSLOG, FIELD_SERVER).Err()
		if err != nil {
			rtnstring += "  server level del db failed!\n"
		}
		dbchange = true
	}

	loglevel = profile.GetPersisLevel()
	/*no command ,valid is ok*/
	if pb.LOG_LEVEL_LOG_LEVEL_DEBUG >= loglevel {
		err = dbclient.HDel(KEY_SYSLOG, FIELD_PERSISTENT).Err()
		if err != nil {
			rtnstring += "  persistent del db failed!\n"
		}
		dbchange = true
	}

	intsize = profile.GetLogbufsize()
	/*0无效值,代表未配置*/
	if (syslog_bufsize_max >= intsize) && (0 < intsize) {
		err = dbclient.HDel(KEY_SYSLOG, FIELD_BUFFERSIZE).Err()
		if err != nil {
			rtnstring += "  buffer reset db failed!\n"
		}
		dbchange = true
	}

	intsize = profile.GetLogmemsize()
	/*0无效值,代表未配置*/
	if (syslog_memsize_max >= intsize) && (0 < intsize) {
		err = dbclient.HDel(KEY_SYSLOG, FIELD_STORAGESIZE).Err()
		if err != nil {
			rtnstring += "  memsize reset db failed!\n"
		}
		dbchange = true
	}

	intsize = profile.GetLogfileduration()
	/*0无效值,代表未配置*/
	if (syslog_default_filedur >= intsize) && (0 < intsize) {
		err = dbclient.HDel(KEY_SYSLOG, FIELD_FILEDURATION).Err()
		if err != nil {
			rtnstring += "  memsize reset db failed!\n"
		}
		dbchange = true
	}

	/*field use: serverip1 , serverip2; max 2 server*/
	/*syslog server*/
	syslogserver := profile.GetLogserver()
	var serverfield string
	for _, serverinfo := range syslogserver {
		if serverinfo.GetAddrtype() < 20 {
			serverfield = FIELD_SYSLOG_SERVER1
		} else {
			serverfield = FIELD_SYSLOG_SERVER2
		}
		fmt.Printf("delete server:%s", serverfield)
		/*check if exist*/
		err = dbclient.HDel(KEY_SYSLOG_SEVER, serverfield).Err()
		if err != nil {
			rtnstring += "server info reset db error! \n"
		}
		dbchange = true
	}

	/*syslog source & port*/
	syslogsource := profile.GetLogsourceinfo()
	if nil != syslogsource {
		/*0.0.0.0 is no,255.255.255.255 is not config*/
		if syslogsource.GetIpaddr() == "0.0.0.0" {
			err = dbclient.HDel(KEY_SYSLOG, FIELD_SORCEIP).Err()
			if err != nil {
				rtnstring += "source reset db error! \n"
			}
			dbchange = true
		}
		/*0 is no,65535 is not config*/
		if syslogsource.GetPort() == 0 {
			err = dbclient.HDel(KEY_SYSLOG, FIELD_SOURCEPORT).Err()
			if err != nil {
				rtnstring += "source reset db error! \n"
			}
			dbchange = true
		}
	}

	if dbchange {

		err = dbclient.HSet(KEY_SYSLOG, FIELD_UPDATETIME, strconv.FormatUint(flushtime, 10)).Err()
		if err != nil {
			/*retry or alarm?*/
			fmt.Println("db flush updatetime failed ", err)
		}
		CmlDbCfgPubstring(dbname, "syslog.config", strconv.FormatUint(flushtime, 10))
	}

	return rtnstring, err

}

/*修改告警码,写数据库 Key:AlarmNew|module */
/*只修改告警码对应的级别,不需要修改告警码产生位置,位置是业务发送时填写的*/
func LogLevelSetRedis(dbname string, profile *pb.LogLevelProfile) (string, error) {
	var err error = nil
	var dkey, okey, data, flushtime string
	var splitdata []string
	var dbclient *redis.Client
	var rtnstring string = ""

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect syslog db failed!\n"
		return rtnstring, err
	}

	flushtime = strconv.FormatUint(profile.GetUpdatedAt(), 10)

	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]
	/*模块名统一转为小写*/
	okey = KEY_ALARM_LEVEL_ORIGINAL + strings.ToLower(profile.GetModulename())
	dkey = KEY_ALARM_LEVEL_NEW + strings.ToLower(profile.GetModulename())

	/*先读取原始的数据,取出level:position*/
	data, err = dbclient.HGet(okey, fmt.Sprintf("%d", profile.GetCode())).Result()
	if err != nil {
		rtnstring = "Check alarm code failed!\n"
		fmt.Printf("%s\n", rtnstring)
		return rtnstring, err
	}
	/*将oldlevel替换为newlevel,position不变*/
	splitdata = strings.Split(data, DATA_SPLIT_STR)
	splitdata[0] = fmt.Sprintf("%d", profile.GetLevel())

	data = strings.Join(splitdata, DATA_SPLIT_STR)
	/*按DB格式写入数据库 code-level,updatetime 域*/
	err = dbclient.HSet(dkey, fmt.Sprintf("%d", profile.GetCode()), data).Err()
	if err != nil {
		rtnstring = fmt.Sprintf("Db write failed,key:%s code:%d!\n", dkey, profile.GetCode())
	} else {
		dbclient.HSet(KEY_ALARM_UPDATEINFO, FIELD_ALARM_UPDATETIME, flushtime)
	}

	CmlDbCfgPubstring(dbname, PUB_SYSLOGCHAN, flushtime)

	return rtnstring, err
}

func LogLevelUndoSetRedis(dbname string, profile *pb.LogLevelProfile) (string, error) {
	var err error = nil
	var dkey, flushtime string
	var dbclient *redis.Client
	var rtnstring string = ""

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstring = "Connect syslog db failed!\n"
		return rtnstring, err
	}

	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]
	dkey = KEY_ALARM_LEVEL_NEW + strings.ToLower(profile.GetModulename())

	/*按DB格式写入数据库 code-level,updatetime 域*/
	err = dbclient.HDel(dkey, fmt.Sprintf("%d", profile.GetCode())).Err()
	if err != nil {
		rtnstring = fmt.Sprintf("Db del failed,key:%s!\n", dkey)
	} else {
		/*更新updatetime*/
		flushtime = strconv.FormatUint(profile.GetUpdatedAt(), 10)
		dbclient.HSet(KEY_ALARM_UPDATEINFO, FIELD_ALARM_UPDATETIME, flushtime)
	}

	CmlDbCfgPubstring(dbname, PUB_SYSLOGCHAN, flushtime)

	return rtnstring, err
}

/*修改后的*/
/*对每个模块先把修改后的读出存在一个map中,如果显示所有的就再读出原来的然后组织显示*/
/*以原来的数据为所以逐个组织,组织时查询map中是否有修改的,若有增加修改的显示*/
func LogLevelShow(dbname string, profile *pb.LogLevelShowProfile) (string, error) {
	var err error = nil
	var orinkey, newkey, hkey, data, module, code, tempstr, rtnstr string
	var modulename, codeindex, tcodestr string
	var splitdata []string
	var dbclient *redis.Client
	var option int32
	var kcursor, fcursor uint64 = 0, 0
	var newdata map[string]string //code为索引,level为值,不同模块的code不重复

	err = CmlDbReconnect(dbname)
	if err != nil {
		rtnstr = "Connect syslog db failed!\n"
		return rtnstr, err
	}

	option = profile.GetOption()

	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]
	if "" != profile.GetModulename() {
		/*显示指定模块的code,level*/
		module = strings.ToLower(profile.GetModulename())
	} else {
		module = "*" /*不指定显示所有模块的*/
	}
	if -1 != profile.GetCode() {
		code = fmt.Sprintf("%d", profile.GetCode())
	} else {
		code = ""
	}
	orinkey = KEY_ALARM_LEVEL_ORIGINAL + module
	newkey = KEY_ALARM_LEVEL_NEW + module
	newdata = make(map[string]string, 200) /*如果不同模块code重复则就需要每个模块都及时完成数据组织*/
	/*scan会根据key自动判断,如果不是通配符则查找指定的,每次找10个,找完为止*/

	for {
		var keys []string
		keys, kcursor, err = dbclient.Scan(kcursor, newkey, 10).Result()
		if err != nil {
			break
		}
		/*从每个key(模块）中读取code-level*/
		fmt.Printf("len:%d\n", len(keys))
		for _, hkey = range keys {
			if KEY_ALARM_UPDATEINFO == hkey {
				continue
			}
			fcursor = 0
			//modulename = strings.TrimLeft(hkey, KEY_ALARM_LEVEL_NEW)
			modulename = strings.TrimPrefix(hkey, KEY_ALARM_LEVEL_NEW)
			for {
				var fields []string
				/*一次最多读100个field,也即一个模块下超过100个告警码的要多次循环*/
				if "" == code { /*显示所有的*/
					fields, fcursor, _ = dbclient.HScan(hkey, fcursor, code, 100).Result()
					codeindex = "0"
					for i, value := range fields {
						/*第一个值为code,第二个为value(level::postion)*/
						/*module|code 为索引*/
						if (i % 2) == 0 {
							codeindex = modulename + "|" + value
						} else {
							newdata[codeindex] = value
						}
					}
					if 0 == fcursor {
						break
					}
				} else { /*显示指定码*/
					data, err = dbclient.HGet(hkey, code).Result()
					if err == nil {
						codeindex = modulename + "|" + code
						newdata[codeindex] = data
					}
					break
				}
			}
		}
		if kcursor == 0 {
			break
		}
	}
	newlen := len(newdata)

	tempstr = ""
	if 0 == option {
		/*只显示有修改的*/
		if 0 == newlen {
			/*没有查询到相关的修改数据*/
			rtnstr = "no data"
		} else {
			/*组织显示数据*/
			rtnstr = fmt.Sprintf("%16s   %8s  %s  %s\n", "module", "code", "level", "position")
			for k, v := range newdata {
				/*将level::position拆分显示*/
				splitdata = strings.Split(v, DATA_SPLIT_STR)
				csplit := strings.Split(k, "|")
				rtnstr += fmt.Sprintf("%16s   %8s  %s  %s\n", csplit[0], csplit[1], splitdata[0], splitdata[1])
			}
		}
	} else {
		/*显示所有的,读取原始数据,然后组织数据时判断是否有修改*/
		kcursor = 0
		for {
			var keys []string
			keys, kcursor, err = dbclient.Scan(kcursor, orinkey, 20).Result()
			if err != nil {
				break
			}
			/*从每个key(模块）中读取code-level*/
			for _, hkey = range keys {
				if KEY_ALARM_UPDATEINFO == hkey {
					continue
				}
				fcursor = 0
				//modulename = strings.TrimLeft(hkey, KEY_ALARM_LEVEL_ORIGINAL)
				modulename = strings.TrimPrefix(hkey, KEY_ALARM_LEVEL_ORIGINAL)
				fmt.Printf("hkey:%s, modulename:%s\n", hkey, modulename)
				for {
					var fields []string
					/*一次最多读100个field,也即一个模块下超过100个告警码的要多次循环*/
					fields, fcursor, err = dbclient.HScan(hkey, fcursor, "*", 100).Result()
					codeindex = "0"
					for i, value := range fields {
						/*第一个值为code,第二个为value(level::postion)*/
						if (i % 2) == 0 {
							codeindex = modulename + "|" + value
							tcodestr = value
						} else {
							/*wait to modify   */
							/*将level::position拆分显示*/
							splitdata = strings.Split(value, DATA_SPLIT_STR)
							if _, ok := newdata[codeindex]; ok {
								/*增加new显示*/
								newsplit := strings.Split(newdata[codeindex], DATA_SPLIT_STR)
								tempstr += fmt.Sprintf("%16s   %8s  %s  %s  newlevel:%s\n", modulename, tcodestr, splitdata[0], splitdata[1], newsplit[0])
							} else {
								tempstr += fmt.Sprintf("%16s   %8s  %s  %s\n", modulename, tcodestr, splitdata[0], splitdata[1])
							}
						}
					}
					if 0 == fcursor {
						break
					}
				}
			}
			if kcursor == 0 {
				break
			}
		}
		if 0 < len(tempstr) {
			rtnstr = fmt.Sprintf("%16s   %8s  %s  %s\n", "module", "code", "level", "position")
			rtnstr += tempstr
		}
	}

	return rtnstr, err
}

/*显示所有配置信息*/
func SyslogShowConfig(dbname string, rtnstr *string) {

	var dbclient *redis.Client
	var err error
	var fieldvalue, outstr string
	var filedexist bool

	err = CmlDbReconnect(dbname)
	if err != nil {
		*rtnstr = "Connect syslog db failed!\n"
		return
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	outstr = "#syslog config, last change:"

	fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_UPDATETIME).Result()
	if "" != fieldvalue {
		fmt.Println(fieldvalue)
		outstr += fieldvalue + " msec\n"
	} else {
		outstr += "0 msec\n"
	}

	filedexist, err = dbclient.HExists(KEY_SYSLOG, FIELD_CONSOLE).Result()
	if filedexist {
		fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_CONSOLE).Result()
		if "" != fieldvalue {
			if level, ok := pb.LOG_LEVEL_value[fieldvalue]; ok {
				if syslog_level_max > level {
					outstr += "  syslog console " + fmt.Sprintf("%d", level) + "\n"
				}

			}
		}
	}

	filedexist, err = dbclient.HExists(KEY_SYSLOG, FIELD_MONITOR).Result()
	if filedexist {
		fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_MONITOR).Result()
		if "" != fieldvalue {
			if level, ok := pb.LOG_LEVEL_value[fieldvalue]; ok {
				if syslog_level_max > level {
					outstr += "  syslog monitor " + fmt.Sprintf("%d", level) + "\n"
				}
			}
		}
	}

	filedexist, err = dbclient.HExists(KEY_SYSLOG, FIELD_TRAP).Result()
	if filedexist {
		fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_TRAP).Result()
		if "" != fieldvalue {
			if level, ok := pb.LOG_LEVEL_value[fieldvalue]; ok {
				if syslog_level_max > level {
					outstr += "  syslog trap " + fmt.Sprintf("%d", level) + "\n"
				}
			}
		}
	}

	filedexist, err = dbclient.HExists(KEY_SYSLOG, FIELD_SERVER).Result()
	if filedexist {
		fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_SERVER).Result()
		if "" != fieldvalue {
			if level, ok := pb.LOG_LEVEL_value[fieldvalue]; ok {
				if syslog_level_max > level {
					outstr += "  syslog loghost " + fmt.Sprintf("%d", level) + "\n"
				}
			}
		}
	}

	filedexist, err = dbclient.HExists(KEY_SYSLOG, FIELD_PERSISTENT).Result()
	if filedexist {
		fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_PERSISTENT).Result()
		if "" != fieldvalue {
			if level, ok := pb.LOG_LEVEL_value[fieldvalue]; ok {
				if syslog_level_max > level {
					outstr += "  syslog persistent " + fmt.Sprintf("%d", level) + "\n"
				}
			}
		}
	}

	filedexist, err = dbclient.HExists(KEY_SYSLOG, FIELD_BUFFERSIZE).Result()
	if filedexist {
		fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_BUFFERSIZE).Result()
		if "" != fieldvalue {
			outstr += "  syslog buffer " + fieldvalue + "\n"
		}
	}

	/*数据库中以m为单位,读取到syslog内存变量时转为byte*/
	filedexist, err = dbclient.HExists(KEY_SYSLOG, FIELD_STORAGESIZE).Result()
	if filedexist {
		fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_STORAGESIZE).Result()
		if "" != fieldvalue {
			outstr += "  syslog memsize " + fieldvalue + "\n"
		}
	}

	filedexist, err = dbclient.HExists(KEY_SYSLOG, FIELD_FILEDURATION).Result()
	if filedexist {
		fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_FILEDURATION).Result()
		if "" != fieldvalue {
			outstr += "  syslog fileduration " + fieldvalue + "\n"
		}
	}

	filedexist, err = dbclient.HExists(KEY_SYSLOG, FIELD_SORCEIP).Result()
	if filedexist {
		fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_SORCEIP).Result()
		if "" != fieldvalue {
			outstr += "  syslog source-address " + fieldvalue + "\n"
		}
	}

	filedexist, err = dbclient.HExists(KEY_SYSLOG, FIELD_SOURCEPORT).Result()
	if filedexist {
		fieldvalue, err = dbclient.HGet(KEY_SYSLOG, FIELD_SOURCEPORT).Result()
		if "" != fieldvalue {
			outstr += "  syslog source-port " + fieldvalue + "\n"
		}
	}

	var serverinfo string = "  syslog host "
	fieldvalue, err = dbclient.HGet(KEY_SYSLOG_SEVER, FIELD_SYSLOG_SERVER1).Result()
	if "" != fieldvalue {
		serverinfo += "server1 "
		/*"vrf:%s,type:%d,ip:%s,port:%d"*/
		paramv := strings.Split(fieldvalue, ",")
		for _, v := range paramv {
			if strings.Contains(v, "vrf:") {
				if "" == strings.TrimLeft(v, "vrf:") {
					serverinfo += "global "
				} else {
					serverinfo += "vrf " + strings.TrimLeft(v, "vrf:") + " "
				}
			}
			if strings.Contains(v, "ip:") {
				serverinfo += "ipv4 " + strings.TrimLeft(v, "ip:") + " "
			}
			if strings.Contains(v, "port:") {
				serverinfo += "port " + strings.TrimLeft(v, "port:")
			}
		}
		outstr += serverinfo + "\n"
	}

	fieldvalue, err = dbclient.HGet(KEY_SYSLOG_SEVER, FIELD_SYSLOG_SERVER2).Result()
	if "" != fieldvalue {
		serverinfo = "  syslog host server2 "
		/*"vrf:%s,type:%d,ip:%s,port:%d"*/
		paramv := strings.Split(fieldvalue, ",")
		for _, v := range paramv {
			if strings.Contains(v, "vrf:") {
				if "" == strings.TrimLeft(v, "vrf:") {
					serverinfo += "global "
				} else {
					serverinfo += "vrf " + strings.TrimLeft(v, "vrf:") + " "
				}
			}
			if strings.Contains(v, "ip:") {
				serverinfo += "ipv4 " + strings.TrimLeft(v, "ip:") + " "
			}
			if strings.Contains(v, "port:") {
				serverinfo += "port " + strings.TrimLeft(v, "port:")
			}
		}
		outstr += serverinfo + "\n"
	}

	*rtnstr = outstr

}

/*显示当前存在的告警信息,state db中*/
func SyslogShowAlarm(dbname string) string {

	var dbclient *redis.Client
	var err error
	var outstr, dkey, hkey, level, timestr, fieldkey, fieldv string
	var splitdata []string
	var kcursor uint64
	var unixsec, unixnano int64
	var occurtime time.Time
	var showfile *os.File

	tm := time.Now()
	timestamp := tm.Unix()
	bakName := fmt.Sprintf("%s%d.bak", "/var/log/", timestamp)
	showfile, err = os.OpenFile(bakName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("Open write file:%s failed, err:%s\n", bakName, err)
		return ""
	}
	defer showfile.Close()

	err = CmlDbReconnect(dbname)
	if err != nil {
		showfile.WriteString("#syslog config, Connect db failed\n")
		return bakName
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	outstr = ""

	//dkey = KEY_ALARM_CURRNET_INFO + "*"
	dkey = "*"
	kcursor = 0
	for {
		var keys []string
		keys, kcursor, err = dbclient.Scan(kcursor, dkey, 20).Result()
		if err != nil {
			break
		}
		/*从每个key(模块）中读取code-level*/
		for _, hkey = range keys {
			/*key:AlarmCur|module|code|position*/
			splitdata = strings.Split(hkey, KEY_DATA_SPLIT_STR)
			if 4 != len(splitdata) {
				continue //该key有错误
			}
			/*目前每个告警key下使用time作为field-key,descrip 作为field-value*/
			/*getall返回为 map[field]value*/
			field, _ := dbclient.HGetAll(hkey).Result()
			for fieldkey, fieldv = range field {
				/*显示告警数据,时间转换为time格式,时间是unix-nano,后9位是纳秒*/
				unixnano = 0
				timestr = fieldkey[0 : len(fieldkey)-9]
				unixsec, err = strconv.ParseInt(timestr, 10, 64)
				if nil != err {
					fmt.Printf("%s\n", err)
					timestr = "unknown"
				} else {
					/*再复用timestr转换纳秒值,如果错误纳秒值默认为0*/
					timestr = fieldkey[(len(fieldkey) - 9):len(fieldkey)]
					fmt.Printf("timestr-nano:%s\n", timestr)
					unixnano, _ = strconv.ParseInt(timestr, 10, 64)
				}
				occurtime = time.Unix(unixsec, unixnano)
				timestr = occurtime.String()
				/*查找level*/
				level = AlarmLevelQuery(splitdata[1], splitdata[2], CFG_DB)
				outstr += fmt.Sprintf("%16s %8s %5s %30s %8s %s\n", splitdata[1], splitdata[2], level, timestr, splitdata[3], fieldv)
			}
		}
		if kcursor == 0 {
			break
		}
	}

	if "" != outstr {
		showfile.WriteString(fmt.Sprintf("%16s %8s %5s %30s %8s %s\n", "module", "code", "level", "time", "position", "description"))
		showfile.WriteString(outstr)
	} else {
		showfile.WriteString("No alarm information\n")
	}
	return bakName
}

/*从db查询告警码对应的告警严重级别,返回严重级别字符串,如果没查到,返回“-”*/
func AlarmLevelQuery(module, code, dbname string) string {

	var dkey, levelstr string
	var err error
	var findcode bool = false
	var splitdata []string
	var fcursor uint64 = 0
	var dbclient *redis.Client

	err = CmlDbReconnect(dbname)
	if err != nil {
		return "-"
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[dbname]]

	/*key为 Alarm|module,读取code-level::position, 域*/
	/*先查有无修改的,如果没有再查原始表*/
	dkey = KEY_ALARM_LEVEL_NEW + module
	fcursor = 0
	levelstr = ""
	for {
		var fields []string
		/*一次最多读100个field,也即一个模块下超过100个告警码的要多次循环*/
		fields, fcursor, err = dbclient.HScan(dkey, fcursor, "*", 100).Result()
		if err != nil {
			break
		}
		for i, value := range fields {
			/*第一个值为code,第二个为value(level::postion)*/
			if (i % 2) == 0 {
				if value == code {
					findcode = true
				}
			} else {
				if true == findcode {
					/*将level::position拆分显示,返回对应的level*/
					splitdata = strings.Split(value, DATA_SPLIT_STR)
					levelstr = splitdata[0]
					break
				}
			}
		}
		if (0 == fcursor) || (true == findcode) {
			break
		}
	}
	if "" != levelstr {
		return levelstr
	}

	/*再查询原始表*/
	dkey = KEY_ALARM_LEVEL_ORIGINAL + module
	for {
		var fields []string
		/*一次最多读100个field,也即一个模块下超过100个告警码的要多次循环*/
		fields, fcursor, err = dbclient.HScan(dkey, fcursor, "*", 100).Result()
		for i, value := range fields {
			/*第一个值为code,第二个为value(level::postion)*/
			if (i % 2) == 0 {
				if value == code {
					findcode = true
				}
			} else {
				if true == findcode {
					/*将level::position拆分显示,返回对应的level*/
					splitdata = strings.Split(value, DATA_SPLIT_STR)
					levelstr = splitdata[0]
					break
				}
			}
		}
		if (0 == fcursor) || (true == findcode) {
			break
		}
	}
	/*没查到返回"-"*/
	if "" == levelstr {
		levelstr = "-"
	}
	return levelstr
}

/*if not exist or invalid value,return false & max-id(invalid)*/
/*check the level(one configure) whether exist*/
var LOG_LEVEL_trasn = map[int32]pb.LOG_LEVEL{
	0: 0,
	1: 1,
	2: 2,
	3: 3,
	4: 4,
	5: 5,
	6: 6,
	7: 7,
}

func SyslogDbLevelExist(dbclient *redis.Client, key, field string) (pb.LOG_LEVEL, bool) {
	var fieldvalue string

	filedexist, _ := dbclient.HExists(key, field).Result()
	if !filedexist {
		/*return default value*/
		return 0, false
	} else {
		fieldvalue, _ = dbclient.HGet(key, field).Result()
		if "" != fieldvalue {
			if level, ok := pb.LOG_LEVEL_value[fieldvalue]; ok {
				if syslog_level_max > level {
					return LOG_LEVEL_trasn[level], true
				}
			}
		}
	}
	return 0, false
}

/*set level config*/
/*before call ,level must be valid*/
func SyslogDbLevelSet(dbclient *redis.Client, key, field string, level pb.LOG_LEVEL) error {

	err := dbclient.HSet(key, field, fmt.Sprintf("%v", level)).Err()
	return err
}

/*set console config*/
func SyslogDbLevelDel(dbclient *redis.Client, key, field string) error {

	err := dbclient.HDel(key, field).Err()
	return err
}

func syslogprintdbcfg() {

	var cursor uint64
	var n int
	var err error
	var keys []string

	dbclient := cmldbinfo.cmldbclient[CmlDbNo["CONFIG_DB"]]
	for {
		keys, cursor, err = dbclient.Scan(cursor, "syslog", 10).Result()
		if err != nil {
			panic(err)
		}
		n += len(keys)
		fmt.Println("keys:\n", keys)

		if cursor == 0 {
			break
		}
	}

	fmt.Printf("found %d keys\n", n)

}

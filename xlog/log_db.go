// log_db.go
// 创建者 陈维 2019.11.20
// syslog server从数据库读取配置的处理,配置来自于CML进程写入数据库,然后通过publish通知syslog
// server读取. syslog server重启后也通过该处理实现配置的恢复(在整个系统重启恢复后,需
// 要有pub发布,以通知业务读取.通过该方式可减少重启时序的控制要求)

package xlog

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type T_LOG_DBINFO struct {
	clientcfgdb     *redis.Client
	logpubsub       *redis.PubSub
	stateinfodb     *redis.Client /*STATE_DB,6号,存放告警信息*/
	cfgdbstate      bool
	cfgsubstate     bool
	statedbstate    bool
	cfgdbupdatetime string
}

var logdbinfo T_LOG_DBINFO
var pubchannel string = "syslog.config"

var LogLevelMap = map[string]LOG_LEVEL{
	"LOG_LEVEL_EMERGENCY": LOG_LEVEL_EMERGENCY,
	"LOG_LEVEL_ALERT":     LOG_LEVEL_ALERT,
	"LOG_LEVEL_CRITICAL":  LOG_LEVEL_CRITICAL,
	"LOG_LEVEL_ERROR":     LOG_LEVEL_ERROR,
	"LOG_LEVEL_WARNING":   LOG_LEVEL_WARNING,
	"LOG_LEVEL_NOTICE":    LOG_LEVEL_NOTICE,
	"LOG_LEVEL_INFO":      LOG_LEVEL_INFO,
	"LOG_LEVEL_DEBUG":     LOG_LEVEL_DEBUG,
	"0":                   LOG_LEVEL_EMERGENCY,
	"1":                   LOG_LEVEL_ALERT,
	"2":                   LOG_LEVEL_CRITICAL,
	"3":                   LOG_LEVEL_ERROR,
	"4":                   LOG_LEVEL_WARNING,
	"5":                   LOG_LEVEL_NOTICE,
	"6":                   LOG_LEVEL_INFO,
	"7":                   LOG_LEVEL_DEBUG,
}

const (
	KEY_ALARM_LEVEL_ORIGINAL string = "Alarm|"
	KEY_ALARM_LEVEL_NEW             = "AlarmNew|"
	KEY_ALARM_UPDATEINFO            = "Alarm|updateinfo"
	FIELD_ALARM_UPDATETIME          = "updatetime"
	DATA_SPLIT_STR                  = "::" /*用于分割level与position*/
	DATA_UPDATEINFO_STR             = "updateinfo"
	ALARM_OCCUR_STR                 = "[alarm:occur]"
	ALARM_CLEAR_STR                 = "[alarm:clear]"
)

func InitCmlDbInfo() {

	var err error

	logdbinfo.cfgdbupdatetime = "0"

	logdbinfo.clientcfgdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       4,  // use config DB,
	})
	_, err = logdbinfo.clientcfgdb.Ping().Result()
	if err == nil {
		logdbinfo.cfgdbstate = true
		logdbinfo.logpubsub = logdbinfo.clientcfgdb.Subscribe("syslog.config")
		/*在cml没创建syslog相关db及channel时,可能为空*/
		if nil == logdbinfo.logpubsub {
			fmt.Println("subscibe syslog.config failed!")
		} else {
			logdbinfo.cfgsubstate = true
		}
	}

	logdbinfo.stateinfodb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       6,  // use STATE DB,
	})
	_, err = logdbinfo.stateinfodb.Ping().Result()
	if err == nil {
		logdbinfo.statedbstate = true
	}
	return
}

func LogCfgDbReconnect() error {
	var err error = nil

	if !logdbinfo.cfgdbstate {
		logdbinfo.clientcfgdb = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       4,  // use config DB,
		})
		_, err = logdbinfo.clientcfgdb.Ping().Result()
		if err == nil {
			logdbinfo.cfgdbstate = true
			logdbinfo.logpubsub = logdbinfo.clientcfgdb.Subscribe("syslog.config")
			if nil == logdbinfo.logpubsub {
				fmt.Println("subscibe syslog.config failed!")
			} else {
				logdbinfo.cfgsubstate = true
			}
		}
	}
	/*数据库可能一开始就连接成功,但订阅的频道可能未创建时为空*/
	if !logdbinfo.cfgsubstate {
		logdbinfo.logpubsub = logdbinfo.clientcfgdb.Subscribe("syslog.config")
		if nil == logdbinfo.logpubsub {
			fmt.Println("subscibe syslog.config failed!")
		} else {
			logdbinfo.cfgsubstate = true
		}
	}

	if !logdbinfo.statedbstate {
		logdbinfo.stateinfodb = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       6,  // use STATE DB,
		})
		_, err = logdbinfo.stateinfodb.Ping().Result()
		if err == nil {
			logdbinfo.statedbstate = true
		}
	}

	return err
}

func LogRedisGet() error {

	var i, filedur int
	var storagesize int64 = 0
	var err error = nil
	var fieldvalue string
	var filedexist bool

	for i = 0; i < 3; i++ {
		err = LogCfgDbReconnect()
		if nil != err {
			return err
		} else {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}

	fieldvalue, err = logdbinfo.clientcfgdb.HGet("syslog", "updatetime").Result()
	if "" != fieldvalue {
		if logdbinfo.cfgdbupdatetime == fieldvalue {
			/*data had flushed */
			fmt.Println("time equal,return", logdbinfo.cfgdbupdatetime, fieldvalue)
			return err
		} else {
			logdbinfo.cfgdbupdatetime = fieldvalue
		}
	} else {
		return err
	}

	filedexist, _ = logdbinfo.clientcfgdb.HExists("syslog", "console").Result()
	if filedexist {
		fieldvalue = logdbinfo.clientcfgdb.HGet("syslog", "console").Val()
		if "" != fieldvalue {
			if level, ok := LogLevelMap[fieldvalue]; ok {
				if LOG_LEVEL_DEBUG >= level {
					SetGlobalConsoleLogLevel(level)
				}
			}
		}
	} else {
		/*已执行no命令或没有配置,恢复默认值*/
		SetGlobalConsoleLogLevel(LOG_LEVEL_INFO)
	}

	filedexist, _ = logdbinfo.clientcfgdb.HExists("syslog", "monitor").Result()
	if filedexist {
		fieldvalue, err = logdbinfo.clientcfgdb.HGet("syslog", "monitor").Result()
		if "" != fieldvalue {
			if level, ok := LogLevelMap[fieldvalue]; ok {
				if LOG_LEVEL_DEBUG >= level {
					logService.monitorLevel = level
				}
			}
		}
	} else {
		/*已执行no命令或没有配置,恢复默认值*/
		logService.monitorLevel = LOG_LEVEL_INFO
	}

	filedexist, _ = logdbinfo.clientcfgdb.HExists("syslog", "trap").Result()
	if filedexist {
		fieldvalue = logdbinfo.clientcfgdb.HGet("syslog", "trap").Val()
		if "" != fieldvalue {
			if level, ok := LogLevelMap[fieldvalue]; ok {
				if LOG_LEVEL_DEBUG >= level {
					logService.trapLevel = level //trap
				}
			}
		}
	} else {
		/*恢复默认值,应该是不发送,但由于0为有效值,需要重新定义*/
		logService.trapLevel = 0
	}

	filedexist, _ = logdbinfo.clientcfgdb.HExists("syslog", "server").Result()
	if filedexist {
		fieldvalue = logdbinfo.clientcfgdb.HGet("syslog", "server").Val()
		if "" != fieldvalue {
			if level, ok := LogLevelMap[fieldvalue]; ok {
				if LOG_LEVEL_DEBUG >= level {
					logService.serverLevel = level //server level
				}
			}
		}
	} else {
		/*恢复默认值,应该是不发送,但由于0为有效值,需要重新定义*/
		logService.serverLevel = 0
	}

	filedexist, _ = logdbinfo.clientcfgdb.HExists("syslog", "persistent").Result()
	if filedexist {
		fieldvalue = logdbinfo.clientcfgdb.HGet("syslog", "persistent").Val()
		if "" != fieldvalue {
			if level, ok := LogLevelMap[fieldvalue]; ok {
				if LOG_LEVEL_DEBUG >= level {
					SetGlobalPersitentLogLevel(level)
					//logService.persistentLevel = level
				}
			}
		}
	} else {
		/*恢复默认值,需要重新定义*/
		SetGlobalPersitentLogLevel(LOG_LEVEL_ERROR)
	}

	filedexist, _ = logdbinfo.clientcfgdb.HExists("syslog", "memsize").Result()
	if filedexist {
		fieldvalue = logdbinfo.clientcfgdb.HGet("syslog", "memsize").Val()
		if "" != fieldvalue {
			storagesize, _ = strconv.ParseInt(fieldvalue, 10, 64)
			storagesize = storagesize * (1 << 20) //trans bytes
			if LOG_MAX_STORAGE_SIZE > storagesize {
				/*only change storagesize, don't modify period*/
				SetLogPersitentParam(storagesize, -1)
				//logService.maxStorageSize = storagesize
			}
		}
	} else {
		SetLogPersitentParam(LOG_MAX_STORAGE_SIZE, -1)
	}

	filedexist, _ = logdbinfo.clientcfgdb.HExists("syslog", "fileduration").Result()
	if filedexist {
		fieldvalue = logdbinfo.clientcfgdb.HGet("syslog", "fileduration").Val()
		if "" != fieldvalue {
			filedur, _ = strconv.Atoi(fieldvalue)
			if (2 * LOG_MAX_STORAGE_DURATION) > filedur {
				SetLogPersitentParam(-1, filedur)
			}
		}
	} else {
		SetLogPersitentParam(-1, LOG_MAX_STORAGE_DURATION)
	}

	fieldvalue = logdbinfo.clientcfgdb.HGet("logserver", "server1").Val()
	var numberstr string
	if "" != fieldvalue {
		/*"vrf:%s,type:%d,ip:%s,port:%d"*/
		paramv := strings.Split(fieldvalue, ",")
		for _, v := range paramv {
			if strings.Contains(v, "vrf:") {
				logService.outSyslogServer[0].vrfname = strings.TrimLeft(v, "vrf:")
			}
			if strings.Contains(v, "ip:") {
				numberstr = strings.TrimLeft(v, "ip:")
				logService.outSyslogServer[0].servAddr.IP = net.ParseIP(numberstr)
			}
			if strings.Contains(v, "port:") {
				numberstr = strings.TrimLeft(v, "port:")
				logService.outSyslogServer[0].servAddr.Port, _ = strconv.Atoi(numberstr)
			}
		}
		if 0 == logService.outSyslogServer[0].servAddr.Port {
			logService.outSyslogServer[0].servAddr.Port = 514 //set default
		}
	} else {
		if 0 != logService.outSyslogServer[0].servAddr.Port {
			/*del config*/
			/*udp connect need close*/
			logService.outSyslogServer[0].vrfname = ""
			logService.outSyslogServer[0].servAddr.IP = nil
			logService.outSyslogServer[0].servAddr.Port = 0
		}
	}

	fieldvalue = logdbinfo.clientcfgdb.HGet("logserver", "server2").Val()
	if "" != fieldvalue {
		/*"vrf:%s,type:%d,ip:%s,port:%d"*/
		paramv := strings.Split(fieldvalue, ",")
		for _, v := range paramv {
			if strings.Contains(v, "vrf:") {
				logService.outSyslogServer[1].vrfname = strings.TrimLeft(v, "vrf:")
			}
			if strings.Contains(v, "ip:") {
				numberstr = strings.TrimLeft(v, "ip:")
				logService.outSyslogServer[1].servAddr.IP = net.ParseIP(numberstr)
			}
			if strings.Contains(v, "port:") {
				numberstr = strings.TrimLeft(v, "port:")
				logService.outSyslogServer[1].servAddr.Port, _ = strconv.Atoi(numberstr)
			}
		}
		if 0 == logService.outSyslogServer[1].servAddr.Port {
			logService.outSyslogServer[1].servAddr.Port = 514 //set default
		}
	} else {
		if 0 != logService.outSyslogServer[1].servAddr.Port {
			/*del config*/
			/*udp connect need close*/
			logService.outSyslogServer[1].vrfname = ""
			logService.outSyslogServer[1].servAddr.IP = nil
			logService.outSyslogServer[1].servAddr.Port = 0
		}
	}
	/*for debug, need del*/
	PrintLogServiceStats()

	return nil
}

/*告警信息表,key:AlarmCur|module|code|position*/
/*mo,co,po 三个组合是唯一的*/
/*field: time(fieldname)-descritption*/
/*level从内存map表中获取,只有show,或snmp get时需要*/
/*status默认为occur,clear后就从表中删除*/
func AlarmStateUpdate(msgStru *InnerLogMsg) {
	var dkey, field string
	var err error
	var tnano int64

	dkey = "AlarmCur|" /*KEY_ALARM_CURRNET_INFO*/ + msgStru.ModuleName + "|" + fmt.Sprintf("%d", msgStru.ErrCode) + "|" + msgStru.Position

	/*状态是通过descript带的,格式为 [alarm:occur] */
	if strings.HasPrefix(msgStru.LogDesc, ALARM_OCCUR_STR) {
		/*告警发生*/
		tnano = msgStru.OccureTime.UnixNano()
		field = strconv.FormatInt(tnano, 10)
		err = logdbinfo.stateinfodb.HSet(dkey, field, msgStru.LogDesc).Err()
		if nil != err {
			fmt.Println("%s\n", err)
		}
	} else if strings.HasPrefix(msgStru.LogDesc, ALARM_CLEAR_STR) {
		/*告警消除*/
		err = logdbinfo.stateinfodb.Del(dkey).Err()
		if nil != err {
			fmt.Println("%s\n", err)
		}
	} else {
		/*错误告警信息*/
		fmt.Println("Error info,descrip:%s\n", msgStru.LogDesc)
	}
}

/*将告警码,告警级别从db读取存入内存map表中.告警处理中转syslog日志时从map表中查level更新,*/
/*另在告警信息的show, snmp get,trap时也需要查询获取level*/
func AlarmLoadCodeInfo() {
	var dkey, hkey, modulename, codeindex string
	var err error
	var splitdata []string
	var kcursor, fcursor uint64 = 0, 0

	dkey = KEY_ALARM_LEVEL_ORIGINAL + "*"

	/*先读取定义值*/
	for {
		var keys []string
		keys, kcursor, err = logdbinfo.clientcfgdb.Scan(kcursor, dkey, 20).Result()
		if err != nil {
			break
		}
		/*从每个key(模块）中读取code-level*/
		for _, hkey = range keys {
			if KEY_ALARM_UPDATEINFO == hkey {
				continue
			}
			fcursor = 0
			modulename = strings.TrimPrefix(hkey, KEY_ALARM_LEVEL_ORIGINAL)
			for {
				var fields []string
				/*一次最多读100个field,也即一个模块下超过100个告警码的要多次循环*/
				fields, fcursor, err = logdbinfo.clientcfgdb.HScan(hkey, fcursor, "*", 100).Result()
				codeindex = "0"
				for i, value := range fields {
					/*第一个值为code,第二个为value(level::postion)*/
					if (i % 2) == 0 {
						codeindex = modulename + "|" + value
					} else {
						/*将level::position拆分显示*/
						splitdata = strings.Split(value, DATA_SPLIT_STR)
						/*如果是有效的level,即表中有对应项*/
						if _, ok := LogLevelMap[splitdata[0]]; ok {
							logService.alarmCodeLevel[codeindex] = LogLevelMap[splitdata[0]]
						} else {
							fmt.Println("Error level,%s\n", splitdata[0])
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

	/*再查询有无命令修改,如有更新内存map中的值*/
	dkey = KEY_ALARM_LEVEL_NEW + "*"
	kcursor = 0
	for {
		var keys []string
		keys, kcursor, err = logdbinfo.clientcfgdb.Scan(kcursor, dkey, 20).Result()
		if err != nil {
			break
		}
		/*从每个key(模块）中读取code-level*/
		for _, hkey = range keys {
			if KEY_ALARM_UPDATEINFO == hkey {
				continue
			}
			fcursor = 0
			modulename = strings.TrimPrefix(hkey, KEY_ALARM_LEVEL_NEW)
			for {
				var fields []string
				/*一次最多读100个field,也即一个模块下超过100个告警码的要多次循环*/
				fields, fcursor, err = logdbinfo.clientcfgdb.HScan(hkey, fcursor, "*", 100).Result()
				codeindex = "0"
				for i, value := range fields {
					/*第一个值为code,第二个为value(level::postion)*/
					if (i % 2) == 0 {
						codeindex = modulename + "|" + value
					} else {
						/*将level::position拆分显示*/
						splitdata = strings.Split(value, DATA_SPLIT_STR)
						/*如果是有效的level,即表中有对应项*/
						if _, ok := LogLevelMap[splitdata[0]]; ok {
							logService.alarmCodeLevel[codeindex] = LogLevelMap[splitdata[0]]
						} else {
							fmt.Println("Error level,%s\n", splitdata[0])
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
	return
}

/*收到db发布的订阅信息,读取db更新内存中map表的level值*/
func AlarmUpdateCodeInfo() {
	AlarmLoadCodeInfo()
}

func LogCfgDbListen() error {
	var cfgchan string = "syslog.config"
	var showchan string = "syslog.show"

	for {
		time.Sleep(5000 * time.Millisecond)
		if logdbinfo.cfgsubstate {
			msg, err := logdbinfo.logpubsub.ReceiveMessage()
			if err != nil {
				fmt.Print("Error in ReceiveTimeout()", err)
				continue
			}
			/*没有收到订阅消息,继续等待*/
			if nil == msg {
				fmt.Print("Receive nil, continue ")
				continue
			}
			if msg.Channel == cfgchan {
				/*update config from db*/
				fmt.Println("Receive pub, enter get db", msg.Payload)
				LogRedisGet()
			} else if msg.Channel == showchan {
				fmt.Println("Receive pub, enter get show")
				/*update buffer or statictics dynamic data, for show*/
				LogBufferShow(msg.Payload)
			} else {
				fmt.Println("Receive publish,channel:, payload:", msg.Channel, msg.Payload)
			}
		} else {
			/*延迟一段时间再尝试订阅*/
			time.Sleep(5 * time.Second)
			LogCfgDbReconnect()
		}
	}
}

func LogBufferShow(showinfo string) {

	fmt.Println("Show command:%s", showinfo)

}

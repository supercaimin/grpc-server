// syslogconfig.go
// syslog cml process

package config

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

/*数据库订阅*/
const (
	CFG_DB    string = "CONFIG_DB"
	APP_DB           = "APP_DB"
	COUTER_DB        = "COUNTERS_DB"
	ASIC_DB          = "ASIC_DB"
	STATE_DB         = "STATE_DB"
)

/*0-7 db*/
type T_CML_DBINFO struct {
	cmldbclient [8]*redis.Client
	cmldbstate  [8]bool
}

var CmlDbNo = map[string]int{
	"APP_DB":          0,
	"ASIC_DB":         1,
	"COUNTERS_DB":     2,
	"LOGLEVEL_DB":     3,
	"CONFIG_DB":       4,
	"PFC_WD_DB":       5,
	"STATE_DB":        6,
	"SNMP_OVERLAY_DB": 7,
}

var CmlDbName = map[int]string{
	0: "APP_DB",
	1: "ASIC_DB",
	2: "COUNTERS_DB",
	3: "LOGLEVEL_DB",
	4: "CONFIG_DB",
	5: "PFC_WD_DB",
	6: "STATE_DB",
	7: "SNMP_OVERLAY_DB",
}

/*in CML SERVER*/
/*every db has multi channel subscribe*/
/*every channel has multi app to subscribe*/
/*db->[]channel->[](app,proc)*/

type T_CML_SUB_CHANINFO struct {
	pubsub *redis.PubSub
	//channel string, map string index
	modulenum int /*频道订阅者数量*/
	modulefun map[string]*CmlDbSubListen
}

var cmldbinfo T_CML_DBINFO
var aclRedisClient *redis.Client

/*subscribe app db channel and proc func info*/
/*the same channel maybe multi subscribers*/

/*subscribe counter db channel and proc func info*/
/*int present db,0-7*/
var cmlsubchaninfo [8]map[string]T_CML_SUB_CHANINFO

/*channel callback interface*/
type CmlDbSubListen interface {
	/**/
	SubDbProcess(channel, payload string) bool
}

func InitCmlDbInfo() {

	var err error

	cmldbinfo.cmldbclient[CmlDbNo[CFG_DB]] = redis.NewClient(&redis.Options{
		Addr:     "192.168.3.103:6379",
		Password: "",              // no password set
		DB:       CmlDbNo[CFG_DB], // use config DB,
		// PoolSize: 20, max number of socket connection,defaul is 10
	})
	_, err = cmldbinfo.cmldbclient[CmlDbNo[CFG_DB]].Ping().Result()
	if err == nil {
		cmldbinfo.cmldbstate[CmlDbNo[CFG_DB]] = true

	}
	/*acl配置后续要切换为统一的db句柄   --*/
	aclRedisClient = cmldbinfo.cmldbclient[CmlDbNo[CFG_DB]]

	cmldbinfo.cmldbclient[CmlDbNo[APP_DB]] = redis.NewClient(&redis.Options{
		Addr:     "192.168.3.103:6379",
		Password: "",              // no password set
		DB:       CmlDbNo[APP_DB], // app DB,
	})
	_, err = cmldbinfo.cmldbclient[CmlDbNo[APP_DB]].Ping().Result()
	if err == nil {
		cmldbinfo.cmldbstate[CmlDbNo[APP_DB]] = true
	}

	cmldbinfo.cmldbclient[CmlDbNo[COUTER_DB]] = redis.NewClient(&redis.Options{
		Addr:     "192.168.3.103:6379",
		Password: "",                 // no password set
		DB:       CmlDbNo[COUTER_DB], // conter DB,
	})
	_, err = cmldbinfo.cmldbclient[CmlDbNo[COUTER_DB]].Ping().Result()
	if err == nil {
		cmldbinfo.cmldbstate[CmlDbNo[COUTER_DB]] = true
	}

}

func CmlDbReconnect(dbname string) error {
	var err error = nil

	if !cmldbinfo.cmldbstate[CmlDbNo[dbname]] {
		cmldbinfo.cmldbclient[CmlDbNo[dbname]] = redis.NewClient(&redis.Options{
			Addr:     "192.168.3.103:6379",
			Password: "", // no password set
			DB:       CmlDbNo[dbname],
		})
		_, err = cmldbinfo.cmldbclient[CmlDbNo[dbname]].Ping().Result()
		if err == nil {
			cmldbinfo.cmldbstate[CmlDbNo[dbname]] = true
		}
	}

	return err
}

/*init subscribe info*/
func InitCmlDbSubscribeInfo() {

	/*max 8db,first app 2(app & counter); subscribe channel */
	for i, _ := range cmlsubchaninfo {
		/*default 8 channel per db, max 32 channel*/
		cmlsubchaninfo[i] = make(map[string]T_CML_SUB_CHANINFO)
	}

	return
}

func CmlDbCfgPubstring(dbname, channel, message string) error {
	var err error = nil

	/*cmldbinfo.clientcfgdb.Publish(channel, message) return type is *redis.IntCmd*/
	if !cmldbinfo.cmldbstate[CmlDbNo[dbname]] {
		err = CmlDbReconnect(dbname)
	}

	if nil == err {
		err = cmldbinfo.cmldbclient[CmlDbNo[dbname]].Publish(channel, message).Err()
	}

	return err
}

/*pub interface * */
func CmlDbCfgPublish(dbname, channel string, message interface{}) error {

	marshalstr, err := json.Marshal(message)

	if nil == err {
		messagestr := string(marshalstr)
		err = CmlDbCfgPubstring(dbname, channel, messagestr)
	}

	return err
}

/*subscribe , app register process function*/
/*一个数据库可能存在多个频道订阅,一个频道存在多个模块订阅*/
/*如果一个模块订阅一个已经存在的频道,则不用再向数据库发订阅,只需更新模块处理函数链*/
/*在频道全部都关闭订阅后再关闭该数据库的订阅句柄*/
func CmlDbSubchan(dbname, channel, modulename string, procfun CmlDbSubListen) error {
	//var err error
	//var subch *redis.PubSub
	var chanexist bool = false
	var temp T_CML_SUB_CHANINFO

	/*check channel had subscribe?*/
	if ch, ok := cmlsubchaninfo[CmlDbNo[dbname]][channel]; ok {
		chanexist = true
		if v, ok := ch.modulefun[modulename]; ok {
			/*exist, replace*/
			if v != &procfun {
				ch.modulefun[modulename] = &procfun
			}
		} else {
			/*not exist,new module*/
			ch.modulenum++
			ch.modulefun[modulename] = &procfun
		}
	}

	if !chanexist {
		temp = cmlsubchaninfo[CmlDbNo[dbname]][channel]
		temp.modulenum = 1
		temp.modulefun[modulename] = &procfun
		temp.pubsub = cmldbinfo.cmldbclient[CmlDbNo[dbname]].Subscribe(channel)
	}

	return nil
}

func CmlDbUnSubchan(dbname, channel, modulename string) {
	//var err error
	//var subch *redis.PubSub
	var chanexist, ok bool = false, false
	var ch T_CML_SUB_CHANINFO

	/*check channel had subscribe?*/
	if ch, ok = cmlsubchaninfo[CmlDbNo[dbname]][channel]; ok {
		chanexist = true
		if _, ok = ch.modulefun[modulename]; ok {
			/*module exist, delete*/
			ch.modulenum--

			delete(ch.modulefun, modulename)
		}
	}

	/*再检查该channel是否还有订阅者,如果没有则取消该频道订阅*/
	if chanexist {
		ch = cmlsubchaninfo[CmlDbNo[dbname]][channel]
		//for k, v := range ch.modulefun {
		//moduleexist = true
		//	break
		//}

		/*if !moduleexist {*/
		if 0 == ch.modulenum {
			/*No module subscribe,*/
			ch.pubsub.Unsubscribe(channel)
			delete(cmlsubchaninfo[CmlDbNo[dbname]], channel)
		}
	}

	return
}

/*sub app service callback*/
func CmlDbSubCallback(dbname, channel, payload string) {

	if module, ok := cmlsubchaninfo[CmlDbNo[dbname]][channel]; ok {

		for k, v := range module.modulefun {
			fmt.Printf("db:%s,channle:%s, module:%s callback", dbname, channel, k)
			go (*v).SubDbProcess(channel, payload)
		}
	}

	return
}

func CmlDbPubSubListen(dbname string) error {

	var pubsub *redis.PubSub = nil

	for _, v := range cmlsubchaninfo[CmlDbNo[dbname]] {
		pubsub = v.pubsub
		break
	}
	if nil == pubsub {
		return nil
	}
	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			fmt.Print("Error in ReceiveTimeout()", err)
			continue
		}

		// Process the message
		CmlDbSubCallback(dbname, msg.Channel, msg.Payload)
	}
}

/*debug function*/
func cmldbinfoprint() {

	fmt.Printf("cml db connect state:\n")

	for i := range cmldbinfo.cmldbstate {
		if !cmldbinfo.cmldbstate[i] {
			continue
		}
		fmt.Printf("  %s connect ok!", CmlDbName[i])

		for k, v := range cmlsubchaninfo[i] {
			fmt.Printf("    subscribe channel: %v ", k)

			for module, _ := range v.modulefun {
				fmt.Printf("      subscribe module: %v ", module)
			}
		}
	}

}

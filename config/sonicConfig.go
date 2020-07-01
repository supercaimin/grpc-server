// sonicConfig.go
// SONIC原生命令服务端的处理函数具体实现
// 处理来源可以是klish,也有可能是web-gui,netconf等

package config

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	pb "hstcmler/cml"

	"github.com/go-redis/redis"
)

const (
	CMD_END_RTNSTR  string = "cmdend-"    /*业务执行完命令后显示输出该字符串*/
	CMD_SPLIT_LABEL        = ":0HSTHST0:" /*视图多个命令切分字符串*/
	CMD_END_STR            = "end"
)

//sonic的命令来源信息以及业务对处理的返回信息等
//因为对SONIC原生命令，CML处理后还需要业务继续处理。因对原生的命令处理，配置的有效性检查需要
//到业务处处理才行(或者就需要业务提供配置检查,或者把klish中konfd的纯文本检查移植过来，选择？
type SonicConfig struct {
}

func NewSonicConfig() *SonicConfig {
	return new(SonicConfig)
}

//sonic 原生命令的处理分发函数
/*通过rpush方式放入链数据库,业务从列表中取出逐条执行*/
/*命令输入所有终端放在一个链表key中以避免业务的分队列调度,但执行时要分vty终*/
/*端,避免乱序,因为有进模式的命令以及"end"回特权模式的命令*/
/*命令执行输出按vtyno分开,各通道只关注自己的命令*/
/*每条命令执行完以后都再显示的输出"cmdend",以便cml能知道执行到第几条命令了*/
/*因为在命令执行时需要先执行视图中的命令,然后才是当前的命令*/
/*每次都将视图的命令执行一次,可避免进视图的命令要同步给业务,解耦*/
/*当前命令入队后需要再入队一条"end"命令,回到特权模式*/
/*第一节点只考虑统一从CML到FRR,暂不考虑多vty终端并行发送到FRR的情况 chenwei 20191231*/
func (self *SonicConfig) ExecSonicCfgProfile(ctx context.Context, profile *pb.SONICCmdInputProfile) (*pb.SONICCmdRtnProfile, error) {

	var cmdRtnInfo = &pb.SONICCmdRtnProfile{CmdId: 1, CmdRtncode: 0}
	var err error
	var cmdnum, rtncode int = 0, 0
	var dbclient *redis.Client
	var datavalue, rtnstr /*, vtystr*/ string
	var cmdstr []string
	//var re *regexp.

	cmdRtnInfo.CmdRtncode = 0
	cmdRtnInfo.CmdId = profile.GetCmdId()
	cmdRtnInfo.UpdatedAt = profile.GetUpdatedAt()
	cmdRtnInfo.KlishVtyNo = profile.GetKlishVtyNo()
	fmt.Printf("Rcv gRpc Call id:%d,cmdline:%s\n", profile.CmdId, profile.CmdinputLine)

	//vtystr = fmt.Sprintf("%d", profile.GetKlishVtyNo())

	err = CmlDbReconnect(CFG_DB)
	if nil != err {
		cmdRtnInfo.CmdrtnStr = "db connect failed!\n"
		return cmdRtnInfo, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[CFG_DB]]
	/*先将vtyno对应的输出队列清空,避免之前超时退出的残留返回影响当前的判断*/
	datavalue = KEY_FRR_CMD_OUTPUT // + vtystr
	rtnlens, _ := dbclient.LLen(datavalue).Result()
	dbclient.LTrim(datavalue, rtnlens+1, rtnlens)

	/*匹配表达式串,前4个为数字,后4个为16进制(数字加a-f组合)*/
	re, _ := regexp.Compile("^[0-9]{4}[0-9A-Fa-f]{4}")
	/*将视图命令字符串压入列表*/
	if "" != profile.GetCmdviewList() {
		cmdstr = strings.Split(profile.GetCmdviewList(), CMD_SPLIT_LABEL)
		for _, viewstr := range cmdstr {
			if "" != viewstr {
				/*将vtyno也带入,用于命令输出的返回*/
				datavalue = /*vtystr + " " + */ viewstr
				/*视图命令有的可能带了huastart的cmdid,有的不带;带的前4个字节应全是数字,后*/
				/*4字节可能是数字或16进制的a-f的组合;带id的需要去掉id*/
				if re.MatchString(datavalue) {
					/*去掉前面8个16进制字符的commandid*/
					datavalue = datavalue[8:len(datavalue)]
				}
				err = dbclient.RPush(KEY_FRR_CMD_INPUT, datavalue).Err()
				if nil != err {
					fmt.Println("RPush fail,err:\n", err)
				}
				fmt.Printf("RPush key:%s value:%s\n", KEY_FRR_CMD_INPUT, datavalue)
				cmdnum++
			}
		}
	}
	/*将当前命令字符串压入列表*/
	if "" != profile.GetCmdinputLine() {
		datavalue = /*vtystr + " " +*/ profile.GetCmdinputLine()
		err = dbclient.RPush(KEY_FRR_CMD_INPUT, datavalue).Err()
		if nil != err {
			fmt.Println("RPush fail,err:\n", err)
		}
		fmt.Printf("RPush key:%s value:%s\n", KEY_FRR_CMD_INPUT, datavalue)
		cmdnum++
	}
	fmt.Printf("cmdnum:%d\n", cmdnum)
	if 0 < cmdnum {
		/*压入"end"命令*/
		datavalue = /* vtystr + " " +*/ CMD_END_STR
		err = dbclient.RPush(KEY_FRR_CMD_INPUT, datavalue).Err()
		if nil != err {
			fmt.Println("RPush fail,err:\n", err)
		}
		fmt.Printf("RPush key:%s value:%s\n", KEY_FRR_CMD_INPUT, datavalue)
		cmdnum++
		/*守护输出,直到所有命令都执行完*/
		datavalue = KEY_FRR_CMD_OUTPUT //+ vtystr
		for 0 < cmdnum {
			/*设置5秒,因grpc设置了10秒超时*/
			cmdstr, err = dbclient.BRPop(5*time.Second, datavalue).Result()
			/*返回错误或超时(cmdstr==nil)*/
			if (nil != err) || (nil == cmdstr) {
				cmdRtnInfo.CmdrtnStr += "Get command result from db failed!\n"
				break
			}
			fmt.Printf("cmdnum:%d,rtnstring:%s\n", cmdnum, cmdstr[1])
			cmdnum--
			if strings.HasPrefix(cmdstr[1], CMD_END_RTNSTR) {
				/*去掉前缀获取返回码,如果返回码不等于succ,succ_deamon,warning,表示执行失败*/
				rtnstr = strings.TrimPrefix(cmdstr[1], CMD_END_RTNSTR)
				rtncode, _ = strconv.Atoi(rtnstr)
				/*frr define rtncode,0 SUCCESS,1 WARNING,10 SUCCESS_DEAMON*/
				if (0 != rtncode) && (1 != rtncode) && (10 != rtncode) {
					cmdRtnInfo.CmdRtncode = 1 /*execute failed*/
					break
				}
			} else {
				cmdRtnInfo.CmdrtnStr += cmdstr[1] + "\n"
			}
		}
		/*根据返回判断,write db; 发送到业务执行的命令,在数据库中存储命令字符串*/
		/*用于show running时显示*/
	} else {
		cmdRtnInfo.CmdrtnStr = "No command!"
	}

	return cmdRtnInfo, nil
}

func RedisTest(profile *pb.SONICCmdInputProfile) {
	clientcfgdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.103:6379",
		Password: "", // no password set
		DB:       4,  // use config DB,
	})

	//	clientappdb := redis.NewClient(&redis.Options{
	//		Addr:     "127.0.0.1:6379",
	//		Password: "",
	//		DB:       0,
	//	})

	pong, err := clientcfgdb.Ping().Result()
	fmt.Println(pong, err)

	err = clientcfgdb.Set(fmt.Sprintf("key%x", profile.CmdId), profile.CmdinputLine, 0).Err()
	if err != nil {
		panic(err)
	}

	//rdb.FlushDB()

	var cursor uint64
	var n int
	for {
		var keys []string
		var err error
		keys, cursor, err = clientcfgdb.Scan(cursor, "key*", 10).Result()
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

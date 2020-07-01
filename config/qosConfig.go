// qosconfig.go
// qos cml process, don't include db operate

package config

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	pb "hstcmler/cml"

	"github.com/go-redis/redis"
)

type QosConfig struct {
}

func NewQosConfig() *QosConfig {
	return new(QosConfig)
}

const (
	QosCarMaxNum   = 256
	QosDiffServNum = 6
)

//qos业务配置处理命令
func (self *QosConfig) SetQosMapCfg(ctx context.Context, req *pb.Qosdiffservmap) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	//var err error

	/*配置参数检查,避免参数间的冲突,以及业务之间的冲突*/
	if "" == req.GetDiffservName() {
		cmdRtnInfo.Rtnstr = "Error, no diffserv domain name!\n"
		cmdRtnInfo.Rtncode = 2
	} else {
		/*write db*/
		cmdRtnInfo.Rtnstr, _ = QosMapSetRedis(CFG_DB, req)
		cmdRtnInfo.Rtncode = 1
	}

	return &cmdRtnInfo, nil
}

func (self *QosConfig) DelQosMapCfg(ctx context.Context, req *pb.Qosdiffservmap) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	//var err error

	/*配置参数检查,避免参数间的冲突,以及业务之间的冲突*/
	if "" == req.GetDiffservName() {
		cmdRtnInfo.Rtnstr = "Error, no diffserv domain name!\n"
		cmdRtnInfo.Rtncode = 2
	} else {
		/*write db*/
		cmdRtnInfo.Rtnstr, _ = QosMapUndoSetRedis(CFG_DB, req)
		cmdRtnInfo.Rtncode = 1
	}

	return &cmdRtnInfo, nil
}

func (self *QosConfig) ShowQosMapCfg(ctx context.Context, req *pb.Qosservershowinput) (*pb.Qosservershowrtn, error) {

	var showRtninfo pb.Qosservershowrtn
	//var outstr string

	/*show config,statistic use rtn buffer*/
	/*show buffer use file return data*/
	showRtninfo.Cmdid = req.GetCmdid()
	showRtninfo.Rtnstr = "qos ShowQosMapCfg command return!!!"

	//switch req.GetCmdid() {
	//}
	//switch(req.GetShowoption()){
	/*过滤等选项*/
	//}

	return &showRtninfo, nil

}

func (self *QosConfig) SetDiffservifCfg(ctx context.Context, req *pb.Qosdiffservapply) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	var diffserv, iflist string

	/*配置参数检查*/
	diffserv = req.GetTrustdiffserv()
	iflist = req.GetIflist()
	if ("" == diffserv) || ("" == iflist) {
		cmdRtnInfo.Rtnstr = "Error, no diffserv domain name or if info!\n"
		cmdRtnInfo.Rtncode = 2
	} else {
		/*write db*/
		cmdRtnInfo.Rtnstr, _ = QosDiffIfSetRedis(CFG_DB, req)
		cmdRtnInfo.Rtncode = 1
	}

	return &cmdRtnInfo, nil
}

func (self *QosConfig) DelDiffservifCfg(ctx context.Context, req *pb.Qosdiffservapply) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	var diffserv, iflist string

	diffserv = req.GetTrustdiffserv()
	iflist = req.GetIflist()
	if ("" == diffserv) || ("" == iflist) {
		cmdRtnInfo.Rtnstr = "Error, no diffserv domain name or if info!\n"
		cmdRtnInfo.Rtncode = 2
	} else {
		/*write db*/
		cmdRtnInfo.Rtnstr, _ = QosDiffIfUndoSetRedis(CFG_DB, req)
		cmdRtnInfo.Rtncode = 1
	}

	return &cmdRtnInfo, nil
}

func (self *QosConfig) ShowDiffservifCfg(ctx context.Context, req *pb.Qosservershowinput) (*pb.Qosservershowrtn, error) {

	var showRtninfo pb.Qosservershowrtn
	//var outstr string

	/*show config,statistic use rtn buffer*/
	/*show buffer use file return data*/
	showRtninfo.Cmdid = req.GetCmdid()
	showRtninfo.Rtnstr = "qos ShowDiffservifCfg command return!!!"

	//switch req.GetCmdid() {
	//}
	//switch(req.GetShowoption()){
	/*过滤等选项*/
	//}

	return &showRtninfo, nil

}

func (self *QosConfig) SetQosLableModeCfg(ctx context.Context, req *pb.Labelqosmode) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	//var diffserv, iflist string

	/*配置参数检查*/
	//diffserv = req.GetTrustdiffserv()
	//iflist = req.GetIflist()
	//if ("" == diffserv) || ("" == iflist) {
	//	cmdRtnInfo.Rtnstr = "Error, no diffserv domain name or if info!\n"
	//	cmdRtnInfo.Rtncode = 2
	//} else {
	/*write db*/
	cmdRtnInfo.Rtnstr, _ = QosLableModeCfgSetRedis(CFG_DB, req)
	cmdRtnInfo.Rtncode = 1
	//}

	return &cmdRtnInfo, nil
}

func (self *QosConfig) DelQosLableModeCfg(ctx context.Context, req *pb.Labelqosmode) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	//var diffserv, iflist string

	//diffserv = req.GetTrustdiffserv()
	//iflist = req.GetIflist()
	//if ("" == diffserv) || ("" == iflist) {
	//	cmdRtnInfo.Rtnstr = "Error, no diffserv domain name or if info!\n"
	//	cmdRtnInfo.Rtncode = 2
	//} else {
	/*write db*/
	cmdRtnInfo.Rtnstr, _ = QosLableModeCfgUndoSetRedis(CFG_DB, req)
	cmdRtnInfo.Rtncode = 1
	//}

	return &cmdRtnInfo, nil
}

func (self *QosConfig) ShowQosLableModeCfg(ctx context.Context, req *pb.Qosservershowinput) (*pb.Qosservershowrtn, error) {

	var showRtninfo pb.Qosservershowrtn
	//var outstr string

	/*show config,statistic use rtn buffer*/
	/*show buffer use file return data*/
	showRtninfo.Cmdid = req.GetCmdid()
	showRtninfo.Rtnstr = "qos ShowDiffservifCfg command return!!!"

	//switch req.GetCmdid() {
	//}
	//switch(req.GetShowoption()){
	/*过滤等选项*/
	//}

	return &showRtninfo, nil

}

func (self *QosConfig) SetQosPhbInfoCfg(ctx context.Context, req *pb.Qosphbcfg) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	//var diffserv, iflist string

	/*配置参数检查*/
	//diffserv = req.GetTrustdiffserv()
	//iflist = req.GetIflist()
	//if ("" == diffserv) || ("" == iflist) {
	cmdRtnInfo.Rtnstr = "Error, no diffserv domain name or if info!\n"
	cmdRtnInfo.Rtncode = 2
	//} else {
	/*write db*/
	cmdRtnInfo.Rtnstr, _ = QosPhbInfoCfgSetRedis(CFG_DB, req)
	cmdRtnInfo.Rtncode = 1
	//}

	return &cmdRtnInfo, nil
}

func (self *QosConfig) DelQosPhbInfoCfg(ctx context.Context, req *pb.Qosphbcfg) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	//var diffserv, iflist string

	//diffserv = req.GetTrustdiffserv()
	//iflist = req.GetIflist()
	//if ("" == diffserv) || ("" == iflist) {
	cmdRtnInfo.Rtnstr = "Error, no diffserv domain name or if info!\n"
	cmdRtnInfo.Rtncode = 2
	//} else {
	/*write db*/
	cmdRtnInfo.Rtnstr, _ = QosPhbInfoCfgUndoSetRedis(CFG_DB, req)
	cmdRtnInfo.Rtncode = 1
	//}

	return &cmdRtnInfo, nil
}

func (self *QosConfig) ShowQosPhbInfoCfg(ctx context.Context, req *pb.Qosservershowinput) (*pb.Qosservershowrtn, error) {

	var showRtninfo pb.Qosservershowrtn
	//var outstr string

	/*show config,statistic use rtn buffer*/
	/*show buffer use file return data*/
	showRtninfo.Cmdid = req.GetCmdid()
	showRtninfo.Rtnstr = "qos ShowDiffservifCfg command return!!!"

	//switch req.GetCmdid() {
	//}
	//switch(req.GetShowoption()){
	/*过滤等选项*/
	//}

	return &showRtninfo, nil

}

func (self *QosConfig) SetQosCarGloCfg(ctx context.Context, req *pb.Qoscarglobalcfg) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	var fielddata string
	var dbclient *redis.Client
	var newadd, i int = 0, 0
	var carinfo []*pb.Qoscar
	var keys []string = nil
	var err error

	err = CmlDbReconnect(CFG_DB)
	if err != nil {
		cmdRtnInfo.Rtnstr = "db state error!\n"
		return &cmdRtnInfo, err
	}
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[CFG_DB]]

	carinfo = req.GetQoscarcfg()
	if nil != carinfo {
		/*从数据库中读出car name的总数量*/
		fielddata = KEY_QOS_POLICER + "*"
		keys, err = dbclient.Keys(fielddata).Result()
		if (nil == err) && (nil != keys) {
			/*配置参数检查,避免参数间的冲突,以及业务之间的冲突*/
			/*当输入多个car时,是不重名的,需要在输入端(cli,web)查重*/
			/*此处只检查总数是否超标,相同名的car是覆盖关系*/
			for _, datainfo := range carinfo {
				for _, carname := range keys {
					fielddata = strings.TrimLeft(carname, KEY_QOS_POLICER)
					/*检查是否存在,不存在就统计新增数*/
					if fielddata == datainfo.GetCarname() {
						break
					}
					i++
				}
				if i == len(keys) {
					newadd++
				}
			}
		} else {
			/*数据读取失败,或没有(keys==nil)全部认为是新增*/
			newadd = len(carinfo)
		}
	}
	/*color ifg的配置不需要查重或冲突检查,直接更新数据库*/

	/*判断car总数是否超标.不放在for循环中判断是因为多数情况都不存在超标,不用每次循环都*/
	/*判断,高效一点。另外,放在循环结束判断,代码实现流程更清晰*/
	/*如果本次没有配置qos car,则cartotal为0 ,不影响只配置color等的处理流程*/
	if QosCarMaxNum < newadd+len(keys) {
		/*总数超标,返回并提示*/
		cmdRtnInfo.Rtncode = 2
		cmdRtnInfo.Rtnstr = "Error, qos car exceed max number!\n"
	} else {
		/*write db,若原来有配置直接覆盖*/
		cmdRtnInfo.Rtnstr, err = QosCarGloSetRedis(CFG_DB, req)
		cmdRtnInfo.Rtncode = 1
	}

	/*作为数据库操作,需要发布更新,带上时间等*/

	return &cmdRtnInfo, nil
}

/*删除相关配置,不需要做特殊的处理*/
/*对color的动作是恢复默认值,默认值不用在呈现层补充,减少与呈现层的耦合*/
func (self *QosConfig) DelQosCarGloCfg(ctx context.Context, req *pb.Qoscarglobalcfg) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo

	cmdRtnInfo.Rtnstr, _ = QosCarGloSetRedis(CFG_DB, req)
	/*判断数据库操作结果,并发布更新*/
	cmdRtnInfo.Rtncode = 1

	return &cmdRtnInfo, nil
}

func (self *QosConfig) ShowQosCarGloCfg(ctx context.Context, req *pb.Qosservershowinput) (*pb.Qosservershowrtn, error) {

	var showRtninfo pb.Qosservershowrtn
	//var outstr string

	/*show config,statistic use rtn buffer*/
	/*show buffer use file return data*/
	showRtninfo.Cmdid = req.GetCmdid()
	showRtninfo.Rtnstr = "qos ShowQosCarGloCfg command return!!!"

	//switch req.GetCmdid() {
	//}
	//switch(req.GetShowoption()){
	/*过滤等选项*/
	//}

	return &showRtninfo, nil

}

func (self *QosConfig) SetQosCommIfCfg(ctx context.Context, req *pb.Qoscommifcfg) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	var qoscar, iflist string

	qoscar = req.GetQoscarname()
	iflist = req.GetIflist()
	if ("" == iflist) || ("" == qoscar) {
		cmdRtnInfo.Rtnstr = "Error, no qos car name or if info!\n"
		cmdRtnInfo.Rtncode = 2
	} else {
		/*write db*/
		cmdRtnInfo.Rtnstr, _ = QosCommIfSetRedis(CFG_DB, req)
		cmdRtnInfo.Rtncode = 1
	}

	return &cmdRtnInfo, nil
}

func (self *QosConfig) DelQosCommIfCfg(ctx context.Context, req *pb.Qoscommifcfg) (*pb.Qoscfgrtninfo, error) {

	var cmdRtnInfo pb.Qoscfgrtninfo
	var qoscar, iflist string

	qoscar = req.GetQoscarname()
	iflist = req.GetIflist()
	if ("" == iflist) || ("" == qoscar) {
		cmdRtnInfo.Rtnstr = "Error, no qos car name or if info!\n"
		cmdRtnInfo.Rtncode = 2
	} else {
		/*write db*/
		cmdRtnInfo.Rtnstr, _ = QosCommIfUndoSetRedis(CFG_DB, req)
		cmdRtnInfo.Rtncode = 1
	}

	return &cmdRtnInfo, nil
}

func (self *QosConfig) ShowQosCommIfCfg(ctx context.Context, req *pb.Qosservershowinput) (*pb.Qosservershowrtn, error) {

	var showRtninfo pb.Qosservershowrtn
	//var outstr string

	/*show config,statistic use rtn buffer*/
	/*show buffer use file return data*/
	showRtninfo.Cmdid = req.GetCmdid()
	showRtninfo.Rtnstr = "qos ShowQosCarGloCfg command return!!!"

	//switch req.GetCmdid() {
	//}
	//switch(req.GetShowoption()){
	/*过滤等选项*/
	//}

	return &showRtninfo, nil

}

func QosShowStatistic() string {
	/*return filename(ms.bak), in klish,cp file-bak,and read to display, then delete*/
	/*mulit vty has problem, maybe conflict*/
	var logfile, showfile *os.File
	var err error

	tm := time.Now()
	timestamp := tm.Unix()
	bakName := fmt.Sprintf("%s%d.bak", "/var/log/", timestamp)

	fileName := fmt.Sprintf("%s_%04d%02d%02d.log", "/var/log/sonic", tm.Year(), tm.Month(), tm.Day())
	logfile, err = os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("syslog command,open Log File Failed, fileName=%s, err:%s\n", fileName, err)
		return "No buffer data"
	}
	defer logfile.Close()

	showfile, err = os.OpenFile(bakName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("Open write file:%s failed, err:%s\n", fileName, err)
		return "Get data file failed!"
	}
	defer showfile.Close()

	_, err = io.Copy(showfile, logfile)
	if err != nil {
		fmt.Printf("write data to file:%s failed, err:%s\n", fileName, err)
		return "Get data file failed!"
	}

	return bakName
}

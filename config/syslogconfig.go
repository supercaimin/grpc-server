// syslogconfig.go
// syslog cml process

package config

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	pb "hstcmler/cml"
	//"github.com/go-redis/redis"
)

//sonic的命令来源信息以及业务对处理的返回信息等
//因为对SONIC原生命令，CML处理后还需要业务继续处理。因对原生的命令处理，配置的有效性检查需要
//到业务处处理才行(或者就需要业务提供配置检查,或者把klish中konfd的纯文本检查移植过来，选择？
//
type SyslogConfig struct {
}

func NewSyslogConfig() *SyslogConfig {
	return new(SyslogConfig)
}

//命令的处理分发函数
func (self *SyslogConfig) SyslogSetConfigure(ctx context.Context, req *pb.SyslogConfProfile) (*pb.SyslogConfRtnProfile, error) {

	var cmdRtnInfo pb.SyslogConfRtnProfile

	cmdRtnInfo.Rtncode = 1

	/*write db*/
	cmdRtnInfo.Rtnstring, _ = SyslogSetRedis(CFG_DB, req)
	/*不能把err的信息透传到客户端,只把错误信息的字符串转换为提示信息带回*/
	return &cmdRtnInfo, nil
}

func (self *SyslogConfig) SyslogUndoSetConfigure(ctx context.Context, req *pb.SyslogConfProfile) (*pb.SyslogConfRtnProfile, error) {

	var cmdRtnInfo pb.SyslogConfRtnProfile

	cmdRtnInfo.Rtncode = 2
	cmdRtnInfo.Rtnstring = "Syslog command rcv ok!!!"

	/*write db*/
	cmdRtnInfo.Rtnstring, _ = SyslogUndoSetRedis(CFG_DB, req)
	/*不能把err的信息透传到客户端,只把错误信息的字符串转换为提示信息带回*/
	return &cmdRtnInfo, nil

}

func (self *SyslogConfig) LogLevelModify(ctx context.Context, req *pb.LogLevelProfile) (*pb.SyslogConfRtnProfile, error) {

	var cmdRtnInfo pb.SyslogConfRtnProfile

	cmdRtnInfo.Rtncode = 2
	cmdRtnInfo.Rtnstring = "Syslog command rcv ok!!!"

	/*write db*/
	cmdRtnInfo.Rtnstring, _ = LogLevelSetRedis(CFG_DB, req)
	/*不能把err的信息透传到客户端,只把错误信息的字符串转换为提示信息带回*/
	return &cmdRtnInfo, nil

}

func (self *SyslogConfig) LogLevelUndoModify(ctx context.Context, req *pb.LogLevelProfile) (*pb.SyslogConfRtnProfile, error) {

	var cmdRtnInfo pb.SyslogConfRtnProfile

	cmdRtnInfo.Rtncode = 2
	cmdRtnInfo.Rtnstring = "Syslog command rcv ok!!!"

	/*write db*/
	cmdRtnInfo.Rtnstring, _ = LogLevelUndoSetRedis(CFG_DB, req)
	/*不能把err的信息透传到客户端,只把错误信息的字符串转换为提示信息带回*/
	return &cmdRtnInfo, nil

}

func (self *SyslogConfig) LogLevelshow(ctx context.Context, req *pb.LogLevelShowProfile) (*pb.SyslogShowRtnProfile, error) {

	var showRtninfo pb.SyslogShowRtnProfile

	/*show config,statistic use rtn buffer*/
	/*show buffer use file return data*/

	showRtninfo.CmdrtnStr, _ = LogLevelShow(CFG_DB, req)
	return &showRtninfo, nil

}

func (self *SyslogConfig) Syslogshow(ctx context.Context, req *pb.SyslogShowProfile) (*pb.SyslogShowRtnProfile, error) {

	var showRtninfo pb.SyslogShowRtnProfile

	/*show config,statistic use rtn buffer*/
	/*show buffer use file return data*/
	showRtninfo.ShowOption = req.GetShowOption()
	showRtninfo.CmdrtnStr = "syslog show command return!!!"

	switch req.GetShowOption() {
	case pb.LOG_SHOW_LOG_SHOW_CURRENT_ALARM: //show current alarm
		showRtninfo.CmdrtnStr = SyslogShowAlarm(STATE_DB)
		break
	case pb.LOG_SHOW_LOG_SHOW_CURRENT_LOG: //show current log
		//use db to trans data
		//SyslogShowstatis(showRtninfo.CmdrtnStr)
		/*show buffer,display buffer file*/
		showRtninfo.CmdrtnStr = SyslogShowBufferfile()
		break
	case pb.LOG_SHOW_LOG_SHOW_CONFIG:
		SyslogShowConfig(CFG_DB, &showRtninfo.CmdrtnStr)
		break
	case pb.LOG_SHOW_LOG_SHOW_CLEARSTATIS:
		break
	case pb.LOG_SHOW_LOG_SHOW_SAVEFILE:
		break
	default:
		break
	}

	return &showRtninfo, nil

}

func SyslogShowBufferfile() string {
	/*return filename(ms.bak), in klish,cp file-bak,and read to display, then delete*/
	/*mulit vty has problem, maybe conflict*/
	var logfile, showfile *os.File
	var err error

	tm := time.Now()
	timestamp := tm.Unix()
	bakName := fmt.Sprintf("%s%d.bak", "/var/log/", timestamp)

	showfile, err = os.OpenFile(bakName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("Open write file:%s failed, err:%s\n", bakName, err)
		return ""
	}
	defer showfile.Close()

	fileName := fmt.Sprintf("%s_%04d%02d%02d.log", "/var/log/sonic", tm.Year(), tm.Month(), tm.Day())
	logfile, err = os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("syslog command,open Log File Failed, fileName=%s, err:%s\n", fileName, err)
		showfile.WriteString("Open log file failed, No data\n")
		return bakName
	}
	defer logfile.Close()

	_, err = io.Copy(showfile, logfile)
	if err != nil {
		fmt.Printf("write data to file:%s failed, err:%s\n", fileName, err)
		showfile.WriteString("Get data file failed, No data\n")
		return bakName
	}

	return bakName
}

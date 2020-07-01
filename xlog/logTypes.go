/* logTypes.go
	作者：	汪军
	日期：2019-9-4

Log模块常数定义

*/

package xlog

import (
	"net"
	"time"

	"hstcmler/errcode"
)

type LOG_LEVEL int

const (
	LOG_LEVEL_EMERGENCY LOG_LEVEL = 0
	LOG_LEVEL_ALERT               = 1
	LOG_LEVEL_CRITICAL            = 2
	LOG_LEVEL_ERROR               = 3
	LOG_LEVEL_WARNING             = 4
	LOG_LEVEL_NOTICE              = 5
	LOG_LEVEL_INFO                = 6
	LOG_LEVEL_DEBUG               = 7
)

const (
	LOG_TYPE_SYSTEM    int = 0
	LOG_TYPE_OPERATION     = 1
	LOG_TYPE_SECURITY      = 2
	LOG_TYEP_ALARM         = 3
)

/*Log Service的常量定义UDP端口号*/

const (
	LOG_DEF_SERVICE_PORT               = 10000
	LOG_MAX_MSG_SIZE                   = 32768
	LOG_MAX_STORAGE_SIZE         int64 = 32 * (1 << 20)    //日志存储的最大磁盘空间，字节作为单位
	LOG_MAX_STORAGE_DURATION           = 30                //日志存储的默认最大时间
	LOG_DEF_PERSISTENT_LEVEL           = LOG_LEVEL_WARNING //日志默认存盘的级别
	LOG_DEF_CONSOLE_OUTPUT_LEVEL       = LOG_LEVEL_INFO    //日志默认输出到控制台的级别
	SYSLOG_DEF_UDP_PORT                = 514
)

/*内部传递的Log数据结构*/
type InnerLogMsg struct {
	ModuleName string         `ModuleName`
	OccureTime time.Time      `OccureTime`
	LogType    int            `LogType`
	LogLevel   LOG_LEVEL      `LogLevel`
	Position   string         `Postision`
	ErrCode    errcode.RESULT `ErrCode`
	LogDesc    string         `LogDesc`
	Context    string         `Context`
	senderAddr net.UDPAddr
	jsonMsgLen int
}

/*Log 监听器接口*/
type LogListener interface {
	/*Log Service回调接口，监听者如果返回true，表示下一步LogService继续处理剩余逻辑；否则终结日志处理
	msgSeq是消息的唯一序列号，上电后单增*/
	NotifyLog(msg *InnerLogMsg, msgSeq int64) bool
}

/* syslogclient.go
	作者：	汪军
	日期：2019-9-4

sysLog标准协议客户端代码，用于和外部Syslog Server进行通信，输出日志
可以配置最大4个Syslog Server，并行发送日志
遵循RFC 5424/5426协议

*/

package xlog

import (
	"fmt"
)

/*将Log的结构体格式化为字符串，符合RFC5424的SysLog格式
存盘采用此格式*/
func LogFormat(msgStru *InnerLogMsg, msgId int64) string {
	var strMsg string = ""

	/*SYSLOG-MSG = HEADER SP STRUCTURED-DATA [SP MSG]
	  HEADER = PRI VERSION SP TIMESTAMP SP HOSTNAME
	  SP APP-NAME SP PROCID SP MSGID */
	strHdr := fmt.Sprintf("<%d> 1 %s %s:%d %s %s %d", 1*8+msgStru.LogLevel, msgStru.OccureTime.Format("2006-01-02 15:04:05.999999Z-0700"),
		msgStru.senderAddr.IP.String(), msgStru.senderAddr.Port, msgStru.ModuleName, msgStru.Position, msgId)

	strMsg = strHdr + fmt.Sprintf(" ErrorCode=\"%d\" %s %s", msgStru.ErrCode, msgStru.LogDesc, msgStru.Context)
	return strMsg
}

/* logClient.go
	作者：	汪军
	创建日期：2019-9-4
	最新修订日期：2019-9-16

   模块功能描述：Log客户端代码，由每个应用自己负责调用；Log客户端和服务端通过网络进行通信，控制级别
每个模块首先调用NewXLogger，传入本地的模块名、本地IP地址、端口，LogService的IP和端口；
如果不需要跨节点通信本地地址可以填写127.0.0.1,如不关心本地地址的可以填写0.0.0.0
*/

package xlog

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"hstcmler/errcode"
)

/*Log实例对象结构*/
type XLogger struct {
	moduleName             string
	selfIp                 net.IP
	selfPort               int
	sentPackets, sentbytes int64
	udpconn                *net.UDPConn
	servAddr               net.UDPAddr
}

/*初始化Log，每个模块使用Log前要初始化一个实例，用此实例输出日志,入参为自己的模块名称，自身IP和端口，LogService的IP和端口
支持LogClient和Service在不同的节点中部署*/
func NewXLogger(moduleName string, selfIp string, selfPort int, servIp string, servPort int) *XLogger {

	localIp := net.ParseIP(selfIp)
	rIp := net.ParseIP(servIp)

	conn := initUdpConn(localIp, rIp, selfPort, servPort)
	if conn == nil {
		return nil
	}

	logger := XLogger{moduleName: moduleName, selfIp: localIp,
		selfPort: selfPort, sentPackets: 0, sentbytes: 0}
	logger.udpconn = conn
	logger.servAddr.IP = rIp
	logger.servAddr.Port = servPort

	return &logger
}

/*初始化UDP连接*/
func initUdpConn(srcIp, dstIp net.IP, srcPort, dstPort int) *net.UDPConn {
	dst := net.UDPAddr{IP: dstIp, Port: dstPort}
	src := net.UDPAddr{IP: srcIp, Port: srcPort}

	fmt.Printf("src ip: %d", src.IP)

	//conn, err := net.DialUDP("udp", &src, &dst)
	conn, err := net.DialUDP("udp", nil, &dst)

	fmt.Printf("Log Init Udp Socket ip=%s, port=%d\n", dstIp.String(), dstPort)

	if err != nil {
		fmt.Printf("Log Init Udp Socket Error, ip=%s, port=%d,err:%s", srcIp.String(), srcPort, err)
		return nil
	}

	return conn
}

/*日志接口，不同的日志等级区分处理*/
func (self *XLogger) Alertf(postion string, err errcode.RESULT, logDesc string, v ...interface{}) {
	desc := fmt.Sprintf(logDesc, v...)
	self.Log(LOG_LEVEL_ALERT, postion, err, desc, "")
}

func (self *XLogger) Errorf(postion string, err errcode.RESULT, logDesc string, v ...interface{}) {
	desc := fmt.Sprintf(logDesc, v...)
	self.Log(LOG_LEVEL_ERROR, postion, err, desc, "")
}

func (self *XLogger) Warningf(postion string, err errcode.RESULT, logDesc string, v ...interface{}) {
	desc := fmt.Sprintf(logDesc, v...)
	self.Log(LOG_LEVEL_WARNING, postion, err, desc, "")
}

func (self *XLogger) Infof(postion string, err errcode.RESULT, logDesc string, v ...interface{}) {
	desc := fmt.Sprintf(logDesc, v...)
	self.Log(LOG_LEVEL_INFO, postion, err, desc, "")
}

func (self *XLogger) Debugf(postion string, err errcode.RESULT, logDesc string, v ...interface{}) {
	desc := fmt.Sprintf(logDesc, v...)
	self.Log(LOG_LEVEL_DEBUG, postion, err, desc, "")
}

/*操作日志，和普通系统日志分开*/
func (self *XLogger) OpLog(postion string, logDesc string, v ...interface{}) {
	desc := fmt.Sprintf(logDesc, v...)

	logMsg := InnerLogMsg{LogType: LOG_TYPE_OPERATION, Position: postion, LogDesc: desc}
	self.sentLog(&logMsg)

}

/*系统日志输出*/
func (self *XLogger) Log(level LOG_LEVEL, position string, err errcode.RESULT, logDesc string, context string) {

	logMsg := InnerLogMsg{LogType: LOG_TYPE_SYSTEM, LogLevel: level, Position: position, ErrCode: err,
		LogDesc: logDesc, Context: context}

	self.sentLog(&logMsg)
}

/*获取当前日志模块的输出统计数据*/
func (self *XLogger) GetLogStats() (sentPackets, sentBytes int64) {
	return self.sentPackets, self.sentbytes
}

/*设置新的服务端IP和端口*/
func (self *XLogger) SetNewServer(newIp net.IP, port int) {
	if self.udpconn != nil {
		self.udpconn.Close()
	}

	self.servAddr.IP = newIp
	self.servAddr.Port = port

	self.udpconn = initUdpConn(self.selfIp, newIp, self.selfPort, port)

}

/*发送日志给LogService*/
func (self *XLogger) sentLog(msg *InnerLogMsg) errcode.RESULT {

	msg.OccureTime = time.Now()
	msg.ModuleName = self.moduleName

	bstr, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("LogClient: Sent Log error, err=%s\n", err)
		return errcode.RESULT_ERROR_FAILED_ENCODING
	}
	if self.udpconn == nil {
		fmt.Printf("LogClient: udp connection is not ready, ip=%s,port=%d\n", self.selfIp.String(), self.selfPort)
		return errcode.RESULT_ERROR_CONNECTION_CLOSED
	}

	len, err := self.udpconn.Write(bstr)

	if len <= 0 {
		fmt.Printf("LogClient: Sent Log error, err=%s\n", err)
		return errcode.RESULT_ERROR_COMMON
	}
	self.sentbytes += int64(len)
	self.sentPackets++
	return errcode.RESULT_SUCCESS
}

func (self *XLogger) PrintStats() {
	fmt.Printf("LogClient: ClientIP=%s:%d, Server=%s:%d\n", self.selfIp, self.selfPort,
		self.servAddr.IP.String(), self.servAddr.Port)
	fmt.Printf("LogClient: sent packets=%d, bytes=%d\n", self.sentPackets, self.sentbytes)
}

//
/* LogService.go
	作者：	汪军
	日期：2019-9-4

Log服务端，负责日志的统一存储，按照模块分类管理,Log客户端和服务端通过网络进行通信，支持跨节点部署。
系统一个节点只支持一个LogService实例，可以用InitLogService进行初始化
支持注册Log事件监听器，方便应用进行进一步开发

*/

package xlog

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"hstcmler/errcode"
)

const (
	MAX_LOG_QUEUE_LEN    = 2048
	MAX_LOG_MSG_LEN      = 65000
	MAX_OOB_MSG_LEN      = 4096
	MAX_UNSYNCED_MSG     = 256
	MAX_OUT_LOG_SERVER   = 2 //最多支持同时向2个2外部服务器发syslog
	LOG_SYNC_DISK_PERIOD = 2 //刷新到磁盘的周期
)

type logModule struct {
	persistent_log_level LOG_LEVEL
	console_log_level    LOG_LEVEL
	name                 string //客户模块的名字
	ip                   net.IP //Log客户模块的IP地址
	port                 int    //Log客户模块的UDP端口号
	logPackets, logBytes int64  //发送Log的条数和
}

/*chenwei add for send to out syslog server*/
/*ipv6在第二阶段完善*/
type logOutServer struct {
	vrfname                string //when global, vrfname is ""
	selfIp                 net.IP //source ip
	selfPort               int    //source port
	sentPackets, sentbytes int64
	udpconn                *net.UDPConn
	servAddr               net.UDPAddr
}

type TLogService struct {
	persistentLevel   LOG_LEVEL
	consoleLevel      LOG_LEVEL
	monitorLevel      LOG_LEVEL
	trapLevel         LOG_LEVEL
	serverLevel       LOG_LEVEL
	maxStorageSize    int64 //默认磁盘空间，以MB计算
	maxStoragePeriod  int   //默认存储时间，以天计算
	isInit            bool
	selfAddr          net.UDPAddr
	conn              *net.UDPConn
	logFilePath       string    //日志文件的存放路径
	logFilePrefix     string    //日志文件的名称前缀，比如anpc，系统会自动加上日期及扩展名，比如anpc_20190311.log
	logFileCreateTime time.Time //当前日志文件的创建时间
	curLogFile        *os.File  //当前打开的文件句柄
	unSyncedMsg       int       //已经写入文件，但是没有存盘的消息计数
	lastWriteTime     time.Time
	fileMutex         sync.Mutex
	outSyslogServer   []logOutServer //发送外部服务器的信息

	queue      *LogQueue
	alarmqueue *LogQueue /*告警队列*/

	logBuf1, logBuf2 []byte

	logMuduleControl     map[string]*logModule //以名称索引到日志的客户模块信息
	logModuleIndex       map[string]string     //以IP:端口索引到名称
	logListener          map[string]*LogListener
	alarmCodeLevel       map[string]LOG_LEVEL /*告警码与level关系,[module|code]-level*/
	logPackets, LogBytes int64
	logSeq               int64 //Log Msg序列号
}

var logService = TLogService{persistentLevel: LOG_LEVEL_ERROR, consoleLevel: LOG_LEVEL_INFO, monitorLevel: LOG_LEVEL_INFO,
	maxStorageSize: LOG_MAX_STORAGE_SIZE, maxStoragePeriod: LOG_MAX_STORAGE_DURATION, isInit: false}

/*初始化日志服务器，由上层应用调用初始化，本接口启动日志接收
logFilePath是指"/var/log/"这样的文件路径，logFilePrefix是指文件名称前缀，比如anpc，系统会生成"anpc_YYYYMMDD.log"这样的文件名
系统一天使用一个Log文件名称，并对前一天的日志进行压缩打包
*/
func InitLogService(localIp string, port int, logFilePath string, logFilePrefix string) errcode.RESULT {

	var err error
	if logService.isInit {
		return errcode.RESULT_ERROR_COMMON
	}
	logService.selfAddr.IP = net.ParseIP(localIp)
	logService.selfAddr.Port = port
	logService.conn, err = net.ListenUDP("udp", &logService.selfAddr)
	if err != nil {
		fmt.Printf("Init LogService Error, Err:%s\n", err)
		return errcode.RESULT_ERROR_COMMON
	}

	/*设置系统默认的日志输出级别*/
	logService.consoleLevel, logService.persistentLevel = LOG_LEVEL_INFO, LOG_LEVEL_WARNING
	logService.logFilePath = formatLogPath(logFilePath)

	logService.logFilePrefix = logFilePrefix
	ret := openLogFile()
	if ret != errcode.RESULT_SUCCESS {
		return ret
	}
	//chenwei add for outer syslog server
	logService.outSyslogServer = make([]logOutServer, MAX_OUT_LOG_SERVER)

	logService.lastWriteTime = time.Now()
	logService.unSyncedMsg = 0

	logService.logListener = make(map[string]*LogListener)
	logService.logModuleIndex = make(map[string]string)
	logService.logMuduleControl = make(map[string]*logModule)
	logService.isInit = true

	logService.alarmCodeLevel = make(map[string]LOG_LEVEL, 200)

	logService.logBuf1 = make([]byte, MAX_LOG_MSG_LEN)
	logService.logBuf2 = make([]byte, MAX_OOB_MSG_LEN)

	/*初始化日志队列，收到消息后即放入消息队列，由单独的日志处理进程进行处理*/
	logService.queue = NewLogQueue(MAX_LOG_QUEUE_LEN)
	logService.alarmqueue = NewLogQueue(MAX_LOG_QUEUE_LEN)

	/*init config db subscribe*/
	InitCmlDbInfo()
	go LogCfgDbListen()

	go recvLog()      //收日志消息的任务
	go processLog()   //处理日志消息的任务，和收日志消息异步处理
	go processAlarm() //处理告警的任务,告警接收与日志走同一套口,收到后告警入alarmqueue
	go flushLogFile() //周期性日志存盘任务

	InitLogCleanTask(&logService) //初始化清理任务，定期清理任务

	return errcode.RESULT_SUCCESS

}

/*设置日志存储的参数，0表示不修改*/
func SetLogPersitentParam(maxDiskSize int64, maxPeriod int) {
	if maxDiskSize >= 0 {
		logService.maxStorageSize = maxDiskSize
	}
	if maxPeriod >= 0 {
		logService.maxStoragePeriod = maxPeriod
	}
}

/*设置全局日志级别,分为存盘级别和控制台输出级别*/
func SetGlobalPersitentLogLevel(newLevel LOG_LEVEL) {
	logService.persistentLevel = newLevel

}

func SetGlobalConsoleLogLevel(newLevel LOG_LEVEL) {
	logService.consoleLevel = newLevel

}

/*设置模块级的持久化日志级别*/
func SetModulePersitentLogLevel(name string, newLevel LOG_LEVEL) errcode.RESULT {
	v, ok := logService.logMuduleControl[name]
	if !ok {
		logService.logMuduleControl[name] = &logModule{name: name, persistent_log_level: newLevel,
			console_log_level: LOG_DEF_CONSOLE_OUTPUT_LEVEL}
	} else {
		v.persistent_log_level = newLevel
	}

	return errcode.RESULT_SUCCESS

}

/*设置模块级的控制台输出日志级别*/
func SetModuleConsoleLogLevel(name string, newLevel LOG_LEVEL) errcode.RESULT {
	v, ok := logService.logMuduleControl[name]
	if !ok {
		logService.logMuduleControl[name] = &logModule{name: name, console_log_level: newLevel,
			persistent_log_level: LOG_DEF_PERSISTENT_LEVEL}
	} else {
		v.console_log_level = newLevel
	}

	return errcode.RESULT_SUCCESS
}

/*日志监听器注册，LogService将所有的日志均发送给监听器*/
func RegisterListener(name string, listener *LogListener) errcode.RESULT {

	_, found := logService.logListener[name]
	if found {
		return errcode.RESULT_ERROR_ALREADY_EXIST
	}
	logService.logListener[name] = listener
	return errcode.RESULT_SUCCESS
}

/*监听器去注册*/
func DeregisterListener(name string) errcode.RESULT {

	_, found := logService.logListener[name]
	if !found {
		return errcode.RESULT_ERROR_NOT_FOUND
	}
	delete(logService.logListener, name)
	return errcode.RESULT_SUCCESS
}

/*日志接收处理,JSON格式
首先进行回调处理，然后进行存盘处理;最后发送给SysLog Server*/
func recvLog() {

	fmt.Println("LogServer: Begin Receiving Log Message")

	for {
		logService.conn.SetDeadline(time.Now().Add(1 * time.Second))
		n1, srcAddr, err := logService.conn.ReadFromUDP(logService.logBuf1)
		if err != nil {
			if !os.IsTimeout(err) {
				fmt.Printf("LogService,recv msg err:%s\n", err)
			}

			continue
		}

		fmt.Printf("udp log raw info:%s\n", logService.logBuf1)

		logStru := new(InnerLogMsg)
		err = json.Unmarshal(logService.logBuf1[0:n1], logStru)
		if err != nil {
			fmt.Printf("LogService,Decoding Json err:%s\n", err)
			continue
		}

		logService.logPackets++
		logService.LogBytes += int64(n1)
		logStru.jsonMsgLen = n1

		logStru.senderAddr = *srcAddr

		/*告警日志单独一个队列*/
		if logStru.LogType == LOG_TYEP_ALARM {
			logService.alarmqueue.PushBack(logStru)
		} else {
			logService.queue.PushBack(logStru)
		}

	}

}

/*真正的处理Log日志的任务，初始化日志实例时创建线程任务运行*/
func processLog() {

	var (
		ret bool
		msg *InnerLogMsg
	)

	for {
		/*每次调度等待100ms*/
		time.Sleep(100 * time.Millisecond)
		for {
			/*首先从日志消息队列中取消息，日志消息队列默认长度为MAX_LOG_QUEUE_LEN,则每秒钟处理日志的最大能力为10*MAX_LOG_QUEUE_LEN*/
			msg = logService.queue.Pop()
			/*假如队列已空，则退出，等待下一次调度*/
			if msg == nil {
				break
			}
			logService.logSeq++
			ret = processListener(msg, logService.logSeq)
			if !ret {
				continue
			}

			/*每次都调用RegisterLogModule是为解决先在服务端初始化模块级日志，后收到日志的问题*/
			m := RegisterLogModule(msg.ModuleName, &msg.senderAddr)

			if m != nil {
				m.logPackets++
				m.logBytes += int64(msg.jsonMsgLen)
			}
			persistentLog(msg, m, logService.logSeq)    //首先进行持久化处理
			consoleOutputLog(msg, m, logService.logSeq) //然后进行控制台处理
			serverOutputLog(msg, m, logService.logSeq)  //进行外部server发送处理
			snmpTrapLog(msg)                            //snmp trap
		}
	}

}

/*处理告警的任务，初始化日志实例时创建线程任务运行*/
/*将接收到的告警在数据库中建表记录,状态是消除的删除数据库中记录*/
/*所有告警(含产生和消除)都同时会产生一条日志记录到syslog中,模块和级别都复用告警信息*/
/*陈维 20200245*/
func processAlarm() {

	var (
		msg       *InnerLogMsg
		codeindex string
	)

	for {
		/*每次调度等待100ms*/
		time.Sleep(100 * time.Millisecond)
		/*处理告警*/
		for {
			/*从告警消息队列中取消息，队列默认长度为MAX_LOG_QUEUE_LEN,则每秒钟处理告警的最大能力为10*MAX_LOG_QUEUE_LEN*/
			msg = logService.alarmqueue.Pop()
			/*假如队列已空，则退出，等待下一次调度*/
			if msg == nil {
				break
			}
			/*接收的告警先进行状态更新操作处理,然后产生一条日志*/
			/*其它就直接按syslog日志处理,不做额外的操作*/
			/*告警数据直接在STATE-DB表中进行操作,不在内存中进行保留处理*/
			AlarmStateUpdate(msg) //告警的db处理,不检查level
			//将类型改为syslog系统日志,信息压入syslog日志队列
			msg.LogType = LOG_TYPE_SYSTEM
			/*根据modulename+code获取level.告警发送借用日志通道level设置为1,现恢复定义值*/
			codeindex = msg.ModuleName + "|" + fmt.Sprintf("%d", msg.ErrCode)
			if _, ok := logService.alarmCodeLevel[codeindex]; ok {
				msg.LogLevel = logService.alarmCodeLevel[codeindex]
			} else {
				msg.LogLevel = LOG_LEVEL_WARNING //没找到,告警对应的日志严重级别设为4
			}
			logService.queue.PushBack(msg)
		}
	}

}

/*处理监听器的回调，被调用者返回false，则表示流程结束*/
func RegisterLogModule(name string, addr *net.UDPAddr) *logModule {

	pM, ok := logService.logMuduleControl[name]
	if !ok {
		m := logModule{persistent_log_level: LOG_LEVEL_WARNING, console_log_level: LOG_LEVEL_INFO,
			name: name, ip: addr.IP, port: addr.Port}

		logService.logMuduleControl[name] = &m
		return &m

	} else if pM.port == 0 {
		pM.ip = addr.IP
		pM.port = addr.Port
		logService.logModuleIndex[fmt.Sprintf("%s:%d", addr.IP.String(), addr.Port)] = name
	}

	return pM

}

/*处理监听器的回调，被调用者返回false，则表示流程结束*/
func processListener(msgStru *InnerLogMsg, msgSeq int64) bool {
	var ret bool

	for _, v := range logService.logListener {
		ret = (*v).NotifyLog(msgStru, msgSeq)
		if !ret {
			return false
		}
	}
	return true
}

/*周期性日志存盘的操作*/
func flushLogFile() {

	for {
		time.Sleep(time.Second * 1) //休眠1秒钟
		//如果超过一定的时间没有刷新磁盘，则强制刷新一次Log到磁盘
		cur := time.Now()

		if logService.unSyncedMsg <= 0 {
			continue
		}

		if logService.unSyncedMsg >= MAX_UNSYNCED_MSG ||
			cur.After(logService.lastWriteTime.Add(time.Second*LOG_SYNC_DISK_PERIOD)) {
			logService.fileMutex.Lock()
			forceSyncLogFile()
			logService.fileMutex.Unlock()
		}

	}

}

/*处理日志持久化流程，首先判断级别是否够*/
/*若将是否创建日志文件移到外部处理,减少每条记录都进行判断,则存在风险,因为Logclean任务压缩后
会删除文件,如果正好在24小时的点上就存在多条记录持久化中途文件被删除,导致写文件异常 */
func persistentLog(msgStru *InnerLogMsg, m *logModule, msgSeq int64) {

	if msgStru.LogLevel > logService.persistentLevel {
		return
	}
	if m != nil && msgStru.LogLevel > m.persistent_log_level {
		return
	}

	logService.fileMutex.Lock()
	defer logService.fileMutex.Unlock()

	if logService.curLogFile == nil {
		return
	}
	cur := time.Now()
	/*已经是另外一天，需要打开新的日志文件*/
	if (cur.YearDay() != logService.lastWriteTime.YearDay()) || (cur.Year() != logService.lastWriteTime.Year()) {
		forceSyncLogFile()
		logService.curLogFile.Close()
		/*自动根据当前日期计算一个新的日志文件*/
		err := openLogFile()
		if err != errcode.RESULT_SUCCESS {
			fmt.Println("Create new log file failed!!!\n")
		}
	}

	strLog := LogFormat(msgStru, msgSeq) + "\n"
	logService.curLogFile.WriteString(strLog)
	logService.unSyncedMsg++
	logService.lastWriteTime = time.Now()

}

/*强制刷新到磁盘文件，将未同步的日志文件，刷新到磁盘*/
func forceSyncLogFile() {
	logService.curLogFile.Sync()
	logService.unSyncedMsg = 0

}

/*处理日志控制台输出流程，首先判断日志输出级别是否可以输出*/
func consoleOutputLog(msgStru *InnerLogMsg, m *logModule, msgSeq int64) {
	if msgStru.LogLevel > logService.consoleLevel {
		return
	}
	if m != nil && msgStru.LogLevel > m.console_log_level {
		return
	}
	fmt.Println(LogFormat(msgStru, msgSeq))

}

/*snmp trap*/
func snmpTrapLog(msgStru *InnerLogMsg) {
	var sver net.UDPAddr

	if msgStru.LogLevel > logService.trapLevel {
		return
	}
	codestr := fmt.Sprintf("%d", msgStru.ErrCode)
	levelstr := fmt.Sprintf("%d", msgStru.LogLevel)
	for _, v := range logService.outSyslogServer {
		/*port 为0表示无效配置*/
		if 0 != (v).servAddr.Port {
			sver = (v).servAddr
			sver.Port = 162 //trap 162 port
			snmptrap(msgStru.ModuleName, codestr, levelstr, msgStru.LogDesc, true, sver)
		}
	}

}

/*向syslog 服务器发送流程*/
func serverOutputLog(msgStru *InnerLogMsg, m *logModule, msgSeq int64) {
	if msgStru.LogLevel > logService.serverLevel {
		return
	}

	strLog := LogFormat(msgStru, msgSeq) + "\n"

	for i, v := range logService.outSyslogServer {
		if 0 != (v).servAddr.Port {
			/*先判断不等于,是大多数情况都属于此,减少多次重复判断.但这导致发送处理代码有重复*/
			if nil != (v).udpconn {
				len, _ := (v).udpconn.Write([]byte(strLog))

				if len > 0 {
					logService.outSyslogServer[i].sentbytes += int64(len)
					logService.outSyslogServer[i].sentPackets++
				}
			} else {
				/*还有考虑带vpn的情况*/
				conn := initUdpConn(logService.selfAddr.IP, (v).servAddr.IP,
					logService.selfAddr.Port, (v).servAddr.Port)
				if nil != conn {
					logService.outSyslogServer[i].udpconn = conn
					len, _ := conn.Write([]byte(strLog))
					if len > 0 {
						logService.outSyslogServer[i].sentbytes += int64(len)
						logService.outSyslogServer[i].sentPackets++
					}
				}
			}
		}
	}
}

/*格式化处理日志目录，包括处理空路径、添加遗失的"/"等操作*/
func formatLogPath(filePath string) string {

	str := strings.Trim(filePath, " ")
	if len(str) == 0 {
		str = "/var/log/"
	} else {
		if !strings.HasSuffix(str, "/") {
			str += "/"
		}
	}
	return str
}

/*根据自动生成的文件名，打开日志文件*/
func openLogFile() errcode.RESULT {

	var err error
	tm := time.Now()

	fileName := logService.logFilePath + getLogFileName(&tm)

	logService.curLogFile, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("Open Log File Failed, fileName!=%s, err:%s\n", fileName, err)
		return errcode.RESULT_ERROR_FAILED_OPEN_FILE
	}

	logService.logFileCreateTime = tm

	return errcode.RESULT_SUCCESS
}

/*给定一个时间，返回一个规整的日志文件名称，一般为prefix+"_"+YYYYMMDD+".log"*/
func getLogFileName(tm *time.Time) string {
	fileName := fmt.Sprintf("%s_%04d%02d%02d.log", logService.logFilePrefix, tm.Year(), tm.Month(), tm.Day())
	return fileName
}

func PrintLogServiceStats() {
	if logService.isInit {
		fmt.Println("LogService has been initialized")
	} else {
		fmt.Println("LogService has not been initialized!!!")
		return
	}
	fmt.Printf("LogService: Ip=%s,port=%d\n", logService.selfAddr.IP.String(), logService.selfAddr.Port)

	fmt.Printf("LogService: PersitentLevel=%d, ConsoleLevel=%d, MonitorLeve=%d,TrapLevel=%d,ServerLevel=%d,disk_size=%dB,Period=%dDays\n", logService.persistentLevel,
		logService.consoleLevel, logService.monitorLevel, logService.trapLevel, logService.serverLevel, logService.maxStorageSize, logService.maxStoragePeriod)
	len, cap := logService.queue.GetQueueInfo()

	fmt.Printf("LogService: Recv Log Packets=%d, bytes=%d,queue_cap=%d,len=%d\n",
		logService.logPackets, logService.LogBytes, cap, len)

	fmt.Println("--------------Log Module----------------")
	for _, v := range logService.logMuduleControl {
		fmt.Printf("ModuleName=%s, \taddr=%s:%d, PersitentLevel=%d, ConsoleLevel=%d,recv_packets=%d,bytes=%d\n", v.name,
			v.ip.String(), v.port, v.persistent_log_level, v.console_log_level, v.logPackets, v.logBytes)
	}
}

/*Log Service的内部数据获取*/
func (self *TLogService) GetLogFilePath() string {
	return self.logFilePath
}

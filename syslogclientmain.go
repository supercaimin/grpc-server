// syslogclientmain.go
// syslog client进程
package main

import (
	"flag"
	"fmt"

	"hstcmler/errcode"
	"hstcmler/xlog"

	"google.golang.org/grpc"
)

var LogLevelMap = map[int]xlog.LOG_LEVEL{
	0: xlog.LOG_LEVEL_EMERGENCY,
	1: xlog.LOG_LEVEL_ALERT,
	2: xlog.LOG_LEVEL_CRITICAL,
	3: xlog.LOG_LEVEL_ERROR,
	4: xlog.LOG_LEVEL_WARNING,
	5: xlog.LOG_LEVEL_NOTICE,
	6: xlog.LOG_LEVEL_INFO,
	7: xlog.LOG_LEVEL_DEBUG,
}

var syslogcode = map[int]errcode.RESULT{
	0:  0,
	1:  1,
	2:  2,
	3:  3,
	4:  4,
	5:  5,
	6:  6,
	7:  7,
	8:  8,
	9:  9,
	10: 10,
	11: 11,
	12: 12,
	13: 13,
	14: 14,
	15: 15,
	16: 16,
	17: 17,
	18: 18,
	19: 19,
	20: 20,
}

func main() {

	var xlogclient *xlog.XLogger

	intlevel := flag.Int("level", 5, "syslog level:0-7")
	position := flag.String("position", "0/0/1", "position,string")
	logcode := flag.Int("errcode", 12, "syslog code,1-20")
	descript := flag.String("descrip", "", "syslog error description")

	flag.Parse()
	fmt.Println("usage: -level \n -position \n -errcode \n -descrip  \n")

	xlogclient = xlog.NewXLogger("goclient", "127.0.0.1", 30002, "127.0.0.1", 30001)

	xlogclient.Log(LogLevelMap[*intlevel], *position, syslogcode[*logcode], *descript, "")

	fmt.Println(grpc.Version)

}

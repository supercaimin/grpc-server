// syslogmain.go
// syslog server进程
package main

import (
	"fmt"
	"time"

	"hstcmler/errcode"
	"hstcmler/xlog"

	"google.golang.org/grpc"
)

func main() {

	var i int
	var ch chan int

	ret := xlog.InitLogService("127.0.0.1", 30001, "/var/log/", "sonic")
	if ret != errcode.RESULT_SUCCESS {
		fmt.Println("Syslog server init failed!\n")
		return
	}
	/*read db to get config*/
	xlog.LogRedisGet()

	fmt.Println("hello,syslog server test!")
	fmt.Println(grpc.Version)

	//循环等待,以便klish有时间发告警,验证处理过程
	for i = 0; i < 2; i++ {
		time.Sleep(time.Second) //阻塞一秒
	}
	<-ch //程序不退出（go协程并没有返回信息,所以程序会停在此处

	xlog.PrintLogServiceStats()

}
